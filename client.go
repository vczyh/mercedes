package mercedes

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/vczyh/mercedes/proto/pb"
	"google.golang.org/protobuf/proto"
	"net/http"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

var (
	ErrNotAllowedCommand = errors.New("not allowed command")
)

type EventListenFun func(event Event, err error)

type Client struct {
	accessToken        string
	refreshToken       string
	expireIn           int
	refreshWhenConnect bool
	region             Region
	sessionId          string
	conn               *websocket.Conn
	eventListen        EventListenFun
	cmdStatus          *SyncMap[string, chan *pb.AppTwinCommandStatus]
	api                *API
	vehicles           *GetVehiclesResponse

	err    error
	closed atomic.Bool

	mu sync.Mutex
}

func New(opts ...ClientOption) *Client {
	c := &Client{
		region:    RegionChina,
		sessionId: uuid.New().String(),
		cmdStatus: new(SyncMap[string, chan *pb.AppTwinCommandStatus]),
		api:       NewAPI(WithAPIRegion(RegionChina)),
	}

	for _, opt := range opts {
		opt.apply(c)
	}

	return c
}

func (c *Client) Connect(ctx context.Context) error {
	if c.refreshWhenConnect && c.refreshToken != "" {
		if err := c.refreshAccessToken(ctx); err != nil {
			return err
		}
	}

	vehicles, err := c.api.GetVehicles(ctx, c.accessToken)
	if err != nil {
		return err
	}
	c.vehicles = vehicles

	dialer := websocket.DefaultDialer
	header := http.Header{
		"Authorization":      {c.accessToken},
		"X-SessionId":        {c.sessionId},
		"X-TrackingId":       {uuid.New().String()},
		"RIS-OS-Name":        {"ios"},
		"RIS-OS-Version":     {"17.5.1"},
		"ris-websocket-type": {"ios-native"},
		"RIS-SDK-Version":    {"2.113.0"},
		"X-Locale":           {"zh-CN"},
		"Accept-Language":    {"zh-CN,zh-Hans;q=0.9"},
	}
	conn, _, err := dialer.DialContext(ctx, RegionProviders[c.region].WebSocketURL, header)
	if err != nil {
		return err
	}
	c.conn = conn

	go func() {
		if err := c.handleWebSocket(ctx); err != nil {
			c.mu.Lock()
			c.err = err
			c.mu.Unlock()
			_ = c.Close()
			if c.eventListen != nil {
				c.eventListen(nil, err)
			}
		}
	}()

	return nil
}

func (c *Client) DoorsUnLock(ctx context.Context, vin, pin string) error {
	if err := c.Error(); err != nil {
		return err
	}

	if err := c.checkCommandCapability(ctx, vin, CommandNameDoorsUnlock); err != nil {
		return err
	}

	requestId := uuid.New().String()
	command := &pb.ClientMessage{
		Msg: &pb.ClientMessage_CommandRequest{
			CommandRequest: &pb.CommandRequest{
				Vin:       vin,
				RequestId: requestId,
				Command: &pb.CommandRequest_DoorsUnlock{
					DoorsUnlock: &pb.DoorsUnlock{
						Pin: pin,
					},
				},
			},
		},
	}
	if err := c.writeMessage(command); err != nil {
		return err
	}

	return c.waitCommand(requestId)
}

func (c *Client) Closed() bool {
	return c.closed.Load()
}

func (c *Client) Close() error {
	if !c.closed.Load() {
		_ = c.conn.Close()
		c.cmdStatus.Range(func(_ string, status chan *pb.AppTwinCommandStatus) bool {
			close(status)
			return true
		})
		c.closed.Store(true)
	}
	return nil
}

func (c *Client) AccessToken() string {
	return c.accessToken
}

func (c *Client) checkCommandCapability(ctx context.Context, vin string, commandName string) error {
	for _, vehicle := range c.vehicles.AssignedVehicles {
		if vehicle.Vin == vin {
			capabilities, err := c.api.GetCommandCapabilities(ctx, c.accessToken, vin)
			if err != nil {
				return err
			}

			for _, command := range capabilities.Commands {
				if command.CommandName == commandName {
					if command.IsAvailable {
						return nil
					} else {
						return ErrNotAllowedCommand
					}
				}
			}
		}
	}
	return ErrNotAllowedCommand
}

