package main

import (
	"image/color"
	"time"
)

func actions() {
	err := robot.SetChestLED(color.RGBA{R: 255, G: 0, B: 0})
	if err != nil {
		println(err)
	}

	robot.DriveForward(30, 500)
	time.Sleep(3 * time.Second)

	robot.DriveBackward(30, 500)
	time.Sleep(3 * time.Second)
}
