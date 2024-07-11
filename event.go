package mercedes

import (
	"fmt"
	"github.com/vczyh/mercedes/proto/pb"
	"time"
)

const (
	AttributeStarterBatteryState                            = "starterBatteryState"
	AttributeEngineState                                    = "engineState"
	AttributeDistanceReset                                  = "distanceReset"
	AttributeAverageSpeedReset                              = "averageSpeedReset"
	AttributeDrivenTimeReset                                = "drivenTimeReset"
	AttributeLiquidConsumptionReset                         = "liquidconsumptionreset"
	AttributeDistanceZEReset                                = "distanceZEReset"
	AttributeDrivenTimeZEReset                              = "drivenTimeZEReset"
	AttributeDistanceStart                                  = "distanceStart"
	AttributeAverageSpeedStart                              = "averageSpeedStart"
	AttributeDrivenTimeStart                                = "drivenTimeStart"
	AttributeLiquidConsumptionStart                         = "liquidconsumptionstart"
	AttributeDistanceZEStart                                = "distanceZEStart"
	AttributeDrivenTimeZEStart                              = "drivenTimeZEStart"
	AttributeOdo                                            = "odo"
	AttributeOilLevel                                       = "oilLevel"
	AttributeRangeLiquid                                    = "rangeliquid"
	AttributeTankLevelPercent                               = "tanklevelpercent"
	AttributeRoofTopStatus                                  = "rooftopstatus"
	AttributeDoorStatusOverall                              = "doorStatusOverall"
	AttributeDoorStatusFrontLeft                            = "doorstatusfrontleft"
	AttributeDoorStatusFrontRight                           = "doorstatusfrontright"
	AttributeDoorStatusRearLeft                             = "doorstatusrearleft"
	AttributeDoorStatusRearRight                            = "doorstatusrearright"
	AttributeDeckLidStatus                                  = "decklidstatus"
	AttributeDoorStatusGas                                  = "doorstatusgas"
	AttributeDoorLockStatusOverall                          = "doorLockStatusOverall"
	AttributeDoorLockStatusFrontLeft                        = "doorlockstatusfrontleft"
	AttributeDoorLockStatusFrontRight                       = "doorlockstatusfrontright"
	AttributeDoorLockStatusRearLeft                         = "doorlockstatusrearleft"
	AttributeDoorLockStatusRearRight                        = "doorlockstatusrearright"
	AttributeDoorLockStatusGas                              = "doorlockstatusgas"
	AttributeDoorLockStatusDeckLid                          = "doorlockstatusdecklid"
	AttributeDoorLockStatusVehicle                          = "doorlockstatusvehicle"
	AttributeWindowStatusOverall                            = "windowStatusOverall"
	AttributeWindowStatusRearBlind                          = "windowStatusRearBlind"
	AttributeWindowStatusRearLeftBlind                      = "windowStatusRearLeftBlind"
	AttributeWindowStatusRearRightBlind                     = "windowStatusRearRightBlind"
	AttributeWindowStatusFrontLeft                          = "windowstatusfrontleft"
	AttributeWindowStatusFrontRight                         = "windowstatusfrontright"
	AttributeWindowStatusRearLeft                           = "windowstatusrearleft"
	AttributeWindowStatusRearRight                          = "windowstatusrearright"
	AttributeSunroofStatus                                  = "sunroofstatus"
	AttributeWarningWashWater                               = "warningwashwater"
	AttributeWarningCoolantLevelLow                         = "warningcoolantlevellow"
	AttributeWarningBrakeFluid                              = "warningbrakefluid"
	AttributeWarningBrakeLiningWear                         = "warningbrakeliningwear"
	AttributeTirePressureFrontLeft                          = "tirepressureFrontLeft"
	AttributeTirePressureFrontRight                         = "tirepressureFrontRight"
	AttributeTirePressureRearLeft                           = "tirepressureRearLeft"
	AttributeTirePressureRearRight                          = "tirepressureRearRight"
	AttributeEngineHoodStatus                               = "engineHoodStatus"
	AttributeFilterParticleLoading                          = "filterParticleLoading"
	AttributeIgnitionState                                  = "ignitionstate"
	AttributeLanguageHU                                     = "languageHU"
	AttributeLiquidRangeSkipIndication                      = "liquidRangeSkipIndication"
	AttributeOverallRange                                   = "overallRange"
	AttributeParkBrakeStatus                                = "parkbrakestatus"
	AttributePositionHeading                                = "positionHeading"
	AttributePositionLat                                    = "positionLat"
	AttributePositionLong                                   = "positionLong"
	AttributeVehiclePositionErrorCode                       = "vehiclePositionErrorCode"
	AttributeProximityCalculationForVehiclePositionRequired = "proximityCalculationForVehiclePositionRequired"
	AttributeRemoteStartActive                              = "remoteStartActive"
	AttributeRemoteStartEndTime                             = "remoteStartEndtime"
	AttributeRemoteStartTemperature                         = "remoteStartTemperature"
	AttributeServiceIntervalDays                            = "serviceintervaldays"
	AttributeServiceIntervalDistance                        = "serviceintervaldistance"
	AttributeSoc                                            = "soc"
	AttributeSpeedUnitFromIC                                = "speedUnitFromIC"
	AttributeSunroofEvent                                   = "sunroofEvent"
	AttributeSunroofEventActive                             = "sunroofEventActive"
	AttributeSunroofStatusFrontBlind                        = "sunroofStatusFrontBlind"
	AttributeSunroofStatusRearBlind                         = "sunroofStatusRearBlind"
	AttributeTemperaturePoints                              = "temperaturePoints"
	AttributeTemperatureUnitHU                              = "temperatureUnitHU"
	AttributeTireMarkerFrontLeft                            = "tireMarkerFrontLeft"
	AttributeTireMarkerFrontRight                           = "tireMarkerFrontRight"
	AttributeTireMarkerRearLeft                             = "tireMarkerRearLeft"
	AttributeTireMarkerRearRight                            = "tireMarkerRearRight"
	AttributeTirePressMeasTimestamp                         = "tirePressMeasTimestamp"
	AttributeTireSensorAvailable                            = "tireSensorAvailable"
	AttributeTireWarningsRDK                                = "tirewarningsrdk"
	AttributeTireWarningLamp                                = "tirewarninglamp"
	AttributeTireWarningSprw                                = "tirewarningsprw"
	AttributeTireWarningLevelPrw                            = "tireWarningLevelPrw"
	AttributeTrackingStateHU                                = "trackingStateHU"
	AttributeVehicleDataConnectionState                     = "vehicleDataConnectionState"
	AttributeVehicleLockState                               = "vehicleLockState"
	AttributeVTime                                          = "vtime"
	AttributeWiperHealthPercent                             = "wiperHealthPercent"
	AttributeWiperLifetimeExceeded                          = "wiperLifetimeExceeded"
)