func (c *Client) waitCommand(requestId string) error {
	status := make(chan *pb.AppTwinCommandStatus)
	c.cmdStatus.Set(requestId, status)
	defer c.cmdStatus.Delete(requestId)

	for msg := range status {
		if msg.State == pb.VehicleAPI_FINISHED {
			return nil
		} else if msg.State == pb.VehicleAPI_FAILED {
			var errStrs []string
			for _, err := range msg.Errors {
				errStrs = append(errStrs, err.String())
			}
			return fmt.Errorf("execute command failed: %s", strings.Join(errStrs, " "))
		}
	}

	return c.Error()
}

func (c *Client) handleWebSocket(ctx context.Context) error {
	for {
		_, b, err := c.conn.ReadMessage()
		if err != nil {
			return err
		}
		var pm pb.PushMessage
		if err := proto.Unmarshal(b, &pm); err != nil {
			return err
		}

		events, err := c.handlePushMessage(&pm)
		if err != nil {
			return err
		}

		if c.eventListen != nil {
			for _, e := range events {
				c.eventListen(e, nil)
			}
		}
	}
}

func (c *Client) handlePushMessage(pm *pb.PushMessage) ([]Event, error) {
	switch v := pm.Msg.(type) {
	case *pb.PushMessage_VepUpdates:
		return c.handleVepUpdates(v), nil
	case *pb.PushMessage_ApptwinCommandStatusUpdatesByVin:
		return nil, c.handleApptwinCommandStatusUpdatesByVin(v)
	default:
		fmt.Printf("should handle %T push message\n", v)
		return nil, nil
	}
}

