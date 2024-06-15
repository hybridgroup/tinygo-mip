package mip

type GameMode byte

const (
	// App - The same as cancel Gesture and Radar
	GameModeApp GameMode = 0x01

	// Cage - Play back
	GameModeCage GameMode = 0x02

	// Tracking - The same as enable Radar
	GameModeTracking GameMode = 0x03

	// Dance - Play back
	GameModeDance GameMode = 0x04

	// Default Mip Mode - The same as enable Gesture(0x0A)
	GameModeDefault GameMode = 0x05

	// Stack - Play back
	GameModeStack GameMode = 0x06

	// Trick - programming and playback
	GameModeTrick GameMode = 0x07

	// Roam Mode - Play back
	GameModeRoam GameMode = 0x08
)
