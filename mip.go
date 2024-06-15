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
	println("found mip receive service", r.receive.UUID().String())
	println("found mip send service", r.send.UUID().String())

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

// Stops stops the MIP.
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
