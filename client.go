package mercedes

import (
	"context"
	"fmt"
	"github.com/carlmjohnson/requests"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/vczyh/mercedes/proto/pb"
	"google.golang.org/protobuf/proto"
	"net/http"
	"net/url"
	"time"
)

type ListenFun func(event MercedesEvent)

type Client struct {
	accessToken  string
	refreshToken string
	expireIn     int
	region       Region
	sessionId    string
	conn         *websocket.Conn
}

func New(opts ...ClientOption) *Client {
	c := &Client{
		region:    RegionChina,
		sessionId: uuid.New().String(),
	}
	for _, opt := range opts {
		opt.apply(c)
	}
	return c
}

func (c *Client) Connect(ctx context.Context) error {
	//if c.refreshToken != "" {
	//	if err := c.refreshAccessToken(ctx); err != nil {
	//		return err
	//	}
	//}

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

	return nil
}

func (c *Client) OnListen(f ListenFun) error {
	for {
		_, msg, err := c.conn.ReadMessage()
		if err != nil {
			return err
		}
		events, err := c.handle(msg)
		if err != nil {
			return err
		}
		for _, e := range events {
			f(e)
		}
	}
}
func (c *Client) handle(b []byte) ([]MercedesEvent, error) {
	var pm pb.PushMessage
	err := proto.Unmarshal(b, &pm)
	if err != nil {
		return nil, err
	}
	switch v := pm.Msg.(type) {
	case *pb.PushMessage_VepUpdates:
		return c.handleVepUpdates(v), nil
	default:
		fmt.Printf("should handle %T push message\n", v)
	}
	return nil, nil
}