func (c *Client) handleVepUpdates(message *pb.PushMessage_VepUpdates) []Event {
	var events []Event

	for _, update := range message.VepUpdates.Updates {
		//sequenceNumber := update.SequenceNumber
		for name, status := range update.Attributes {
			statusType := StatusType(status.Status)
			if statusType != StatusTypeValid {
				continue
			}
			attributeStatus := AttributeStatus{
				Vin:     update.Vin,
				Time:    time.UnixMilli(status.TimestampInMs),
				Status:  statusType,
				Changed: status.Changed,
			}
			switch name {
			case AttributeStarterBatteryState:
				e := StarterBatteryStateEvent{
					AttributeStatus: attributeStatus,
					State:           StarterBatteryState(status.AttributeType.(*pb.VehicleAttributeStatus_IntValue).IntValue),
				}
				events = append(events, e)
			case AttributeEngineState:
				e := EngineStateEvent{
					AttributeStatus: attributeStatus,
					Running:         status.AttributeType.(*pb.VehicleAttributeStatus_BoolValue).BoolValue,
				}
				events = append(events, e)
			case AttributeDistanceReset:
				e := DistanceResetEvent{
					AttributeStatus: attributeStatus,
					Value:           status.AttributeType.(*pb.VehicleAttributeStatus_DoubleValue).DoubleValue,
				}
				events = append(events, e)
			case AttributeAverageSpeedReset:
				e := AverageSpeedResetEvent{
					AttributeStatus: attributeStatus,
					Value:           status.AttributeType.(*pb.VehicleAttributeStatus_DoubleValue).DoubleValue,
				}
				events = append(events, e)
			case AttributeDrivenTimeReset:
				minutes := status.AttributeType.(*pb.VehicleAttributeStatus_IntValue).IntValue
				e := DrivenTimeResetEvent{
					AttributeStatus: attributeStatus,
					Value:           time.Minute * time.Duration(minutes),
				}
				events = append(events, e)
			case AttributeLiquidConsumptionReset:
				e := LiquidConsumptionEvent{
					AttributeStatus: attributeStatus,
					Value:           status.AttributeType.(*pb.VehicleAttributeStatus_DoubleValue).DoubleValue,
				}
				events = append(events, e)
			case AttributeDistanceZEReset:
				e := DistanceZEResetEvent{
					AttributeStatus: attributeStatus,
					Value:           status.AttributeType.(*pb.VehicleAttributeStatus_DoubleValue).DoubleValue,
				}
				events = append(events, e)
			case AttributeDrivenTimeZEReset:
				events = append(events, DrivenTimeZEResetEvent{
					AttributeStatus: attributeStatus,
					Value:           time.Minute * time.Duration(status.AttributeType.(*pb.VehicleAttributeStatus_IntValue).IntValue),
				})
			case AttributeDistanceStart:
				e := DistanceStartEvent{
					AttributeStatus: attributeStatus,
					Value:           status.AttributeType.(*pb.VehicleAttributeStatus_DoubleValue).DoubleValue,
				}
				events = append(events, e)
			case AttributeAverageSpeedStart:
				e := AverageSpeedStartEvent{
					AttributeStatus: attributeStatus,
					Value:           status.AttributeType.(*pb.VehicleAttributeStatus_DoubleValue).DoubleValue,
				}
				events = append(events, e)
			case AttributeDrivenTimeStart:
				minutes := status.AttributeType.(*pb.VehicleAttributeStatus_IntValue).IntValue
				e := DrivenTimeStartEvent{
					AttributeStatus: attributeStatus,
					Value:           time.Minute * time.Duration(minutes),
				}
				events = append(events, e)
			case AttributeLiquidConsumptionStart:
				e := LiquidConsumptionStartEvent{
					AttributeStatus: attributeStatus,
					Value:           status.AttributeType.(*pb.VehicleAttributeStatus_DoubleValue).DoubleValue,
				}
				events = append(events, e)
			case AttributeDistanceZEStart:
				e := DistanceZEStartEvent{
					AttributeStatus: attributeStatus,
					Value:           status.AttributeType.(*pb.VehicleAttributeStatus_DoubleValue).DoubleValue,
				}
				events = append(events, e)
			case AttributeDrivenTimeZEStart:
				events = append(events, DrivenTimeZEStartEvent{
					AttributeStatus: attributeStatus,
					Value:           time.Minute * time.Duration(status.AttributeType.(*pb.VehicleAttributeStatus_IntValue).IntValue),
				})
			case AttributeOdo:
				e := OdoEvent{
					AttributeStatus: attributeStatus,
					Value:           int(status.AttributeType.(*pb.VehicleAttributeStatus_IntValue).IntValue),
				}
				events = append(events, e)
			case AttributeOilLevel:
				e := OilLevelEvent{
					AttributeStatus: attributeStatus,
					Level:           int(status.AttributeType.(*pb.VehicleAttributeStatus_IntValue).IntValue),
				}
				events = append(events, e)
			case AttributeRangeLiquid:
				e := RangeLiquidEvent{
					AttributeStatus: attributeStatus,
					Value:           int(status.AttributeType.(*pb.VehicleAttributeStatus_IntValue).IntValue),
				}
				events = append(events, e)
			case AttributeTankLevelPercent:
				e := TankLevelPercentEvent{
					AttributeStatus: attributeStatus,
					Value:           int(status.AttributeType.(*pb.VehicleAttributeStatus_IntValue).IntValue),
				}
				events = append(events, e)
			case AttributeRoofTopStatus:
				e := RoofTopStatusEvent{
					AttributeStatus: attributeStatus,
					State:           RoofTopState(status.AttributeType.(*pb.VehicleAttributeStatus_IntValue).IntValue),
				}
				events = append(events, e)
			case AttributeDoorStatusOverall:
				e := DoorStatusOverallEvent{
					AttributeStatus: attributeStatus,
					State:           DoorOverallStatus(status.AttributeType.(*pb.VehicleAttributeStatus_IntValue).IntValue),
				}
				events = append(events, e)
			case AttributeDoorStatusFrontLeft:
				e := DoorStatusFrontLeftEvent{
					AttributeStatus: attributeStatus,
					Open:            status.AttributeType.(*pb.VehicleAttributeStatus_BoolValue).BoolValue,
				}
				events = append(events, e)
			case AttributeDoorStatusFrontRight:
				e := DoorStatusFrontRightEvent{
					AttributeStatus: attributeStatus,
					Open:            status.AttributeType.(*pb.VehicleAttributeStatus_BoolValue).BoolValue,
				}
				events = append(events, e)
			case AttributeDoorStatusRearLeft:
				e := DoorStatusRearLeftEvent{
					AttributeStatus: attributeStatus,
					Open:            status.AttributeType.(*pb.VehicleAttributeStatus_BoolValue).BoolValue,
				}
				events = append(events, e)
			case AttributeDoorStatusRearRight:
				e := DoorStatusRearRightEvent{
					AttributeStatus: attributeStatus,
					Open:            status.AttributeType.(*pb.VehicleAttributeStatus_BoolValue).BoolValue,
				}
				events = append(events, e)
			case AttributeDeckLidStatus:
				e := DeckLidStatusEvent{
					AttributeStatus: attributeStatus,
					Open:            status.AttributeType.(*pb.VehicleAttributeStatus_BoolValue).BoolValue,
				}
				events = append(events, e)
			case AttributeDoorStatusGas:
				e := DoorStatusGasEvent{
					AttributeStatus: attributeStatus,
					State:           int(status.AttributeType.(*pb.VehicleAttributeStatus_IntValue).IntValue),
				}
				events = append(events, e)
			case AttributeDoorLockStatusOverall:
				e := DoorLockStatusOverallEvent{
					AttributeStatus: attributeStatus,
					State:           DoorLockOverallStatus(status.AttributeType.(*pb.VehicleAttributeStatus_IntValue).IntValue),
				}
				events = append(events, e)
			case AttributeDoorLockStatusFrontLeft:
				e := DoorLockStatusFrontLeftEvent{
					AttributeStatus: attributeStatus,
					UnLocked:        status.AttributeType.(*pb.VehicleAttributeStatus_BoolValue).BoolValue,
				}
				events = append(events, e)
			case AttributeDoorLockStatusFrontRight:
				e := DoorLockStatusFrontRightEvent{
					AttributeStatus: attributeStatus,
					UnLocked:        status.AttributeType.(*pb.VehicleAttributeStatus_BoolValue).BoolValue,
				}
				events = append(events, e)
			case AttributeDoorLockStatusRearLeft:
				e := DoorLockStatusRearLeftEvent{
					AttributeStatus: attributeStatus,
					UnLocked:        status.AttributeType.(*pb.VehicleAttributeStatus_BoolValue).BoolValue,
				}
				events = append(events, e)
			case AttributeDoorLockStatusRearRight:
				e := DoorLockStatusRearRightEvent{
					AttributeStatus: attributeStatus,
					UnLocked:        status.AttributeType.(*pb.VehicleAttributeStatus_BoolValue).BoolValue,
				}
				events = append(events, e)
			case AttributeDoorLockStatusGas:
				e := DoorLockStatusGasEvent{
					AttributeStatus: attributeStatus,
					UnLocked:        status.AttributeType.(*pb.VehicleAttributeStatus_BoolValue).BoolValue,
				}
				events = append(events, e)
			case AttributeDoorLockStatusDeckLid:
				e := DoorLockStatusDeckLidEvent{
					AttributeStatus: attributeStatus,
					UnLocked:        status.AttributeType.(*pb.VehicleAttributeStatus_BoolValue).BoolValue,
				}
				events = append(events, e)
			case AttributeDoorLockStatusVehicle:
				e := DoorLockStatusVehicleEvent{
					AttributeStatus: attributeStatus,
					State:           int(status.AttributeType.(*pb.VehicleAttributeStatus_IntValue).IntValue),
				}
				events = append(events, e)
			case AttributeWindowStatusOverall:
				e := WindowStatusOverallEvent{
					AttributeStatus: attributeStatus,
					State:           WindowsOverallStatus(status.AttributeType.(*pb.VehicleAttributeStatus_IntValue).IntValue),
				}
				events = append(events, e)
			case AttributeWindowStatusRearBlind:
				e := WindowStatusRearBlindEvent{
					AttributeStatus: attributeStatus,
				}
				events = append(events, e)
			case AttributeWindowStatusRearLeftBlind:
				e := WindowStatusRearLeftBlindEvent{
					AttributeStatus: attributeStatus,
				}
				events = append(events, e)
			case AttributeWindowStatusRearRightBlind:
				e := WindowStatusRearRightBlindEvent{
					AttributeStatus: attributeStatus,
				}
				events = append(events, e)
			case AttributeWindowStatusFrontLeft:
				e := WindowStatusFrontLeftEvent{
					AttributeStatus: attributeStatus,
					State:           WindowStatus(status.AttributeType.(*pb.VehicleAttributeStatus_IntValue).IntValue),
				}
				events = append(events, e)
			case AttributeWindowStatusFrontRight:
				e := WindowStatusFrontRightEvent{
					AttributeStatus: attributeStatus,
					State:           WindowStatus(status.AttributeType.(*pb.VehicleAttributeStatus_IntValue).IntValue),
				}
				events = append(events, e)
			case AttributeWindowStatusRearLeft:
				e := WindowStatusRearLeftEvent{
					AttributeStatus: attributeStatus,
					State:           WindowStatus(status.AttributeType.(*pb.VehicleAttributeStatus_IntValue).IntValue),
				}
				events = append(events, e)
			case AttributeWindowStatusRearRight:
				e := WindowStatusRearRightEvent{
					AttributeStatus: attributeStatus,
					State:           WindowStatus(status.AttributeType.(*pb.VehicleAttributeStatus_IntValue).IntValue),
				}
				events = append(events, e)
			case AttributeSunroofStatus:
				e := SunroofStatusEvent{
					AttributeStatus: attributeStatus,
					State:           SunroofState(status.AttributeType.(*pb.VehicleAttributeStatus_IntValue).IntValue),
				}
				events = append(events, e)
			case AttributeWarningWashWater:
				e := WarningWashWaterEvent{
					AttributeStatus: attributeStatus,
					Warning:         status.AttributeType.(*pb.VehicleAttributeStatus_BoolValue).BoolValue,
				}
				events = append(events, e)
			case AttributeWarningCoolantLevelLow:
				e := WarningCoolantLevelLowEvent{
					AttributeStatus: attributeStatus,
					Warning:         status.AttributeType.(*pb.VehicleAttributeStatus_BoolValue).BoolValue,
				}
				events = append(events, e)
			case AttributeWarningBrakeFluid:
				e := WarningBrakeFluidEvent{
					AttributeStatus: attributeStatus,
					Warning:         status.AttributeType.(*pb.VehicleAttributeStatus_BoolValue).BoolValue,
				}
				events = append(events, e)
			case AttributeWarningBrakeLiningWear:
				e := WarningBrakeLiningWearEvent{
					AttributeStatus: attributeStatus,
					Warning:         status.AttributeType.(*pb.VehicleAttributeStatus_BoolValue).BoolValue,
				}
				events = append(events, e)
			case AttributeTirePressureFrontLeft:
				e := TirePressureFrontLeftEvent{
					AttributeStatus: attributeStatus,
					Value:           status.AttributeType.(*pb.VehicleAttributeStatus_DoubleValue).DoubleValue,
				}
				events = append(events, e)
			case AttributeTirePressureFrontRight:
				e := TirePressureFrontRightEvent{
					AttributeStatus: attributeStatus,
					Value:           status.AttributeType.(*pb.VehicleAttributeStatus_DoubleValue).DoubleValue,
				}
				events = append(events, e)
			case AttributeTirePressureRearLeft:
				e := TirePressureRearLeftEvent{
					AttributeStatus: attributeStatus,
					Value:           status.AttributeType.(*pb.VehicleAttributeStatus_DoubleValue).DoubleValue,
				}
				events = append(events, e)
			case AttributeTirePressureRearRight:
				e := TirePressureRearRightEvent{
					AttributeStatus: attributeStatus,
					Value:           status.AttributeType.(*pb.VehicleAttributeStatus_DoubleValue).DoubleValue,
				}
				events = append(events, e)
			case AttributeEngineHoodStatus:
				events = append(events, EngineHoodStatusEvent{
					AttributeStatus: attributeStatus,
					Open:            status.AttributeType.(*pb.VehicleAttributeStatus_BoolValue).BoolValue,
				})
			case AttributeFilterParticleLoading:
				events = append(events, FilterParticleLoadingEvent{
					AttributeStatus: attributeStatus,
					State:           FilterParticleState(status.AttributeType.(*pb.VehicleAttributeStatus_IntValue).IntValue),
				})
			case AttributeIgnitionState:
				events = append(events, IgnitionStateEvent{
					AttributeStatus: attributeStatus,
					State:           IgnitionState(status.AttributeType.(*pb.VehicleAttributeStatus_IntValue).IntValue),
				})
			case AttributeLanguageHU:
				events = append(events, LanguageHUEvent{
					AttributeStatus: attributeStatus,
					State:           LanguageState(status.AttributeType.(*pb.VehicleAttributeStatus_IntValue).IntValue),
				})
			case AttributeLiquidRangeSkipIndication:
				events = append(events, LiquidRangeSkipIndicationEvent{
					AttributeStatus: attributeStatus,
					Skip:            status.AttributeType.(*pb.VehicleAttributeStatus_BoolValue).BoolValue,
				})
			case AttributeOverallRange:
				events = append(events, OverallRangeEvent{
					AttributeStatus: attributeStatus,
					Value:           status.AttributeType.(*pb.VehicleAttributeStatus_DoubleValue).DoubleValue,
				})
			case AttributeParkBrakeStatus:
				events = append(events, ParkBrakeStatusEvent{
					AttributeStatus: attributeStatus,
					Active:          status.AttributeType.(*pb.VehicleAttributeStatus_BoolValue).BoolValue,
				})
			case AttributePositionHeading:
				events = append(events, PositionHeadingEvent{
					AttributeStatus: attributeStatus,
					Value:           status.AttributeType.(*pb.VehicleAttributeStatus_DoubleValue).DoubleValue,
				})
			case AttributePositionLat:
				events = append(events, PositionLatEvent{
					AttributeStatus: attributeStatus,
					Value:           status.AttributeType.(*pb.VehicleAttributeStatus_DoubleValue).DoubleValue,
				})
			case AttributePositionLong:
				events = append(events, PositionLongEvent{
					AttributeStatus: attributeStatus,
					Value:           status.AttributeType.(*pb.VehicleAttributeStatus_DoubleValue).DoubleValue,
				})
			case AttributeVehiclePositionErrorCode:
				events = append(events, PositionErrorCodeEvent{
					AttributeStatus: attributeStatus,
					State:           PositionErrorState(status.AttributeType.(*pb.VehicleAttributeStatus_IntValue).IntValue),
				})
			case AttributeProximityCalculationForVehiclePositionRequired:
				events = append(events, ProximityCalculationForVehiclePositionRequiredEvent{
					AttributeStatus: attributeStatus,
					Required:        status.AttributeType.(*pb.VehicleAttributeStatus_BoolValue).BoolValue,
				})
			case AttributeRemoteStartActive:
				events = append(events, RemoteStartActiveEvent{
					AttributeStatus: attributeStatus,
					Active:          status.AttributeType.(*pb.VehicleAttributeStatus_BoolValue).BoolValue,
				})
			case AttributeRemoteStartEndTime:
				events = append(events, RemoteStartEndTimeEvent{
					AttributeStatus: attributeStatus,
					Time:            time.Unix(status.AttributeType.(*pb.VehicleAttributeStatus_IntValue).IntValue, 0),
				})
			case AttributeRemoteStartTemperature:
				events = append(events, RemoteStartTemperatureEvent{
					AttributeStatus: attributeStatus,
					Value:           status.AttributeType.(*pb.VehicleAttributeStatus_DoubleValue).DoubleValue,
				})
			case AttributeServiceIntervalDays:
				events = append(events, ServiceIntervalDaysEvent{
					AttributeStatus: attributeStatus,
					Value:           int(status.AttributeType.(*pb.VehicleAttributeStatus_IntValue).IntValue),
				})
			case AttributeServiceIntervalDistance:
				events = append(events, ServiceIntervalDistanceEvent{
					AttributeStatus: attributeStatus,
					Value:           int(status.AttributeType.(*pb.VehicleAttributeStatus_IntValue).IntValue),
				})
			case AttributeSoc:
				events = append(events, SocEvent{
					AttributeStatus: attributeStatus,
					Value:           int(status.AttributeType.(*pb.VehicleAttributeStatus_IntValue).IntValue),
				})
			case AttributeSpeedUnitFromIC:
				events = append(events, SpeedUnitFromICEvent{
					AttributeStatus: attributeStatus,
					KmUnit:          !status.AttributeType.(*pb.VehicleAttributeStatus_BoolValue).BoolValue,
				})
			case AttributeSunroofEvent:
				events = append(events, SunroofEventEvent{
					AttributeStatus: attributeStatus,
					State:           SunroofEventState(status.AttributeType.(*pb.VehicleAttributeStatus_IntValue).IntValue),
				})
			case AttributeSunroofEventActive:
				events = append(events, SunroofEventActiveEvent{
					AttributeStatus: attributeStatus,
					Active:          status.AttributeType.(*pb.VehicleAttributeStatus_BoolValue).BoolValue,
				})
			case AttributeSunroofStatusFrontBlind:
				events = append(events, SunroofStatusFrontBlindEvent{
					AttributeStatus: attributeStatus,
					State:           int(status.AttributeType.(*pb.VehicleAttributeStatus_IntValue).IntValue),
				})
			case AttributeSunroofStatusRearBlind:
				events = append(events, SunroofStatusRearBlindEvent{
					AttributeStatus: attributeStatus,
					State:           int(status.AttributeType.(*pb.VehicleAttributeStatus_IntValue).IntValue),
				})
			case AttributeTemperaturePoints:
				e := TemperaturePointsEvent{
					AttributeStatus: attributeStatus,
				}
				points := status.AttributeType.(*pb.VehicleAttributeStatus_TemperaturePointsValue).TemperaturePointsValue
				for _, point := range points.TemperaturePoints {
					temperature := point.Temperature
					switch point.Zone {
					case "frontCenter":
						e.FrontCenter = temperature
					case "frontLeft":
						e.FrontLeft = temperature
					case "frontRight":
						e.FrontRight = temperature
					case "rearCenter":
						e.RearCenter = temperature
					case "rear2center":
						e.Rear2Center = temperature
					case "rearLeft":
						e.RearLeft = temperature
					case "rearRight":
						e.RearRight = temperature
					}
				}
				events = append(events, e)
			case AttributeTemperatureUnitHU:
				events = append(events, TemperatureUnitHUEvent{
					AttributeStatus: attributeStatus,
					CelsiusUnit:     !status.AttributeType.(*pb.VehicleAttributeStatus_BoolValue).BoolValue,
				})
			case AttributeTireMarkerFrontLeft:
				events = append(events, TireMarkerFrontLeftEvent{
					AttributeStatus: attributeStatus,
					WarningLevel:    TireMarkerWarningLevel(status.AttributeType.(*pb.VehicleAttributeStatus_IntValue).IntValue),
				})
			case AttributeTireMarkerFrontRight:
				events = append(events, TireMarkerFrontRightEvent{
					AttributeStatus: attributeStatus,
					WarningLevel:    TireMarkerWarningLevel(status.AttributeType.(*pb.VehicleAttributeStatus_IntValue).IntValue),
				})
			case AttributeTireMarkerRearLeft:
				events = append(events, TireMarkerRearLeftEvent{
					AttributeStatus: attributeStatus,
					WarningLevel:    TireMarkerWarningLevel(status.AttributeType.(*pb.VehicleAttributeStatus_IntValue).IntValue),
				})
			case AttributeTireMarkerRearRight:
				events = append(events, TireMarkerRearRightEvent{
					AttributeStatus: attributeStatus,
					WarningLevel:    TireMarkerWarningLevel(status.AttributeType.(*pb.VehicleAttributeStatus_IntValue).IntValue),
				})
			case AttributeTirePressMeasTimestamp:
				events = append(events, TirePressMeasTimestampEvent{
					AttributeStatus: attributeStatus,
					Time:            time.UnixMilli(status.AttributeType.(*pb.VehicleAttributeStatus_IntValue).IntValue),
				})
			case AttributeTireSensorAvailable:
				events = append(events, TireSensorAvailableEvent{
					AttributeStatus: attributeStatus,
					State:           TireSensorState(status.AttributeType.(*pb.VehicleAttributeStatus_IntValue).IntValue),
				})
			case AttributeTireWarningsRDK:
				events = append(events, TireWarningLevelOverallEvent{
					AttributeStatus: attributeStatus,
					WarningLevel:    TireMarkerWarningLevel(status.AttributeType.(*pb.VehicleAttributeStatus_IntValue).IntValue),
				})
			case AttributeTireWarningLamp:
				events = append(events, TireWarningLampEvent{
					AttributeStatus: attributeStatus,
					State:           TireLampState(status.AttributeType.(*pb.VehicleAttributeStatus_IntValue).IntValue),
				})
			case AttributeTireWarningSprw:
				events = append(events, TireWarningSprwEvent{
					AttributeStatus: attributeStatus,
					Active:          status.AttributeType.(*pb.VehicleAttributeStatus_BoolValue).BoolValue,
				})
			case AttributeTireWarningLevelPrw:
				events = append(events, TireWarningLevelPrwEvent{
					AttributeStatus: attributeStatus,
					Level:           TireLevelPrwWarning(status.AttributeType.(*pb.VehicleAttributeStatus_IntValue).IntValue),
				})
			case AttributeTrackingStateHU:
				events = append(events, TrackingStateHUEvent{
					AttributeStatus: attributeStatus,
					Active:          status.AttributeType.(*pb.VehicleAttributeStatus_BoolValue).BoolValue,
				})
			case AttributeVehicleDataConnectionState:
				events = append(events, VehicleDataConnectionStateEvent{
					AttributeStatus: attributeStatus,
					Active:          status.AttributeType.(*pb.VehicleAttributeStatus_BoolValue).BoolValue,
				})
			case AttributeVehicleLockState:
				events = append(events, VehicleLockStateEvent{
					AttributeStatus: attributeStatus,
					State:           VehicleLockState(status.AttributeType.(*pb.VehicleAttributeStatus_IntValue).IntValue),
				})
			case AttributeVTime:
				events = append(events, VehicleTimeEvent{
					AttributeStatus: attributeStatus,
					Time:            time.Unix(status.AttributeType.(*pb.VehicleAttributeStatus_IntValue).IntValue, 0),
				})
			case AttributeWiperHealthPercent:
				events = append(events, WiperHealthPercentEvent{
					AttributeStatus: attributeStatus,
					Value:           int(status.AttributeType.(*pb.VehicleAttributeStatus_IntValue).IntValue),
				})
			case AttributeWiperLifetimeExceeded:
				events = append(events, WiperLifetimeExceededEvent{
					AttributeStatus: attributeStatus,
					Exceeded:        !status.AttributeType.(*pb.VehicleAttributeStatus_BoolValue).BoolValue,
				})
			default:
				//fmt.Println("unsupported name:", name)
			}
		}
	}

	return events
}