type Event interface {
	mercedesEvent()
}

type UpdateEvent struct {
	*pb.PushMessage
}

type StatusType uint8

func (s StatusType) String() string {
	switch s {
	case StatusTypeInvalid:
		return "invalid"
	case StatusTypeNotAvailable:
		return "not available"
	case StatusTypeNoValue:
		return "no value"
	case StatusTypeValid:
		return "valid"
	default:
		return fmt.Sprintf("unsupported attribute status type: %d", s)
	}
}

const (
	// StatusTypeInvalid invalid attribute data with the timestamp of last attribute change in milliseconds
	StatusTypeInvalid StatusType = 3
	// StatusTypeNotAvailable not available with the timestamp of last attribute change in milliseconds
	StatusTypeNotAvailable StatusType = 4
	// StatusTypeNoValue no value available with the timestamp of last attribute change in milliseconds
	StatusTypeNoValue StatusType = 1
	// StatusTypeValid valid attribute data with the timestamp of last attribute change in milliseconds
	StatusTypeValid StatusType = 0
)

type AttributeStatus struct {
	Vin     string
	Time    time.Time
	Status  StatusType
	Changed bool
}

func (as AttributeStatus) mercedesEvent() {}

type StarterBatteryState uint8

const (
	StarterBatteryStateGreen               StarterBatteryState = 0
	StarterBatteryStateYellow              StarterBatteryState = 1
	StarterBatteryStateRed                 StarterBatteryState = 2
	StarterBatteryStateServiceDisabled     StarterBatteryState = 3
	StarterBatteryStateVehicleNotAvailable                     = 4
)

