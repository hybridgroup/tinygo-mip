package main

import (
	"image/color"
	"time"

	mip "github.com/hybridgroup/tinygo-mip"
)

func actions() {
	err := robot.SetChestLED(color.RGBA{R: 0, G: 255, B: 255})
	if err != nil {
		println(err)
	}

	robot.SetHeadLED(mip.HeadLEDOn, mip.HeadLEDOn, mip.HeadLEDOn, mip.HeadLEDOn)

	robot.TurnLeft(45, 50)
	time.Sleep(3 * time.Second)

	robot.TurnRight(90, 50)
	time.Sleep(3 * time.Second)

	robot.TurnLeft(45, 50)
	time.Sleep(3 * time.Second)

	robot.SetGameMode(mip.GameModeDance)
	time.Sleep(15 * time.Second)
}
