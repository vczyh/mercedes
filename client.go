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
	accessToken  string
	refreshToken string
	expireIn     int
	region       Region
	sessionId    string
	conn         *websocket.Conn
	eventListen  EventListenFun
	cmdStatus    *SyncMap[string, chan *pb.AppTwinCommandStatus]
	api          *API
	vehicles     *GetVehiclesResponse

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
	if c.refreshToken != "" {
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
				Vin:           update.Vin,
				TimestampInMs: status.TimestampInMs,
				Status:        statusType,
				Changed:       status.Changed,
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
					State:           int(status.AttributeType.(*pb.VehicleAttributeStatus_IntValue).IntValue),
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
			case AttributeSunRoofStatus:
				e := SunRoofStatusEvent{
					AttributeStatus: attributeStatus,
					State:           SunRoofStatus(status.AttributeType.(*pb.VehicleAttributeStatus_IntValue).IntValue),
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