func (t StarterBatteryState) String() string {
	switch t {
	case StarterBatteryStateGreen:
		return "Vehicle ok"
	case StarterBatteryStateYellow:
		return "Battery partly charged"
	case StarterBatteryStateRed:
		return "Vehicle not available"
	case StarterBatteryStateServiceDisabled:
		return "Remote service disabled"
	case StarterBatteryStateVehicleNotAvailable:
		return "Vehicle no longer available"
	default:
		return fmt.Sprintf("unsupported starter battery state type: %d", t)
	}
}

type StarterBatteryStateEvent struct {
	AttributeStatus
	State StarterBatteryState
}

type EngineStateEvent struct {
	AttributeStatus
	Running bool
}

type DistanceResetEvent struct {
	AttributeStatus
	Value float64
}

type AverageSpeedResetEvent struct {
	AttributeStatus
	Value float64
}

type DrivenTimeResetEvent struct {
	AttributeStatus
	Value time.Duration
}

type LiquidConsumptionEvent struct {
	AttributeStatus
	Value float64
}

type DistanceZEResetEvent struct {
	AttributeStatus
	Value float64
}

type DrivenTimeZEResetEvent struct {
	AttributeStatus
	Value time.Duration
}

type DistanceStartEvent struct {
	AttributeStatus
	Value float64
}

type AverageSpeedStartEvent struct {
	AttributeStatus
	Value float64
}

type DrivenTimeStartEvent struct {
	AttributeStatus
	Value time.Duration
}

type LiquidConsumptionStartEvent struct {
	AttributeStatus
	Value float64
}

type DistanceZEStartEvent struct {
	AttributeStatus
	Value float64
}

type DrivenTimeZEStartEvent struct {
	AttributeStatus
	Value time.Duration
}

type OdoEvent struct {
	AttributeStatus
	Value int
}

type OilLevelEvent struct {
	AttributeStatus
	Level int
}

type RangeLiquidEvent struct {
	AttributeStatus
	Value int
}

type TankLevelPercentEvent struct {
	AttributeStatus
	Value int
}

type RoofTopState uint8

func (s RoofTopState) String() string {
	switch s {
	case RoofTopStateUnlocked:
		return "unlocked"
	case RoofTopStateOpenLocked:
		return "open & locked"
	case RoofTopStateCloseLocked:
		return "closed & locked"
	default:
		return fmt.Sprintf("unsupported roof top state type: %d", s)
	}
}

const (
	RoofTopStateUnlocked    RoofTopState = 0
	RoofTopStateOpenLocked  RoofTopState = 1
	RoofTopStateCloseLocked RoofTopState = 2
)

type RoofTopStatusEvent struct {
	AttributeStatus
	State RoofTopState
}

type DoorOverallStatus uint8

func (s DoorOverallStatus) String() string {
	switch s {
	case DoorOverallStatusOpen:
		return "open"
	case DoorOverallStatusClosed:
		return "closed"
	case DoorOverallStatusNotExisting:
		return "not existing"
	case DoorOverallStatusUnknown:
		return "unknown"
	default:
		return fmt.Sprintf("unsupported door overall status type: %d", s)
	}
}

const (
	DoorOverallStatusOpen        DoorOverallStatus = 0
	DoorOverallStatusClosed      DoorOverallStatus = 1
	DoorOverallStatusNotExisting DoorOverallStatus = 2
	DoorOverallStatusUnknown     DoorOverallStatus = 3
)

type DoorStatusOverallEvent struct {
	AttributeStatus
	State DoorOverallStatus
}

type DoorStatusFrontLeftEvent struct {
	AttributeStatus
	Open bool
}

type DoorStatusFrontRightEvent struct {
	AttributeStatus
	Open bool
}

type DoorStatusRearLeftEvent struct {
	AttributeStatus
	Open bool
}

type DoorStatusRearRightEvent struct {
	AttributeStatus
	Open bool
}

type DeckLidStatusEvent struct {
	AttributeStatus
	Open bool
}

// TODO
type DoorStatusGasEvent struct {
	AttributeStatus
	State int
}

type DoorLockStatusDeckLidEvent struct {
	AttributeStatus
	UnLocked bool
}

type DoorLockOverallStatus uint8