func (c *Client) handleApptwinCommandStatusUpdatesByVin(message *pb.PushMessage_ApptwinCommandStatusUpdatesByVin) error {
	update := message.ApptwinCommandStatusUpdatesByVin
	clientMessage := &pb.ClientMessage{
		Msg: &pb.ClientMessage_AcknowledgeApptwinCommandStatusUpdateByVin{
			AcknowledgeApptwinCommandStatusUpdateByVin: &pb.AcknowledgeAppTwinCommandStatusUpdatesByVIN{
				SequenceNumber: update.SequenceNumber,
			}},
	}
	if err := c.writeMessage(clientMessage); err != nil {
		return err
	}

	for _, updateByPid := range update.UpdatesByVin {
		for _, commandStatus := range updateByPid.UpdatesByPid {
			c, ok := c.cmdStatus.Get(commandStatus.RequestId)
			if !ok {
				continue
			}
			c <- commandStatus
		}
	}
	return nil
}

func (c *Client) refreshAccessToken(ctx context.Context) error {
	res, err := c.api.RefreshToken(ctx, c.refreshToken)
	if err != nil {
		return err
	}

	c.accessToken = res.AccessToken
	c.expireIn = res.ExpiresIn
	return nil
}

func (c *Client) writeMessage(message proto.Message) error {
	b, err := proto.Marshal(message)
	if err != nil {
		return err
	}
	return c.conn.WriteMessage(websocket.BinaryMessage, b)
}

func (c *Client) Error() error {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.err
}

func WithAccessToken(accessToken string) ClientOption {
	return clientOptionFun(func(c *Client) {
		c.accessToken = accessToken
	})
}

func WithRefreshToken(refreshToken string) ClientOption {
	return clientOptionFun(func(c *Client) {
		c.refreshToken = refreshToken
	})
}

func WithRefreshWhenConnect(refresh bool) ClientOption {
	return clientOptionFun(func(c *Client) {
		c.refreshWhenConnect = refresh
	})
}

func WithRegion(region Region) ClientOption {
	return clientOptionFun(func(c *Client) {
		c.region = region
	})
}

func WithEventListen(f EventListenFun) ClientOption {
	return clientOptionFun(func(c *Client) {
		c.eventListen = f
	})
}

type ClientOption interface {
	apply(*Client)
}
type clientOptionFun func(c *Client)

func (f clientOptionFun) apply(c *Client) {
	f(c)
}
