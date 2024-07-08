package main

import (
	"image/color"
	"time"

	mip "github.com/hybridgroup/tinygo-mip"
)

func actions() {
	println("chest led")
	err := robot.SetChestLED(color.RGBA{R: 255, G: 0, B: 0})
	if err != nil {
		println(err)
	}

	robot.DriveForward(50, 100)
	time.Sleep(3 * time.Second)

	robot.TurnLeft(45, 50)
	time.Sleep(3 * time.Second)

	robot.TurnRight(45, 50)
	time.Sleep(3 * time.Second)

	robot.DriveBackward(50, 100)
	time.Sleep(3 * time.Second)

	robot.SetGameMode(mip.GameModeDance)
	time.Sleep(15 * time.Second)
}