func (s DoorLockOverallStatus) String() string {
	switch s {
	case DoorLockOverallStatusLocked:
		return "locked"
	case DoorLockOverallStatusUnLocked:
		return "unlocked"
	case DoorLockOverallStatusNotExisting:
		return "not existing"
	case DoorLockOverallStatusUnknown:
		return "unknown"
	default:
		return fmt.Sprintf("unsupported door lock overall status type: %d", s)
	}
}

const (
	DoorLockOverallStatusLocked      DoorLockOverallStatus = 0
	DoorLockOverallStatusUnLocked    DoorLockOverallStatus = 1
	DoorLockOverallStatusNotExisting DoorLockOverallStatus = 2
	DoorLockOverallStatusUnknown     DoorLockOverallStatus = 3
)

type DoorLockStatusOverallEvent struct {
	AttributeStatus
	State DoorLockOverallStatus
}

type DoorLockStatusFrontLeftEvent struct {
	AttributeStatus
	UnLocked bool
}

type DoorLockStatusFrontRightEvent struct {
	AttributeStatus
	UnLocked bool
}

type DoorLockStatusRearLeftEvent struct {
	AttributeStatus
	UnLocked bool
}

type DoorLockStatusRearRightEvent struct {
	AttributeStatus
	UnLocked bool
}

type DoorLockStatusGasEvent struct {
	AttributeStatus
	UnLocked bool
}

type DoorLockStatusVehicleEvent struct {
	AttributeStatus
	State int
}

type WindowsOverallStatus uint8

func (s WindowsOverallStatus) String() string {
	switch s {
	case WindowsOverallStatusOpen:
		return "open"
	case WindowsOverallStatusClosed:
		return "closed"
	case WindowsOverallStatusCompletelyOpen:
		return "completely open"
	case WindowsOverallStatusTypeAiring:
		return "airing"
	case WindowsOverallStatusTypeRunning:
		return "running"
	default:
		return fmt.Sprintf("unsupported windows overall status type: %d", s)
	}
}

const (
	WindowsOverallStatusOpen           WindowsOverallStatus = 0
	WindowsOverallStatusClosed         WindowsOverallStatus = 1
	WindowsOverallStatusCompletelyOpen WindowsOverallStatus = 2
	WindowsOverallStatusTypeAiring     WindowsOverallStatus = 3
	WindowsOverallStatusTypeRunning    WindowsOverallStatus = 4
)

type WindowStatusOverallEvent struct {
	AttributeStatus
	State WindowsOverallStatus
}

// TODO
type WindowStatusRearBlindEvent struct {
	AttributeStatus
}

func (e WindowStatusRearBlindEvent) mercedesEvent() {}

// TODO
type WindowStatusRearLeftBlindEvent struct {
	AttributeStatus
}

// TODO
type WindowStatusRearRightBlindEvent struct {
	AttributeStatus
}

type WindowStatus uint8

func (s WindowStatus) String() string {
	switch s {
	case WindowStatusIntermediate:
		return "intermediate"
	case WindowStatusOpen:
		return "open"
	case WindowStatusClosed:
		return "closed"
	case WindowStatusAiringPosition:
		return "airing position"
	case WindowStatusAiringIntermediate:
		return "airing intermediate"
	case WindowStatusRunning:
		return "running"
	default:
		return fmt.Sprintf("unsupported window status type: %d", s)
	}
}

const (
	WindowStatusIntermediate       WindowStatus = 0
	WindowStatusOpen               WindowStatus = 1
	WindowStatusClosed             WindowStatus = 2
	WindowStatusAiringPosition     WindowStatus = 3
	WindowStatusAiringIntermediate WindowStatus = 4
	WindowStatusRunning            WindowStatus = 5
)

type WindowStatusFrontLeftEvent struct {
	AttributeStatus
	State WindowStatus
}

type WindowStatusFrontRightEvent struct {
	AttributeStatus
	State WindowStatus
}

type WindowStatusRearLeftEvent struct {
	AttributeStatus
	State WindowStatus
}

type WindowStatusRearRightEvent struct {
	AttributeStatus
	State WindowStatus
}

type SunroofState int