func (c *Client) handleVepUpdates(message *pb.PushMessage_VepUpdates) []MercedesEvent {
	var events []MercedesEvent

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
				e := StarterBatteryState{
					AttributeStatus: attributeStatus,
					State:           StarterBatteryStateType(status.AttributeType.(*pb.VehicleAttributeStatus_IntValue).IntValue),
				}
				events = append(events, e)
			case AttributeEngineState:
				e := EngineState{
					AttributeStatus: attributeStatus,
					Running:         status.AttributeType.(*pb.VehicleAttributeStatus_BoolValue).BoolValue,
				}
				events = append(events, e)
			case AttributeDistanceReset:
				e := DistanceReset{
					AttributeStatus: attributeStatus,
					Value:           status.AttributeType.(*pb.VehicleAttributeStatus_DoubleValue).DoubleValue,
				}
				events = append(events, e)
			case AttributeAverageSpeedReset:
				e := AverageSpeedReset{
					AttributeStatus: attributeStatus,
					Value:           status.AttributeType.(*pb.VehicleAttributeStatus_DoubleValue).DoubleValue,
				}
				events = append(events, e)
			case AttributeDrivenTimeReset:
				minutes := status.AttributeType.(*pb.VehicleAttributeStatus_IntValue).IntValue
				e := DrivenTimeReset{
					AttributeStatus: attributeStatus,
					Value:           time.Minute * time.Duration(minutes),
				}
				events = append(events, e)
			case AttributeLiquidConsumptionReset:
				e := LiquidConsumptionReset{
					AttributeStatus: attributeStatus,
					Value:           status.AttributeType.(*pb.VehicleAttributeStatus_DoubleValue).DoubleValue,
				}
				events = append(events, e)
			case AttributeDistanceStart:
				e := DistanceStart{
					AttributeStatus: attributeStatus,
					Value:           status.AttributeType.(*pb.VehicleAttributeStatus_DoubleValue).DoubleValue,
				}
				events = append(events, e)
			case AttributeAverageSpeedStart:
				e := AverageSpeedStart{
					AttributeStatus: attributeStatus,
					Value:           status.AttributeType.(*pb.VehicleAttributeStatus_DoubleValue).DoubleValue,
				}
				events = append(events, e)
			case AttributeDrivenTimeStart:
				minutes := status.AttributeType.(*pb.VehicleAttributeStatus_IntValue).IntValue
				e := DrivenTimeStart{
					AttributeStatus: attributeStatus,
					Value:           time.Minute * time.Duration(minutes),
				}
				events = append(events, e)
			case AttributeLiquidConsumptionStart:
				e := LiquidConsumptionStart{
					AttributeStatus: attributeStatus,
					Value:           status.AttributeType.(*pb.VehicleAttributeStatus_DoubleValue).DoubleValue,
				}
				events = append(events, e)
			case AttributeOdo:
				e := Odo{
					AttributeStatus: attributeStatus,
					Value:           int(status.AttributeType.(*pb.VehicleAttributeStatus_IntValue).IntValue),
				}
				events = append(events, e)
			case AttributeOilLevel:
				e := OilLevel{
					AttributeStatus: attributeStatus,
					Level:           int(status.AttributeType.(*pb.VehicleAttributeStatus_IntValue).IntValue),
				}
				events = append(events, e)
			case AttributeRangeLiquid:
				e := RangeLiquid{
					AttributeStatus: attributeStatus,
					Value:           int(status.AttributeType.(*pb.VehicleAttributeStatus_IntValue).IntValue),
				}
				events = append(events, e)
			case AttributeTankLevelPercent:
				e := TankLevelPercent{
					AttributeStatus: attributeStatus,
					Value:           int(status.AttributeType.(*pb.VehicleAttributeStatus_IntValue).IntValue),
				}
				events = append(events, e)
			case AttributeRoofTopStatus:
				e := RoofTopStatus{
					AttributeStatus: attributeStatus,
				}
				events = append(events, e)
			case AttributeDoorStatusOverall:
				e := DoorStatusOverall{
					AttributeStatus: attributeStatus,
					State:           int(status.AttributeType.(*pb.VehicleAttributeStatus_IntValue).IntValue),
				}
				events = append(events, e)
			case AttributeDoorStatusFrontLeft:
				e := DoorStatusFrontLeft{
					AttributeStatus: attributeStatus,
					Open:            status.AttributeType.(*pb.VehicleAttributeStatus_BoolValue).BoolValue,
				}
				events = append(events, e)
			case AttributeDoorStatusFrontRight:
				e := DoorStatusFrontRight{
					AttributeStatus: attributeStatus,
					Open:            status.AttributeType.(*pb.VehicleAttributeStatus_BoolValue).BoolValue,
				}
				events = append(events, e)
			case AttributeDoorStatusRearLeft:
				e := DoorStatusRearLeft{
					AttributeStatus: attributeStatus,
					Open:            status.AttributeType.(*pb.VehicleAttributeStatus_BoolValue).BoolValue,
				}
				events = append(events, e)
			case AttributeDoorStatusRearRight:
				e := DoorStatusRearRight{
					AttributeStatus: attributeStatus,
					Open:            status.AttributeType.(*pb.VehicleAttributeStatus_BoolValue).BoolValue,
				}
				events = append(events, e)
			case AttributeDeckLidStatus:
				e := DeckLidStatus{
					AttributeStatus: attributeStatus,
					Open:            status.AttributeType.(*pb.VehicleAttributeStatus_BoolValue).BoolValue,
				}
				events = append(events, e)
			case AttributeDoorStatusGas:
				e := DoorStatusGas{
					AttributeStatus: attributeStatus,
					State:           int(status.AttributeType.(*pb.VehicleAttributeStatus_IntValue).IntValue),
				}
				events = append(events, e)
			case AttributeDoorLockStatusOverall:
				e := DoorLockStatusOverall{
					AttributeStatus: attributeStatus,
					State:           DoorLockOverallStatusType(status.AttributeType.(*pb.VehicleAttributeStatus_IntValue).IntValue),
				}
				events = append(events, e)
			case AttributeDoorLockStatusFrontLeft:
				e := DoorLockStatusFrontLeft{
					AttributeStatus: attributeStatus,
					UnLocked:        status.AttributeType.(*pb.VehicleAttributeStatus_BoolValue).BoolValue,
				}
				events = append(events, e)
			case AttributeDoorLockStatusFrontRight:
				e := DoorLockStatusFrontRight{
					AttributeStatus: attributeStatus,
					UnLocked:        status.AttributeType.(*pb.VehicleAttributeStatus_BoolValue).BoolValue,
				}
				events = append(events, e)
			case AttributeDoorLockStatusRearLeft:
				e := DoorLockStatusRearLeft{
					AttributeStatus: attributeStatus,
					UnLocked:        status.AttributeType.(*pb.VehicleAttributeStatus_BoolValue).BoolValue,
				}
				events = append(events, e)
			case AttributeDoorLockStatusRearRight:
				e := DoorLockStatusRearRight{
					AttributeStatus: attributeStatus,
					UnLocked:        status.AttributeType.(*pb.VehicleAttributeStatus_BoolValue).BoolValue,
				}
				events = append(events, e)
			case AttributeDoorLockStatusGas:
				e := DoorLockStatusGas{
					AttributeStatus: attributeStatus,
					UnLocked:        status.AttributeType.(*pb.VehicleAttributeStatus_BoolValue).BoolValue,
				}
				events = append(events, e)
			case AttributeDoorLockStatusDeckLid:
				e := DoorLockStatusDeckLid{
					AttributeStatus: attributeStatus,
					UnLocked:        status.AttributeType.(*pb.VehicleAttributeStatus_BoolValue).BoolValue,
				}
				events = append(events, e)
			case AttributeDoorLockStatusVehicle:
				e := DoorLockStatusVehicle{
					AttributeStatus: attributeStatus,
					State:           int(status.AttributeType.(*pb.VehicleAttributeStatus_IntValue).IntValue),
				}
				events = append(events, e)
			case AttributeWindowStatusOverall:
				e := WindowStatusOverall{
					AttributeStatus: attributeStatus,
					State:           WindowsOverallStatusType(status.AttributeType.(*pb.VehicleAttributeStatus_IntValue).IntValue),
				}
				events = append(events, e)
			case AttributeWindowStatusRearBlind:
				e := WindowStatusRearBlind{
					AttributeStatus: attributeStatus,
				}
				events = append(events, e)
			case AttributeWindowStatusRearLeftBlind:
				e := WindowStatusRearLeftBlind{
					AttributeStatus: attributeStatus,
				}
				events = append(events, e)
			case AttributeWindowStatusRearRightBlind:
				e := WindowStatusRearRightBlind{
					AttributeStatus: attributeStatus,
				}
				events = append(events, e)
			case AttributeWindowStatusFrontLeft:
				e := WindowStatusFrontLeft{
					AttributeStatus: attributeStatus,
					State:           WindowStatusType(status.AttributeType.(*pb.VehicleAttributeStatus_IntValue).IntValue),
				}
				events = append(events, e)
			case AttributeWindowStatusFrontRight:
				e := WindowStatusFrontRight{
					AttributeStatus: attributeStatus,
					State:           WindowStatusType(status.AttributeType.(*pb.VehicleAttributeStatus_IntValue).IntValue),
				}
				events = append(events, e)
			case AttributeWindowStatusRearLeft:
				e := WindowStatusRearLeft{
					AttributeStatus: attributeStatus,
					State:           WindowStatusType(status.AttributeType.(*pb.VehicleAttributeStatus_IntValue).IntValue),
				}
				events = append(events, e)
			case AttributeWindowStatusRearRight:
				e := WindowStatusRearRight{
					AttributeStatus: attributeStatus,
					State:           WindowStatusType(status.AttributeType.(*pb.VehicleAttributeStatus_IntValue).IntValue),
				}
				events = append(events, e)
			case AttributeSunRoofStatus:
				e := SunRoofStatus{
					AttributeStatus: attributeStatus,
					State:           WindowStatusType(status.AttributeType.(*pb.VehicleAttributeStatus_IntValue).IntValue),
				}
				events = append(events, e)
			case AttributeWarningWashWater:
				e := WarningWashWater{
					AttributeStatus: attributeStatus,
					Warning:         status.AttributeType.(*pb.VehicleAttributeStatus_BoolValue).BoolValue,
				}
				events = append(events, e)
			case AttributeWarningCoolantLevelLow:
				e := WarningCoolantLevelLow{
					AttributeStatus: attributeStatus,
					Warning:         status.AttributeType.(*pb.VehicleAttributeStatus_BoolValue).BoolValue,
				}
				events = append(events, e)
			case AttributeWarningBrakeFluid:
				e := WarningBrakeFluid{
					AttributeStatus: attributeStatus,
					Warning:         status.AttributeType.(*pb.VehicleAttributeStatus_BoolValue).BoolValue,
				}
				events = append(events, e)
			case AttributeWarningBrakeLiningWear:
				e := WarningBrakeLiningWear{
					AttributeStatus: attributeStatus,
					Warning:         status.AttributeType.(*pb.VehicleAttributeStatus_BoolValue).BoolValue,
				}
				events = append(events, e)
			case AttributeTirePressureFrontLeft:
				e := TirePressureFrontLeft{
					AttributeStatus: attributeStatus,
					Value:           status.AttributeType.(*pb.VehicleAttributeStatus_DoubleValue).DoubleValue,
				}
				events = append(events, e)
			case AttributeTirePressureFrontRight:
				e := TirePressureFrontRight{
					AttributeStatus: attributeStatus,
					Value:           status.AttributeType.(*pb.VehicleAttributeStatus_DoubleValue).DoubleValue,
				}
				events = append(events, e)
			case AttributeTirePressureRearLeft:
				e := TirePressureRearLeft{
					AttributeStatus: attributeStatus,
					Value:           status.AttributeType.(*pb.VehicleAttributeStatus_DoubleValue).DoubleValue,
				}
				events = append(events, e)
			case AttributeTirePressureRearRight:
				e := TirePressureRearRight{
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

func (c *Client) refreshAccessToken(ctx context.Context) error {
	var res struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
		ExpiresIn   int    `json:"expires_in"`
	}

	if err := requests.
		URL(RegionProviders[c.region].OAuth2URL).
		Post().
		ContentType("application/x-www-form-urlencoded").
		BodyForm(url.Values{
			"grant_type":    {"refresh_token"},
			"refresh_token": {c.refreshToken},
			"X-Device-Id":   {uuid.New().String()},
			"X-Request-Id":  {uuid.New().String()},
		}).
		ToJSON(&res).
		Fetch(ctx); err != nil {
		return err
	}

	c.accessToken = res.AccessToken
	c.expireIn = res.ExpiresIn

	return nil
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

type ClientOption interface {
	apply(*Client)
}
type clientOptionFun func(c *Client)

func (f clientOptionFun) apply(c *Client) {
	f(c)
}
