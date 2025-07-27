package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"

	mip "github.com/hybridgroup/tinygo-mip"
	"tinygo.org/x/bluetooth"
)

var (
	adapter = bluetooth.DefaultAdapter
	device  bluetooth.Device
	ch      = make(chan bluetooth.ScanResult, 1)

	robot *mip.Robot
	port  string
)

func main() {
	port = *flag.String("port", ":9090", "MCP server port")
	flag.Parse()

	time.Sleep(5 * time.Second)
	println("enabling...")

	must("enable BLE interface", adapter.Enable())

	println("start scan...")

	must("start scan", adapter.Scan(scanHandler))

	var err error
	select {
	case result := <-ch:
		device, err = adapter.Connect(result.Address, bluetooth.ConnectionParams{})
		must("connect to peripheral device", err)

		println("connected to ", result.Address.String())
	}

	robot = mip.NewRobot(&device)
	err = robot.Start()
	if err != nil {
		println(err)
	}

	time.Sleep(3 * time.Second)

	startMCP()
}

func scanHandler(a *bluetooth.Adapter, d bluetooth.ScanResult) {
	println("device:", d.Address.String(), d.RSSI, d.LocalName())
	if d.Address.String() == connectAddress() {
		a.StopScan()
		ch <- d
	}
}

func must(action string, err error) {
	if err != nil {
		for {
			println("failed to " + action + ": " + err.Error())
			time.Sleep(time.Second)
		}
	}
}

func init() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		if httpSrv != nil {
			httpSrv.Shutdown(context.Background())
		}

		if robot != nil {
			robot.Stop()
			device.Disconnect()
		}

		// Run Cleanup
		os.Exit(1)
	}()
}
