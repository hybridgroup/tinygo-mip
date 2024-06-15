package mip

import (
	"errors"
	"image/color"

	"tinygo.org/x/bluetooth"
)

type Robot struct {
	device        *bluetooth.Device
	receive       *bluetooth.DeviceService
	receiveNotify *bluetooth.DeviceCharacteristic
	send          *bluetooth.DeviceService
	sendData      *bluetooth.DeviceCharacteristic

	buf []byte
}

var (
	// BLE services
	mipReceiveDataService              = bluetooth.New16BitUUID(0xffe0)
	mipReceiveDataNotifyCharacteristic = bluetooth.New16BitUUID(0xffe4)

	mipSendDataService             = bluetooth.New16BitUUID(0xffe5)
	mipSendDataWriteCharacteristic = bluetooth.New16BitUUID(0xffe9)
)

// NewRobot creates a new MiP robot.
func NewRobot(dev *bluetooth.Device) *Robot {
	r := &Robot{
		device: dev,
		buf:    make([]byte, 255),
	}

	return r
}

func (r *Robot) Start() (err error) {
	srvcs, err := r.device.DiscoverServices([]bluetooth.UUID{
		mipReceiveDataService,
		mipSendDataService,
	})
	if err != nil || len(srvcs) == 0 {
		return errors.New("could not find services")
	}

	r.receive = &srvcs[0]
	r.send = &srvcs[1]
	if debug {
		println("found mip receive service", r.receive.UUID().String())
		println("found mip send service", r.send.UUID().String())
	}

	chars, err := r.receive.DiscoverCharacteristics([]bluetooth.UUID{
		mipReceiveDataNotifyCharacteristic,
	})
	if err != nil || len(chars) == 0 {
		return errors.New("could not find mip receive characteristic")
	}

	r.receiveNotify = &chars[0]

	chars, err = r.send.DiscoverCharacteristics([]bluetooth.UUID{
		mipSendDataWriteCharacteristic,
	})
	if err != nil || len(chars) == 0 {
		return errors.New("could not find mip write characteristic")
	}

	r.sendData = &chars[0]

	return
}

func (r *Robot) Halt() (err error) {
	return
}

// Stops stops the MIP from moving.
func (r *Robot) Stop() (err error) {
	buf := []byte{Stop}
	_, err = r.sendData.WriteWithoutResponse(buf)

	return
}

// SetChestLED sets the chest LED of the MiP
func (r *Robot) SetChestLED(c color.RGBA) (err error) {
	buf := []byte{SetChestLED, c.R, c.G, c.B}
	_, err = r.sendData.WriteWithoutResponse(buf)

	return err
}

// FlashChestLED flashes the chest LED of the MiP
func (r *Robot) FlashChestLED(c color.RGBA, on, off int) (err error) {
	buf := []byte{FlashChestLED, c.R, c.G, c.B, byte(on), byte(off)}
	_, err = r.sendData.WriteWithoutResponse(buf)

	return err
}

// SetHeadLED sets the head LEDs of the MiP
func (r *Robot) SetHeadLED(l1, l2, l3, l4 int) (err error) {
	buf := []byte{SetHeadLED, byte(l1), byte(l2), byte(l3), byte(l4)}
	_, err = r.sendData.WriteWithoutResponse(buf)

	return err
}

// GetUp tells the MiP to stand up.
func (r *Robot) GetUp(stand int) (err error) {
	buf := []byte{GetUp, byte(stand)}
	_, err = r.sendData.WriteWithoutResponse(buf)

	return err
}

// DriveForward drives the MiP forward at a given speed for a given duration.
func (r *Robot) DriveForward(speed uint8, duration uint8) (err error) {
	buf := []byte{DriveForwardTime, speed, duration}
	_, err = r.sendData.WriteWithoutResponse(buf)

	return err
}

// DriveBackward drives the MiP backward at a given speed for a given duration.
func (r *Robot) DriveBackward(speed uint8, duration uint8) (err error) {
	buf := []byte{DriveBackwardTime, speed, duration}
	_, err = r.sendData.WriteWithoutResponse(buf)

	return err
}

// TurnLeft turns the MiP left to a given angle at at given speed.
func (r *Robot) TurnLeft(angle uint8, speed uint8) (err error) {
	buf := []byte{TurnLeftAngle, angle, speed}
	_, err = r.sendData.WriteWithoutResponse(buf)

	return
}

// TurnRight turns the MiP right to a given angle at at given speed.
func (r *Robot) TurnRight(angle uint8, speed uint8) (err error) {
	buf := []byte{TurnRightAngle, angle, speed}
	_, err = r.sendData.WriteWithoutResponse(buf)

	return
}

// SetGameMode tells the MiP to start playing a game using a [GameMode].
func (r *Robot) SetGameMode(mode GameMode) (err error) {
	buf := []byte{SetGameMode, byte(mode)}
	_, err = r.sendData.WriteWithoutResponse(buf)

	return err
}