func (s SunroofState) String() string {
	switch s {
	case SunroofStateClosed:
		return "closed"
	case SunroofStateOpen:
		return "open"
	case SunroofStateOpenLifting:
		return "open lifting"
	case SunroofStateRunning:
		return "running"
	case SunroofStateAntiBooming:
		return "anti-booming position"
	case SunroofStateIntermediateSliding:
		return "sliding intermediate"
	case SunroofStateIntermediateLifting:
		return "lifting intermediate"
	case SunroofStateOpening:
		return "opening"
	case SunroofStateClosing:
		return "closing"
	case SunroofStateAntiBoomingLifting:
		return "anti-booming lifting"
	case SunroofStateIntermediatePosition:
		return "intermediate position"
	case SunroofStateOpeningLifting:
		return "opening lifting"
	case SunroofStateClosingLifting:
		return "closing lifting"
	default:
		return fmt.Sprintf("unsupported sunroof status type: %d", s)
	}
}

const (
	SunroofStateClosed               SunroofState = 0
	SunroofStateOpen                 SunroofState = 1
	SunroofStateOpenLifting          SunroofState = 2
	SunroofStateRunning              SunroofState = 3
	SunroofStateAntiBooming          SunroofState = 4
	SunroofStateIntermediateSliding  SunroofState = 5
	SunroofStateIntermediateLifting  SunroofState = 6
	SunroofStateOpening              SunroofState = 7
	SunroofStateClosing              SunroofState = 8
	SunroofStateAntiBoomingLifting   SunroofState = 9
	SunroofStateIntermediatePosition SunroofState = 10
	SunroofStateOpeningLifting       SunroofState = 11
	SunroofStateClosingLifting       SunroofState = 12
)

type SunroofStatusEvent struct {
	AttributeStatus
	State SunroofState
}

type WarningWashWaterEvent struct {
	AttributeStatus
	Warning bool
}

type WarningCoolantLevelLowEvent struct {
	AttributeStatus
	Warning bool
}

type WarningBrakeFluidEvent struct {
	AttributeStatus
	Warning bool
}

type WarningBrakeLiningWearEvent struct {
	AttributeStatus
	Warning bool
}

type TirePressureFrontLeftEvent struct {
	AttributeStatus
	Value float64
}

type TirePressureFrontRightEvent struct {
	AttributeStatus
	Value float64
}

type TirePressureRearLeftEvent struct {
	AttributeStatus
	Value float64
}

type TirePressureRearRightEvent struct {
	AttributeStatus
	Value float64
}

type EngineHoodStatusEvent struct {
	AttributeStatus
	Open bool
}

type FilterParticleState uint8

func (s FilterParticleState) String() string {
	switch s {
	case FilerParticleStateHigh:
		return "Air quality high"
	case FilerParticleStateMedium:
		return "Air quality medium"
	case FilerParticleStateLow:
		return "Air quality low"
	default:
		return fmt.Sprintf("unsupported filter particle state type: %d", s)
	}
}

const (
	FilerParticleStateHigh   FilterParticleState = 0
	FilerParticleStateMedium FilterParticleState = 1
	FilerParticleStateLow    FilterParticleState = 2
)

type FilterParticleLoadingEvent struct {
	AttributeStatus
	State FilterParticleState
}

type IgnitionState uint8

func (s IgnitionState) String() string {
	switch s {
	case IgnitionStateLock:
		return "lock"
	case IgnitionStateOff:
		return "off"
	case IgnitionStateAccessory:
		return "accessory"
	case IgnitionStateOn:
		return "on"
	case IgnitionStateStart:
		return "start"
	default:
		return fmt.Sprintf("unsupported ignition state: %d", s)
	}
}

const (
	IgnitionStateLock      IgnitionState = 0
	IgnitionStateOff       IgnitionState = 1
	IgnitionStateAccessory IgnitionState = 2
	IgnitionStateOn        IgnitionState = 4
	IgnitionStateStart     IgnitionState = 5
)

type IgnitionStateEvent struct {
	AttributeStatus
	State IgnitionState
}

type LanguageState uint8

