package main

import (
	"context"
	"fmt"
	"image/color"
	"log"

	mip "github.com/hybridgroup/tinygo-mip"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

var s *server.MCPServer
var httpSrv *server.StreamableHTTPServer

func startMCP() {
	s = server.NewMCPServer(
		"TinyGo MIP",
		"1.0.0",
		server.WithToolCapabilities(true),
	)

	addToolChestLED()
	addToolFlashChestLED()
	addToolHeadLED()
	addToolGetUp()
	addToolDriveForward()
	addToolDriveBackward()
	addToolTurnLeft()
	addToolTurnRight()
	addToolPlaySound()

	httpServer := server.NewStreamableHTTPServer(s)
	log.Printf("HTTP server listening on %s", port)
	if err := httpServer.Start(port); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}

func addToolChestLED() {
	tool := mcp.NewTool("set_chest_led",
		mcp.WithDescription("Set the chest RGB LED to a specific color"),
		mcp.WithNumber("red",
			mcp.Description("Red byte"),
			mcp.Required(),
		),
		mcp.WithNumber("green",
			mcp.Description("Green byte"),
			mcp.Required(),
		),
		mcp.WithNumber("blue",
			mcp.Description("Blue byte"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, chestLEDToolHandler)
}

func chestLEDToolHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	red, err := request.RequireInt("red")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	green, err := request.RequireInt("green")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	blue, err := request.RequireInt("blue")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	if robot == nil {
		return mcp.NewToolResultError("MIP robot not available"), nil
	}

	err = robot.SetChestLED(color.RGBA{R: uint8(red), G: uint8(green), B: uint8(blue)})
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return mcp.NewToolResultText(fmt.Sprintf("chest LED set to RGB %d, %d, %d", red, green, blue)), nil
}

func addToolFlashChestLED() {
	tool := mcp.NewTool("flash_chest_led",
		mcp.WithDescription("Flash the chest RGB LED on and off"),
		mcp.WithNumber("red",
			mcp.Description("Red byte"),
			mcp.Required(),
		),
		mcp.WithNumber("green",
			mcp.Description("Green byte"),
			mcp.Required(),
		),
		mcp.WithNumber("blue",
			mcp.Description("Blue byte"),
			mcp.Required(),
		),
		mcp.WithNumber("on",
			mcp.Description("How many times to turn on"),
			mcp.Required(),
		),
		mcp.WithNumber("off",
			mcp.Description("How many times to turn off"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, flashChestLEDToolHandler)
}

func flashChestLEDToolHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	red, err := request.RequireInt("red")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	green, err := request.RequireInt("green")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	blue, err := request.RequireInt("blue")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	on, err := request.RequireInt("on")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	off, err := request.RequireInt("off")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	if robot == nil {
		return mcp.NewToolResultError("MIP robot not available"), nil
	}

	err = robot.FlashChestLED(color.RGBA{R: uint8(red), G: uint8(green), B: uint8(blue)}, on, off)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return mcp.NewToolResultText(fmt.Sprintf("flash LED set to RGB %d, %d, %d with on %d/off %d", red, green, blue, on, off)), nil
}

func addToolHeadLED() {
	tool := mcp.NewTool("set_head_led",
		mcp.WithDescription("Set the Head LEDs on or off. 0=off, 1=on, 2=blink slow, 3=blink fast."),
		mcp.WithNumber("l1",
			mcp.Description("light 1, outside left of the head"),
			mcp.Required(),
		),
		mcp.WithNumber("l2",
			mcp.Description("light 2, middle left of the head"),
			mcp.Required(),
		),
		mcp.WithNumber("l3",
			mcp.Description("light 3, middle right of the head"),
			mcp.Required(),
		),
		mcp.WithNumber("l4",
			mcp.Description("light 4, outside right of the head"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, headLEDToolHandler)
}

func headLEDToolHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	l1, err := request.RequireInt("l1")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	l2, err := request.RequireInt("l2")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	l3, err := request.RequireInt("l3")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	l4, err := request.RequireInt("l4")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	if robot == nil {
		return mcp.NewToolResultError("MIP robot not available"), nil
	}

	err = robot.SetHeadLED(mip.HeadLED(l1), mip.HeadLED(l2), mip.HeadLED(l3), mip.HeadLED(l4))
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return mcp.NewToolResultText(fmt.Sprintf("Head LED set to %d, %d, %d, %d", l1, l2, l3, l4)), nil
}

func addToolGetUp() {
	tool := mcp.NewTool("get_up",
		mcp.WithDescription("Makes the MiP robot stand up. "),
		mcp.WithNumber("mode",
			mcp.Description("Get up mode. 0x02 to get up when mip has fallen either back or front"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getUpToolHandler)
}

func getUpToolHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	mode, err := request.RequireInt("mode")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	if robot == nil {
		return mcp.NewToolResultError("MIP robot not available"), nil
	}

	err = robot.GetUp(mip.GetUpMode(mode))
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return mcp.NewToolResultText(fmt.Sprintf("getting up with mode %d", mode)), nil
}

func addToolDriveForward() {
	tool := mcp.NewTool("drive_forward",
		mcp.WithDescription("Drive forward"),
		mcp.WithNumber("speed",
			mcp.Description("Speed (0~30)"),
			mcp.Required(),
		),
		mcp.WithNumber("duration",
			mcp.Description("Time in 7ms intervals (0~255)"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, driveForwardToolHandler)
}

func driveForwardToolHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	speed, err := request.RequireInt("speed")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	duration, err := request.RequireInt("duration")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	if robot == nil {
		return mcp.NewToolResultError("MIP robot not available"), nil
	}

	err = robot.DriveForward(speed, duration)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return mcp.NewToolResultText(fmt.Sprintf("driving forward at speed %d, duration %d", speed, duration)), nil
}

func addToolDriveBackward() {
	tool := mcp.NewTool("drive_backward",
		mcp.WithDescription("Drive backward"),
		mcp.WithNumber("speed",
			mcp.Description("Speed (0~30)"),
			mcp.Required(),
		),
		mcp.WithNumber("duration",
			mcp.Description("Time in 7ms intervals (0~255)"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, driveBackwardToolHandler)
}

func driveBackwardToolHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	speed, err := request.RequireInt("speed")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	duration, err := request.RequireInt("duration")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	if robot == nil {
		return mcp.NewToolResultError("MIP robot not available"), nil
	}

	err = robot.DriveBackward(speed, duration)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return mcp.NewToolResultText(fmt.Sprintf("driving backward at speed %d, duration %d", speed, duration)), nil
}

func addToolTurnLeft() {
	tool := mcp.NewTool("turn_left",
		mcp.WithDescription("Turn left"),
		mcp.WithNumber("angle",
			mcp.Description("Angle in intervals of 5 degrees (0~255)"),
			mcp.Required(),
		),
		mcp.WithNumber("speed",
			mcp.Description("Speed (0~24)"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, turnLeftToolHandler)
}

func turnLeftToolHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	angle, err := request.RequireInt("angle")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	speed, err := request.RequireInt("speed")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	if robot == nil {
		return mcp.NewToolResultError("MIP robot not available"), nil
	}

	err = robot.TurnLeft(angle, speed)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return mcp.NewToolResultText(fmt.Sprintf("turning left at angle %d, speed %d", angle, speed)), nil
}

func addToolTurnRight() {
	tool := mcp.NewTool("turn_right",
		mcp.WithDescription("Turn right"),
		mcp.WithNumber("angle",
			mcp.Description("Angle in intervals of 5 degrees (0~255)"),
			mcp.Required(),
		),
		mcp.WithNumber("speed",
			mcp.Description("Speed (0~24)"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, turnRightToolHandler)
}

func turnRightToolHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	angle, err := request.RequireInt("angle")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	speed, err := request.RequireInt("speed")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	if robot == nil {
		return mcp.NewToolResultError("MIP robot not available"), nil
	}

	err = robot.TurnRight(angle, speed)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return mcp.NewToolResultText(fmt.Sprintf("turning right at angle %d, speed %d", angle, speed)), nil
}

func addToolPlaySound() {
	tool := mcp.NewTool("play_sound",
		mcp.WithDescription("Play a sound from the built-in sound library. You must wait for the sound to finish playing before trying to play another sound."),
		mcp.WithNumber("sound",
			mcp.Description("Sound from the sound library to play between 1 and 104. Send 105 to stop playing. Send between 0xF7 and 0xFE to set volume"),
			mcp.Required(),
		),
		mcp.WithNumber("delay",
			mcp.Description("Delay in intervals of 30ms (0~255)"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, playSoundToolHandler)
}

func playSoundToolHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	sound, err := request.RequireInt("sound")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	delay, err := request.RequireInt("delay")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	if robot == nil {
		return mcp.NewToolResultError("MIP robot not available"), nil
	}

	err = robot.PlaySound(mip.Sound(sound), delay)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return mcp.NewToolResultText(fmt.Sprintf("playing sound %d, delay %d", sound, delay)), nil
}
