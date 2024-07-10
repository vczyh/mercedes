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

| Event                                               | Type        | Description                           |
|-----------------------------------------------------|-------------|---------------------------------------|
| StarterBatteryStateEvent                            | `Vehicle`   | Starter battery state                 |
| DoorLockStatusGasEvent                              | `Vehicle`   | Lock status of the gas cap door       |
| RoofTopStatusEvent                                  | `Vehicle`   | Status of the rooftop                 |
| OdoEvent                                            | `Vehicle `  | Odometer reading                      |
| EngineHoodStatusEvent                               | `Vehicle`   | Status of engine hood                 |
| FilterParticleLoadingEvent                          | `Vehicle`   |                                       |
| ParkBrakeStatusEvent                                | `Vehicle`   |                                       |
| ServiceIntervalDaysEvent                            | `Vehicle`   |                                       |
| ServiceIntervalDistanceEvent                        | `Vehicle`   |                                       |
| SocEvent                                            | `Vehicle`   |                                       |
| SpeedUnitFromICEvent                                | `Vehicle`   |                                       |
| VehicleDataConnectionStateEvent                     | `Vehicle`   |                                       |
| VehicleLockStateEvent                               | `Vehicle`   |                                       |
| VehicleTimeEvent                                    | `Vehicle`   |                                       |
| WiperHealthPercentEvent                             | `Vehicle`   |                                       |
| WiperLifetimeExceededEvent                          | `Vehicle`   |                                       |
| EngineStateEvent                                    | `Engine`    | Engine state                          |
| IgnitionStateEvent                                  | `Engine`    |                                       |
| RemoteStartActiveEvent                              | `Engine`    |                                       |
| RemoteStartEndTimeEvent                             | `Engine`    |                                       |
| RemoteStartTemperatureEvent                         | `Engine`    |                                       |
| DistanceResetEvent                                  | `Statistic` | Distance moved after reset            |
| AverageSpeedResetEvent                              | `Statistic` | Average speed after rest              |
| DrivenTimeResetEvent                                | `Statistic` | Driving time after reset              |
| LiquidConsumptionEvent                              | `Statistic` | Liquid consumption after rest         |
| DistanceZEResetEvent                                | `Statistic` |                                       |
| DistanceStartEvent                                  | `Statistic` | Distance moved after start            |
| AverageSpeedStartEvent                              | `Statistic` | Average speed after start             |
| DrivenTimeStartEvent                                | `Statistic` | Driving time after start              |
| LiquidConsumptionStartEvent                         | `Statistic` | Liquid consumption after start        |
| DistanceZEStartEvent                                | `Statistic` |                                       |
| OilLevelEvent                                       |             | Oil level                             |
| TankLevelPercentEvent                               | `Tank`      | Percentage of liquid in the tank      |
| RangeLiquidEvent                                    | `Tank`      | Range available with current liquid   |
| DoorStatusOverallEvent                              | `Door`      | Overall door status                   |
| DoorStatusFrontLeftEvent                            | `Door`      | Status of the front left door         |
| DoorStatusFrontRightEvent                           | `Door`      | Status of the front right door        |
| DoorStatusRearLeftEvent                             | `Door`      | Status of the rear left door          |
| DoorStatusRearRightEvent                            | `Door`      | Status of the rear right door         |
| DeckLidStatusEvent                                  | `Door`      | Status of the deck lid                |
| DoorStatusGasEvent                                  |             | Status of the gas cap door            |
| DoorLockStatusOverallEvent                          | `Door`      | Overall door lock status              |
| DoorLockStatusFrontLeftEvent                        | `Door`      | Lock status of the front left door    |
| DoorLockStatusFrontRightEvent                       | `Door`      | Lock status of the front right door   |
| DoorLockStatusRearLeftEvent                         | `Door`      | Lock status of the rear left door     |
| DoorLockStatusRearRightEvent                        | `Door`      | Lock status of the rear right door    |
| DoorLockStatusDeckLidEvent                          | `Door`      | Lock status of the deck lid           |
| DoorLockStatusVehicleEvent                          |             | Overall lock status of the vehicle    |
| WindowStatusRearBlindEvent                          | `Blind`     | Status of the rear blind              |
| WindowStatusRearLeftBlindEvent                      | `Blind`     | Status of the rear left blind         |
| WindowStatusRearRightBlindEvent                     | `Blind`     | Status of the rear right blind        |
| WindowStatusOverallEvent                            | `Window`    | Overall window status                 |
| WindowStatusFrontLeftEvent                          | `Window`    | Status of the front left window       |
| WindowStatusFrontRightEvent                         | `Window`    | Status of the front right window      |
| WindowStatusRearLeftEvent                           | `Window`    | Status of the rear left window        |
| WindowStatusRearRightEvent                          | `Window`    | Status of the rear right window       |
| SunRoofStatusEvent                                  | `Window`    | Status of the sunroof                 |
| SunroofStatusFrontBlindEvent                        | `Window`    |                                       |
| SunroofStatusRearBlindEvent                         | `Window`    |                                       |
| SunroofEventEvent                                   | `Window`    |                                       |
| SunroofEventActiveEvent                             | `Window`    |                                       |
| WarningWashWaterEvent                               | `Warning`   | Warning for low washer fluid level    |
| WarningCoolantLevelLowEvent                         | `Warning`   | Warning for low coolant level         |
| WarningBrakeFluidEvent                              | `Warning`   | Warning for low brake fluid level     |
| WarningBrakeLiningWearEvent                         | `Warning`   | Warning for brake lining wear         |
| LiquidRangeSkipIndicationEvent                      | `Warning`   | Warning for engine light              |
| TireWarningLampEvent                                | `Warning`   |                                       |
| TireWarningSprwEvent                                | `Warning`   |                                       |
| TireWarningLevelPrwEvent                            | `Warning`   |                                       |
| TirePressureFrontLeftEvent                          | `Tire`      | Tire pressure of the front left tire  |
| TirePressureFrontRightEvent                         | `Tire`      | Tire pressure of the front right tire |
| TirePressureRearLeftEvent                           | `Tire`      | Tire pressure of the rear left tire   |                                     |
| TirePressureRearRightEvent                          | `Tire`      | Tire pressure of the rear right tire  |
| TireMarkerFrontLeftEvent                            | `Tire`      |                                       |
| TireMarkerFrontRightEvent                           | `Tire`      |                                       |
| TireMarkerRearLeftEvent                             | `Tire`      |                                       |
| TireMarkerRearRightEvent                            | `Tire`      |                                       |
| TireWarningLevelOverallEvent                        | `Tire`      |                                       |
| TireSensorAvailableEvent                            | `Tire`      |                                       |
| TirePressMeasTimestampEvent                         | `Tire`      |                                       |
| PositionHeadingEvent                                | `Location`  |                                       |                                       |
| PositionLatEvent                                    | `Location`  |                                       |
| PositionLongEvent                                   | `Location`  |                                       |
| PositionErrorCodeEvent                              | `Location`  |                                       |
| ProximityCalculationForVehiclePositionRequiredEvent | `Location`  |                                       |
| LanguageHUEvent                                     | `HU`        |                                       |
| TemperatureUnitHUEvent                              | `HU`        |                                       |
| TrackingStateHUEvent                                | `HU`        |                                       |

## Thanks

- [ReneNulschDE/mbapi2020](https://github.com/ReneNulschDE/mbapi2020)
- [mercedes-benz/MBSDK-CarKit-iOS](https://github.com/mercedes-benz/MBSDK-CarKit-iOS)