func (s LanguageState) String() string {
	switch s {
	case LanguageStateGerman:
		return "german"
	case LanguageStateEnglishImp:
		return "english imp"
	case LanguageStateFrench:
		return "french"
	case LanguageStateItalian:
		return "italian"
	case LanguageStateSpanish:
		return "spanish"
	case LanguageStateJapanese:
		return "japanese"
	case LanguageStateEnglishMet:
		return "english met"
	case LanguageStateDutch:
		return "dutch"
	case LanguageStateDanisch:
		return "danisch"
	case LanguageStateSwedish:
		return "swedish"
	case LanguageStateTurkish:
		return "turkish"
	case LanguageStatePortuguese:
		return "portuguese"
	case LanguageStateRussian:
		return "russian"
	case LanguageStateArabic:
		return "arabic"
	case LanguageStateChinese:
		return "chinese"
	case LanguageStateEnglishAm:
		return "english am"
	case LanguageStateTradChinese:
		return "chinese trad"
	case LanguageStateKorean:
		return "korean"
	case LanguageStateFinnish:
		return "finnish"
	case LanguageStatePolish:
		return "polish"
	case LanguageStateCzech:
		return "czech"
	case LanguageStateProtugueseBrazil:
		return "portuguese brazil"
	case LanguageStateNorwegian:
		return "norwegian"
	case LanguageStateThai:
		return "thai"
	case LanguageStateIndonesian:
		return "indonesian"
	case LanguageStateBulgarian:
		return "bulgarian"
	case LanguageStateSlovakian:
		return "slovakian"
	case LanguageStateCroatian:
		return "croatian"
	case LanguageStateSerbian:
		return "serbian"
	case LanguageStateHungarian:
		return "hungarian"
	case LanguageStateUkrainian:
		return "ukrainian"
	case LanguageStateMalayan:
		return "malayan"
	case LanguageStateVietnamese:
		return "vietnamese"
	case LanguageStateRomanian:
		return "romanian"
	case LanguageStateTradChineseTw:
		return "chinese trad taiwan"
	case LanguageStateHebrew:
		return "hebrew"
	case LanguageStateUnknown:
		return "unknown"
	default:
		return fmt.Sprintf("unsupported language state: %d", s)
	}
}

const (
	LanguageStateGerman           LanguageState = 0
	LanguageStateEnglishImp       LanguageState = 1
	LanguageStateFrench           LanguageState = 2
	LanguageStateItalian          LanguageState = 3
	LanguageStateSpanish          LanguageState = 4
	LanguageStateJapanese         LanguageState = 5
	LanguageStateEnglishMet       LanguageState = 6
	LanguageStateDutch            LanguageState = 7
	LanguageStateDanisch          LanguageState = 8
	LanguageStateSwedish          LanguageState = 9
	LanguageStateTurkish          LanguageState = 10
	LanguageStatePortuguese       LanguageState = 11
	LanguageStateRussian          LanguageState = 12
	LanguageStateArabic           LanguageState = 13
	LanguageStateChinese          LanguageState = 14
	LanguageStateEnglishAm        LanguageState = 15
	LanguageStateTradChinese      LanguageState = 16
	LanguageStateKorean           LanguageState = 17
	LanguageStateFinnish          LanguageState = 18
	LanguageStatePolish           LanguageState = 19
	LanguageStateCzech            LanguageState = 20
	LanguageStateProtugueseBrazil LanguageState = 21
	LanguageStateNorwegian        LanguageState = 22
	LanguageStateThai             LanguageState = 23
	LanguageStateIndonesian       LanguageState = 24
	LanguageStateBulgarian        LanguageState = 25
	LanguageStateSlovakian        LanguageState = 26
	LanguageStateCroatian         LanguageState = 27
	LanguageStateSerbian          LanguageState = 28
	LanguageStateHungarian        LanguageState = 29
	LanguageStateUkrainian        LanguageState = 30
	LanguageStateMalayan          LanguageState = 31
	LanguageStateVietnamese       LanguageState = 32
	LanguageStateRomanian         LanguageState = 33
	LanguageStateTradChineseTw    LanguageState = 34
	LanguageStateHebrew           LanguageState = 35
	LanguageStateUnknown          LanguageState = 36
)

type LanguageHUEvent struct {
	AttributeStatus
	State LanguageState
}

type LiquidRangeSkipIndicationEvent struct {
	AttributeStatus
	Skip bool
}

type OverallRangeEvent struct {
	AttributeStatus
	Value float64
}

type ParkBrakeStatusEvent struct {
	AttributeStatus
	Active bool
}

type PositionHeadingEvent struct {
	AttributeStatus
	Value float64
}

type PositionLatEvent struct {
	AttributeStatus
	Value float64
}

type PositionLongEvent struct {
	AttributeStatus
	Value float64
}

