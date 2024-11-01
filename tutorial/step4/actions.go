package main

import (
	"image/color"
	"time"

	mip "github.com/hybridgroup/tinygo-mip"
)

func actions() {
	err := robot.SetChestLED(color.RGBA{R: 255, G: 0, B: 0})
	if err != nil {
		println(err)
	}

	robot.DriveForward(50, 500)
	time.Sleep(3 * time.Second)

	robot.SetHeadLED(mip.HeadLEDOn, mip.HeadLEDOn, mip.HeadLEDOn, mip.HeadLEDOn)

	robot.TurnLeft(90, 50)
	time.Sleep(3 * time.Second)

	robot.TurnRight(180, 50)
	time.Sleep(3 * time.Second)

	robot.TurnLeft(90, 50)
	time.Sleep(3 * time.Second)

	robot.DriveBackward(50, 500)
	time.Sleep(3 * time.Second)

	robot.PlaySound(mip.SoundOhhYeah, 0)
}
