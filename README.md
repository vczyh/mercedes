## Mercedes SDK

- Receive push messages from server.
- Control your car by sending command to server. (developing)

## Usage

```go
var f mercedes.EventListenFun = func (event mercedes.Event, err error) {
    if err != nil {
        return
    }
    fmt.Printf("%s %T %+v\n", time.Now(), event, event)
}

c := mercedes.New(
    mercedes.WithAccessToken("access_token"),
    mercedes.WithRefreshToken("refresh_token"),
    mercedes.WithEventListen(f),
)
ctx := context.TODO()

_ = c.Connect(ctx)
defer c.Close()

_ = c.DoorsUnLock(ctx, "vin", "pin")
```

## Get Token

```shell
go install github.com/vczyh/mercedes/cmd/mercedes
```

Then login and get token:

```shell
mercedes --debug login <email>
```

## Events

| Event                           | Type | Description |
|---------------------------------|------|-------------|
| StarterBatteryStateEvent        |      |             |
| EngineStateEvent                |      |             |
| DistanceResetEvent              |      |             |
| AverageSpeedResetEvent          |      |             |
| DrivenTimeResetEvent            |      |             |
| LiquidConsumptionEvent          |      |             |
| DistanceStartEvent              |      |             |
| AverageSpeedStartEvent          |      |             |
| DrivenTimeStartEvent            |      |             |
| LiquidConsumptionStartEvent     |      |             |
| OdoEvent                        |      |             |
| OilLevelEvent                   |      |             |
| RangeLiquidEvent                |      |             |
| TankLevelPercentEvent           |      |             |
| RoofTopStatusEvent              |      |             |
| DoorStatusOverallEvent          |      |             |
| DoorStatusFrontLeftEvent        |      |             |
| DoorStatusFrontRightEvent       |      |             |
| DoorStatusRearLeftEvent         |      |             |
| DoorStatusRearRightEvent        |      |             |
| DeckLidStatusEvent              |      |             |
| DoorStatusGasEvent              |      |             |
| DoorLockStatusOverallEvent      |      |             |
| DoorLockStatusFrontLeftEvent    |      |             |
| DoorLockStatusFrontRightEvent   |      |             |
| DoorLockStatusRearLeftEvent     |      |             |
| DoorLockStatusRearRightEvent    |      |             |
| DoorLockStatusGasEvent          |      |             |
| DoorLockStatusDeckLidEvent      |      |             |
| DoorLockStatusVehicleEvent      |      |             |
| WindowStatusOverallEvent        |      |             |
| WindowStatusRearBlindEvent      |      |             |
| WindowStatusRearLeftBlindEvent  |      |             |
| WindowStatusRearRightBlindEvent |      |             |
| WindowStatusFrontLeftEvent      |      |             |
| WindowStatusFrontRightEvent     |      |             |
| WindowStatusRearLeftEvent       |      |             |
| WindowStatusRearRightEvent      |      |             |
| SunRoofStatusEvent              |      |             |
| WarningWashWaterEvent           |      |             |
| WarningCoolantLevelLowEvent     |      |             |
| WarningBrakeFluidEvent          |      |             |
| WarningBrakeLiningWearEvent     |      |             |
| TirePressureFrontLeftEvent      |      |             |
| TirePressureFrontRightEvent     |      |             |
| TirePressureRearLeftEvent       |      |             |
| TirePressureRearRightEvent      |      |             |

## Thanks

- [ReneNulschDE/mbapi2020](https://github.com/ReneNulschDE/mbapi2020)
- [mercedes-benz/MBSDK-CarKit-iOS](https://github.com/mercedes-benz/MBSDK-CarKit-iOS)