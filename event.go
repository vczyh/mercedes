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

type MercedesEvent interface {
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

type StarterBatteryStateType uint8

const (
	StarterBatteryStateTypeGreen               StarterBatteryStateType = 0
	StarterBatteryStateTypeYellow              StarterBatteryStateType = 1
	StarterBatteryStateTypeRed                 StarterBatteryStateType = 2
	StarterBatteryStateTypeServiceDisabled     StarterBatteryStateType = 3
	StarterBatteryStateTypeVehicleNotAvailable                         = 4
)

func (t StarterBatteryStateType) String() string {
	switch t {
	case StarterBatteryStateTypeGreen:
		return "Vehicle ok"
	case StarterBatteryStateTypeYellow:
		return "Battery partly charged"
	case StarterBatteryStateTypeRed:
		return "Vehicle not available"
	case StarterBatteryStateTypeServiceDisabled:
		return "Remote service disabled"
	case StarterBatteryStateTypeVehicleNotAvailable:
		return "Vehicle no longer available"
	default:
		return fmt.Sprintf("unsupported starter battery state type: %d", t)
	}
}

type StarterBatteryState struct {
	AttributeStatus
	State StarterBatteryStateType
}

type EngineState struct {
	AttributeStatus
	Running bool
}

type DistanceReset struct {
	AttributeStatus
	Value float64
}

type AverageSpeedReset struct {
	AttributeStatus
	Value float64
}

type DrivenTimeReset struct {
	AttributeStatus
	Value time.Duration
}

type LiquidConsumptionReset struct {
	AttributeStatus
	Value float64
}

type DistanceStart struct {
	AttributeStatus
	Value float64
}

type AverageSpeedStart struct {
	AttributeStatus
	Value float64
}

type DrivenTimeStart struct {
	AttributeStatus
	Value time.Duration
}

type LiquidConsumptionStart struct {
	AttributeStatus
	Value float64
}

type Odo struct {
	AttributeStatus
	Value int
}

type OilLevel struct {
	AttributeStatus
	Level int
}

type RangeLiquid struct {
	AttributeStatus
	Value int
}

type TankLevelPercent struct {
	AttributeStatus
	Value int
}

type RoofTopStatus struct {
	AttributeStatus
}

type DoorOverallStatusType uint8

func (s DoorOverallStatusType) String() string {
	switch s {
	case DoorOverallStatusTypeOpen:
		return "open"
	case DoorOverallStatusTypeClosed:
		return "closed"
	case DoorOverallStatusTypeNotExisting:
		return "not existing"
	case DoorOverallStatusTypeUnknown:
		return "unknown"
	default:
		return fmt.Sprintf("unsupported door overall status type: %d", s)
	}
}

const (
	DoorOverallStatusTypeOpen        DoorOverallStatusType = 0
	DoorOverallStatusTypeClosed      DoorOverallStatusType = 1
	DoorOverallStatusTypeNotExisting DoorOverallStatusType = 2
	DoorOverallStatusTypeUnknown     DoorOverallStatusType = 3
)

type DoorStatusOverall struct {
	AttributeStatus
	State int
}

type DoorStatusFrontLeft struct {
	AttributeStatus
	Open bool
}

type DoorStatusFrontRight struct {
	AttributeStatus
	Open bool
}

type DoorStatusRearLeft struct {
	AttributeStatus
	Open bool
}

type DoorStatusRearRight struct {
	AttributeStatus
	Open bool
}

type DeckLidStatus struct {
	AttributeStatus
	Open bool
}

// TODO
type DoorStatusGas struct {
	AttributeStatus
	State int
}

type DoorLockStatusDeckLid struct {
	AttributeStatus
	UnLocked bool
}

type DoorLockOverallStatusType uint8

func (s DoorLockOverallStatusType) String() string {
	switch s {
	case DoorLockOverallStatusTypeLocked:
		return "locked"
	case DoorLockOverallStatusTypeUnLocked:
		return "unlocked"
	case DoorLockOverallStatusTypeNotExisting:
		return "not existing"
	case DoorLockOverallStatusTypeUnknown:
		return "unknown"
	default:
		return fmt.Sprintf("unsupported door lock overall status type: %d", s)
	}
}

const (
	DoorLockOverallStatusTypeLocked      DoorLockOverallStatusType = 0
	DoorLockOverallStatusTypeUnLocked    DoorLockOverallStatusType = 1
	DoorLockOverallStatusTypeNotExisting DoorLockOverallStatusType = 2
	DoorLockOverallStatusTypeUnknown     DoorLockOverallStatusType = 3
)

type DoorLockStatusOverall struct {
	AttributeStatus
	State DoorLockOverallStatusType
}

type DoorLockStatusFrontLeft struct {
	AttributeStatus
	UnLocked bool
}

type DoorLockStatusFrontRight struct {
	AttributeStatus
	UnLocked bool
}

type DoorLockStatusRearLeft struct {
	AttributeStatus
	UnLocked bool
}

type DoorLockStatusRearRight struct {
	AttributeStatus
	UnLocked bool
}

type DoorLockStatusGas struct {
	AttributeStatus
	UnLocked bool
}

type DoorLockStatusVehicle struct {
	AttributeStatus
	State int
}

type WindowsOverallStatusType uint8

func (s WindowsOverallStatusType) String() string {
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
	WindowsOverallStatusOpen           WindowsOverallStatusType = 0
	WindowsOverallStatusClosed         WindowsOverallStatusType = 1
	WindowsOverallStatusCompletelyOpen WindowsOverallStatusType = 2
	WindowsOverallStatusTypeAiring     WindowsOverallStatusType = 3
	WindowsOverallStatusTypeRunning    WindowsOverallStatusType = 4
)

type WindowStatusOverall struct {
	AttributeStatus
	State WindowsOverallStatusType
}

// TODO
type WindowStatusRearBlind struct {
	AttributeStatus
}

func (e WindowStatusRearBlind) mercedesEvent() {}

// TODO
type WindowStatusRearLeftBlind struct {
	AttributeStatus
}

// TODO
type WindowStatusRearRightBlind struct {
	AttributeStatus
}

type WindowStatusType uint8

func (s WindowStatusType) String() string {
	switch s {
	case WindowStatusTypeIntermediate:
		return "intermediate"
	case WindowStatusTypeOpen:
		return "open"
	case WindowStatusTypeClosed:
		return "closed"
	case WindowStatusTypeAiringPosition:
		return "airing position"
	case WindowStatusTypeAiringIntermediate:
		return "airing intermediate"
	case WindowStatusTypeRunning:
		return "running"
	default:
		return fmt.Sprintf("unsupported window status type: %d", s)
	}
}

const (
	WindowStatusTypeIntermediate       WindowStatusType = 0
	WindowStatusTypeOpen               WindowStatusType = 1
	WindowStatusTypeClosed             WindowStatusType = 2
	WindowStatusTypeAiringPosition     WindowStatusType = 3
	WindowStatusTypeAiringIntermediate WindowStatusType = 4
	WindowStatusTypeRunning            WindowStatusType = 5
)

type WindowStatusFrontLeft struct {
	AttributeStatus
	State WindowStatusType
}

type WindowStatusFrontRight struct {
	AttributeStatus
	State WindowStatusType
}

type WindowStatusRearLeft struct {
	AttributeStatus
	State WindowStatusType
}

type WindowStatusRearRight struct {
	AttributeStatus
	State WindowStatusType
}

type SunRoofStatus struct {
	AttributeStatus
	State WindowStatusType
}

type WarningWashWater struct {
	AttributeStatus
	Warning bool
}

type WarningCoolantLevelLow struct {
	AttributeStatus
	Warning bool
}

type WarningBrakeFluid struct {
	AttributeStatus
	Warning bool
}

type WarningBrakeLiningWear struct {
	AttributeStatus
	Warning bool
}

type TirePressureFrontLeft struct {
	AttributeStatus
	Value float64
}

type TirePressureFrontRight struct {
	AttributeStatus
	Value float64
}

type TirePressureRearLeft struct {
	AttributeStatus
	Value float64
}

type TirePressureRearRight struct {
	AttributeStatus
	Value float64
}