type PositionErrorState uint8

func (s PositionErrorState) String() string {
	switch s {
	case PositionErrorStateUnknown:
		return "unknown"
	case PositionErrorStateServiceInactive:
		return "service inactive"
	case PositionErrorStateTrackingInactive:
		return "tracking inactive"
	case PositionErrorStateParked:
		return "parked"
	case PositionErrorStateIgnitionOn:
		return "ignition on"
	case PositionErrorStateOk:
		return "ok"
	default:
		return fmt.Sprintf("unsupported position error state: %d", s)
	}
}

const (
	PositionErrorStateUnknown          PositionErrorState = 0
	PositionErrorStateServiceInactive  PositionErrorState = 1
	PositionErrorStateTrackingInactive PositionErrorState = 2
	PositionErrorStateParked           PositionErrorState = 3
	PositionErrorStateIgnitionOn       PositionErrorState = 4
	PositionErrorStateOk               PositionErrorState = 5
)

type PositionErrorCodeEvent struct {
	AttributeStatus
	State PositionErrorState
}

type ProximityCalculationForVehiclePositionRequiredEvent struct {
	AttributeStatus
	Required bool
}

type RemoteStartActiveEvent struct {
	AttributeStatus
	Active bool
}

type RemoteStartEndTimeEvent struct {
	AttributeStatus
	Time time.Time
}

type RemoteStartTemperatureEvent struct {
	AttributeStatus
	Value float64
}

type ServiceIntervalDaysEvent struct {
	AttributeStatus
	Value int
}

type ServiceIntervalDistanceEvent struct {
	AttributeStatus
	Value int
}

type SocEvent struct {
	AttributeStatus
	Value int
}

type SpeedUnitFromICEvent struct {
	AttributeStatus

	// KmUnit true if the unit is km, false if the unit is miles.
	KmUnit bool
}

type SunroofEventState int

func (s SunroofEventState) String() string {
	switch s {
	case SunroofEventStateNone:
		return "no event"
	case SunroofEventStateRainLiftPosition:
		return "rain lift position"
	case SunroofEventStateAutomaticLiftPosition:
		return "automatic lift position"
	case SunroofEventStateVentilationPosition:
		return "ventilation position (timer expired)"
	default:
		return fmt.Sprintf("unsupported sunroof event state: %d", s)
	}
}

const (
	SunroofEventStateNone                  SunroofEventState = 0
	SunroofEventStateRainLiftPosition      SunroofEventState = 1
	SunroofEventStateAutomaticLiftPosition SunroofEventState = 2
	SunroofEventStateVentilationPosition   SunroofEventState = 3
)

type SunroofEventEvent struct {
	AttributeStatus
	State SunroofEventState
}

type SunroofEventActiveEvent struct {
	AttributeStatus
	Active bool
}

type SunroofStatusFrontBlindEvent struct {
	AttributeStatus
	State int
}

type SunroofStatusRearBlindEvent struct {
	AttributeStatus
	State int
}

type TemperaturePointsEvent struct {
	AttributeStatus
	FrontCenter float64
	FrontLeft   float64
	FrontRight  float64
	RearCenter  float64
	Rear2Center float64
	RearLeft    float64
	RearRight   float64
}

type TemperatureUnitHUEvent struct {
	AttributeStatus

	// CelsiusUnit true if the unit is Celsius, false if the unit is Fahrenheit.
	CelsiusUnit bool
}

type TireMarkerWarningLevel int

func (l TireMarkerWarningLevel) String() string {
	switch l {
	case TireMarkerWarningNone:
		return "no warning"
	case TireMarkerWarningSoft:
		return "soft warning"
	case TireMarkerWarningLow:
		return "low warning"
	case TireMarkerWarningDeflation:
		return "deflation"
	case TireMarkerWarningMark:
		return "unknown warning"
	default:
		return fmt.Sprintf("unsupported tire marker warning: %d", l)
	}
}

const (
	TireMarkerWarningNone      TireMarkerWarningLevel = 0
	TireMarkerWarningSoft      TireMarkerWarningLevel = 1
	TireMarkerWarningLow       TireMarkerWarningLevel = 2
	TireMarkerWarningDeflation TireMarkerWarningLevel = 3
	TireMarkerWarningMark      TireMarkerWarningLevel = 4
)

