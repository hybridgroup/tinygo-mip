package mip

// all of the MiP commands
const (
	ClapDelayTime          = 0x20
	ClapEnabled            = 0x1E
	ClapTimes              = 0x1D
	ContinousDrive         = 0x78
	Detected               = 0x04
	DetectionMode          = 0x0E
	DisconnectApp          = 0xFE
	DistanceDrive          = 0x70
	DriveBackwardTime      = 0x72
	DriveForwardTime       = 0x71
	FlashChestLED          = 0x89
	ForceBLEDisconnect     = 0xFC
	GestureDetect          = 0x0A
	GetGameMode            = 0x82
	GetHardwareInfo        = 0x19
	GetRadarMode           = 0x0D
	GetSoftwareVersion     = 0x14
	GetUp                  = 0x23
	GetUserData            = 0x13
	GetVolume              = 0x16
	IRRemoteEnabled        = 0x10
	PlaySound              = 0x06
	RadarResponse          = 0x0C
	ReadOdometer           = 0x85
	ReceiveIRDongleCode    = 0x03
	RequestChestLED        = 0x83
	RequestClapEnabled     = 0x1F
	RequestDetectionMode   = 0x0F
	RequestHeadLED         = 0x8B
	RequestIRRemoteEnabled = 0x11
	RequestStatus          = 0x79
	RequestWeightUpdate    = 0x81
	ResetOdometer          = 0x86
	SendIRDongleCode       = 0x8C
	SetChestLED            = 0x84
	SetGameMode            = 0x76
	SetGestureRadarMode    = 0x0C
	SetHeadLED             = 0x8a
	SetPosition            = 0x08
	SetUserData            = 0x12
	SetVolume              = 0x15
	ShakeDetected          = 0x1A
	Sleep                  = 0xFA
	Stop                   = 0x77
	TurnLeftAngle          = 0x73
	TurnRightAngle         = 0x74
)

// HeadLED states
type HeadLED byte

const (
	HeadLEDOff HeadLED = iota
	HeadLEDOn
	HeadLEDBlinkSlow
	HeadLEDBlinkFast
)

// GetUpMode states
type GetUpMode byte

const (
	GetUpModeBack GetUpMode = iota
	GetUpModeFront
	GetUpModeAny
)
