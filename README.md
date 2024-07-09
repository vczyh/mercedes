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

| Event                           | Type              | Description |
|---------------------------------|-------------------|-------------|
| StarterBatteryStateEvent        |                   |             |
| EngineStateEvent                | Engine            |             |
| DistanceResetEvent              | `Reset statistic` |             |
| AverageSpeedResetEvent          | `Reset Statistic`   |             |
| DrivenTimeResetEvent            | `Reset Statistic`   |             |
| LiquidConsumptionEvent          | `Reset Statistic`   |             |
| DistanceStartEvent              | Start Statistic   |             |
| AverageSpeedStartEvent          | Start Statistic   |             |
| DrivenTimeStartEvent            | Start Statistic   |             |
| LiquidConsumptionStartEvent     | Start Statistic   |             |
| OdoEvent                        | Vehicle           |             |
| OilLevelEvent                   |                   |             |
| RangeLiquidEvent                |                   |             |
| TankLevelPercentEvent           |                   |             |
| RoofTopStatusEvent              | Vehicle           |             |
| DoorStatusOverallEvent          | Door              |             |
| DoorStatusFrontLeftEvent        | Door              |             |
| DoorStatusFrontRightEvent       | Door              |             |
| DoorStatusRearLeftEvent         | Door              |             |
| DoorStatusRearRightEvent        | Door              |             |
| DeckLidStatusEvent              | Door              |             |
| DoorStatusGasEvent              |                   |             |
| DoorLockStatusOverallEvent      | Door Lock         |             |
| DoorLockStatusFrontLeftEvent    | Door Lock         |             |
| DoorLockStatusFrontRightEvent   | Door Lock         |             |
| DoorLockStatusRearLeftEvent     | Door Lock         |             |
| DoorLockStatusRearRightEvent    | Door Lock         |             |
| DoorLockStatusDeckLidEvent      | Door Lock         |             |
| DoorLockStatusGasEvent          | Vehicle           |             |
| DoorLockStatusVehicleEvent      |                   |             |
| WindowStatusOverallEvent        | Window            |             |
| WindowStatusRearBlindEvent      |                   |             |
| WindowStatusRearLeftBlindEvent  |                   |             |
| WindowStatusRearRightBlindEvent |                   |             |
| WindowStatusFrontLeftEvent      | Window            |             |
| WindowStatusFrontRightEvent     | Window            |             |
| WindowStatusRearLeftEvent       | Window            |             |
| WindowStatusRearRightEvent      | Window            |             |
| SunRoofStatusEvent              | Window            |             |
| WarningWashWaterEvent           |                   |             |
| WarningCoolantLevelLowEvent     |                   |             |
| WarningBrakeFluidEvent          |                   |             |
| WarningBrakeLiningWearEvent     |                   |             |
| TirePressureFrontLeftEvent      | Tire              |             |
| TirePressureFrontRightEvent     | Tire              |             |
| TirePressureRearLeftEvent       | Tire              |             |
| TirePressureRearRightEvent      | Tire              |             |

## Thanks

- [ReneNulschDE/mbapi2020](https://github.com/ReneNulschDE/mbapi2020)
- [mercedes-benz/MBSDK-CarKit-iOS](https://github.com/mercedes-benz/MBSDK-CarKit-iOS)