type TireMarkerFrontLeftEvent struct {
	AttributeStatus
	WarningLevel TireMarkerWarningLevel
}

type TireMarkerFrontRightEvent struct {
	AttributeStatus
	WarningLevel TireMarkerWarningLevel
}

type TireMarkerRearLeftEvent struct {
	AttributeStatus
	WarningLevel TireMarkerWarningLevel
}

type TireMarkerRearRightEvent struct {
	AttributeStatus
	WarningLevel TireMarkerWarningLevel
}

type TirePressMeasTimestampEvent struct {
	AttributeStatus
	Time time.Time
}

type TireSensorState int

func (s TireSensorState) String() string {
	switch s {
	case TireSensorStateAllLocated:
		return "all sensors located"
	case TireSensorStateMissingSome:
		return "1-3 sensors are missing"
	case TireSensorStateMissingAll:
		return "all sensors missing"
	case TireSensorStateError:
		return "system error"
	default:
		return fmt.Sprintf("unsupported tire sensor state: %d", s)
	}
}

const (
	TireSensorStateAllLocated  TireSensorState = 0
	TireSensorStateMissingSome TireSensorState = 1
	TireSensorStateMissingAll  TireSensorState = 2
	TireSensorStateError       TireSensorState = 3
)

type TireSensorAvailableEvent struct {
	AttributeStatus
	State TireSensorState
}

type TireWarningLevelOverallEvent struct {
	AttributeStatus
	WarningLevel TireMarkerWarningLevel
}

type TireLampState int

func (s TireLampState) String() string {
	switch s {
	case TireLampStateInactive:
		return "inactive"
	case TireLampStateTriggered:
		return "triggered"
	case TireLampStateFlashing:
		return "flashing"
	default:
		return fmt.Sprintf("unsupported tire lamp state: %d", s)
	}
}

const (
	TireLampStateInactive  TireLampState = 0
	TireLampStateTriggered TireLampState = 1
	TireLampStateFlashing  TireLampState = 2
)

type TireWarningLampEvent struct {
	AttributeStatus
	State TireLampState
}

type TireWarningSprwEvent struct {
	AttributeStatus
	Active bool
}

type TireLevelPrwWarning int

func (l TireLevelPrwWarning) String() string {
	switch l {
	case TireLevelPrwWarningNone:
		return "no warning"
	case TireLevelPrwWarningWarning:
		return "warning"
	case TireLevelPrwWarningWorkshop:
		return "go to workshop"
	default:
		return fmt.Sprintf("unsupported tire level prw warning: %d", l)
	}
}

const (
	TireLevelPrwWarningNone     TireLevelPrwWarning = 0
	TireLevelPrwWarningWarning  TireLevelPrwWarning = 1
	TireLevelPrwWarningWorkshop TireLevelPrwWarning = 2
)

type TireWarningLevelPrwEvent struct {
	AttributeStatus
	Level TireLevelPrwWarning
}

type TrackingStateHUEvent struct {
	AttributeStatus
	Active bool
}

type VehicleDataConnectionStateEvent struct {
	AttributeStatus
	Active bool
}

type VehicleLockState int

func (s VehicleLockState) String() string {
	switch s {
	case VehicleLockStateUnlocked:
		return "unlocked"
	case VehicleLockStateLockedInternal:
		return "locked internal"
	case VehicleLockStateLockedExternal:
		return "locked external"
	case VehicleLockStateUnlockedSelective:
		return "unlocked selective"
	case VehicleLockStateUnknown:
		return "unknown"
	default:
		return fmt.Sprintf("unsupported vehicle lock state: %d", s)
	}
}

const (
	VehicleLockStateUnlocked          VehicleLockState = 0
	VehicleLockStateLockedInternal    VehicleLockState = 1
	VehicleLockStateLockedExternal    VehicleLockState = 2
	VehicleLockStateUnlockedSelective VehicleLockState = 3
	VehicleLockStateUnknown           VehicleLockState = 4
)

type VehicleLockStateEvent struct {
	AttributeStatus
	State VehicleLockState
}

type VehicleTimeEvent struct {
	AttributeStatus
	Time time.Time
}

type WiperHealthPercentEvent struct {
	AttributeStatus
	Value int
}

type WiperLifetimeExceededEvent struct {
	AttributeStatus
	Exceeded bool
}
