package mercedes

import (
	"fmt"
	"github.com/vczyh/mercedes/proto/pb"
	"time"
)

const (
	AttributeStarterBatteryState        = "starterBatteryState"
	AttributeEngineState                = "engineState"
	AttributeDistanceReset              = "distanceReset"
	AttributeAverageSpeedReset          = "averageSpeedReset"
	AttributeDrivenTimeReset            = "drivenTimeReset"
	AttributeLiquidConsumptionReset     = "liquidconsumptionreset"
	AttributeDistanceStart              = "distanceStart"
	AttributeAverageSpeedStart          = "averageSpeedStart"
	AttributeDrivenTimeStart            = "drivenTimeStart"
	AttributeLiquidConsumptionStart     = "liquidconsumptionstart"
	AttributeOdo                        = "odo"
	AttributeOilLevel                   = "oilLevel"
	AttributeRangeLiquid                = "rangeliquid"
	AttributeTankLevelPercent           = "tanklevelpercent"
	AttributeRoofTopStatus              = "rooftopstatus"
	AttributeDoorStatusOverall          = "doorStatusOverall"
	AttributeDoorStatusFrontLeft        = "doorstatusfrontleft"
	AttributeDoorStatusFrontRight       = "doorstatusfrontright"
	AttributeDoorStatusRearLeft         = "doorstatusrearleft"
	AttributeDoorStatusRearRight        = "doorstatusrearright"
	AttributeDeckLidStatus              = "decklidstatus"
	AttributeDoorStatusGas              = "doorstatusgas"
	AttributeDoorLockStatusOverall      = "doorLockStatusOverall"
	AttributeDoorLockStatusFrontLeft    = "doorlockstatusfrontleft"
	AttributeDoorLockStatusFrontRight   = "doorlockstatusfrontright"
	AttributeDoorLockStatusRearLeft     = "doorlockstatusrearleft"
	AttributeDoorLockStatusRearRight    = "doorlockstatusrearright"
	AttributeDoorLockStatusGas          = "doorlockstatusgas"
	AttributeDoorLockStatusDeckLid      = "doorlockstatusdecklid"
	AttributeDoorLockStatusVehicle      = "doorlockstatusvehicle"
	AttributeWindowStatusOverall        = "windowStatusOverall"
	AttributeWindowStatusRearBlind      = "windowStatusRearBlind"
	AttributeWindowStatusRearLeftBlind  = "windowStatusRearLeftBlind"
	AttributeWindowStatusRearRightBlind = "windowStatusRearRightBlind"
	AttributeWindowStatusFrontLeft      = "windowstatusfrontleft"
	AttributeWindowStatusFrontRight     = "windowstatusfrontright"
	AttributeWindowStatusRearLeft       = "windowstatusrearleft"
	AttributeWindowStatusRearRight      = "windowstatusrearright"
	AttributeSunRoofStatus              = "sunroofstatus"
	AttributeWarningWashWater           = "warningwashwater"
	AttributeWarningCoolantLevelLow     = "warningcoolantlevellow"
	AttributeWarningBrakeFluid          = "warningbrakefluid"
	AttributeWarningBrakeLiningWear     = "warningbrakeliningwear"
	AttributeTirePressureFrontLeft      = "tirepressureFrontLeft"
	AttributeTirePressureFrontRight     = "tirepressureFrontRight"
	AttributeTirePressureRearLeft       = "tirepressureRearLeft"
	AttributeTirePressureRearRight      = "tirepressureRearRight"
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
	Vin           string
	TimestampInMs int64
	Status        StatusType
	Changed       bool
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
	State int
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

type SunRoofStatus uint8

func (s SunRoofStatus) String() string {
	switch s {
	case SunRoofStatusClosed:
		return "closed"
	case SunRoofStatusOpen:
		return "open"
	case SunRoofStatusOpenLifting:
		return "open lifting"
	case SunRoofStatusRunning:
		return "running"
	case SunRoofStatusAntiBooming:
		return "anti-booming position"
	case SunRoofStatusIntermediateSliding:
		return "sliding intermediate"
	case SunRoofStatusIntermediateLifting:
		return "lifting intermediate"
	case SunRoofStatusOpening:
		return "opening"
	case SunRoofStatusClosing:
		return "closing"
	case SunRoofStatusAntiBoomingLifting:
		return "anti-booming lifting"
	case SunRoofStatusIntermediatePosition:
		return "intermediate position"
	case SunRoofStatusOpeningLifting:
		return "opening lifting"
	case SunRoofStatusClosingLifting:
		return "closing lifting"
	default:
		return fmt.Sprintf("unsupported sunroof status type: %d", s)
	}
}

const (
	SunRoofStatusClosed               SunRoofStatus = 0
	SunRoofStatusOpen                 SunRoofStatus = 1
	SunRoofStatusOpenLifting          SunRoofStatus = 2
	SunRoofStatusRunning              SunRoofStatus = 3
	SunRoofStatusAntiBooming          SunRoofStatus = 4
	SunRoofStatusIntermediateSliding  SunRoofStatus = 5
	SunRoofStatusIntermediateLifting  SunRoofStatus = 6
	SunRoofStatusOpening              SunRoofStatus = 7
	SunRoofStatusClosing              SunRoofStatus = 8
	SunRoofStatusAntiBoomingLifting   SunRoofStatus = 9
	SunRoofStatusIntermediatePosition SunRoofStatus = 10
	SunRoofStatusOpeningLifting       SunRoofStatus = 11
	SunRoofStatusClosingLifting       SunRoofStatus = 12
)

type SunRoofStatusEvent struct {
	AttributeStatus
	State SunRoofStatus
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
