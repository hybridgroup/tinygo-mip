package main

import (
	"image/color"
	"time"
)

func actions() {
	println("chest led")
	err := robot.SetChestLED(color.RGBA{R: 255, G: 0, B: 0})
	if err != nil {
		println(err)
	}

	robot.DriveForward(30, 50)
	time.Sleep(3 * time.Second)

	robot.TurnLeft(45, 50)
	time.Sleep(3 * time.Second)

	robot.TurnRight(45, 50)
	time.Sleep(3 * time.Second)

	robot.DriveBackward(30, 50)
	time.Sleep(3 * time.Second)
}
