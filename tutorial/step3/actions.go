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

	for i := 0; i < 3; i++ {
		robot.DriveForward(30, 500)
		time.Sleep(3 * time.Second)

		robot.TurnLeft(90, 50)
		time.Sleep(3 * time.Second)
	}
}
