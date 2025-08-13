# Tutorial

This tutorial contains a series of small activities to help you learn how to control the MiP robot using Bluetooth.

## Finding the MAC address or Bluetooth ID for the MiP robot

You will need to determine what the MAC address or Bluetooth ID is for the MiP robot you want to connect to.

On Linux and Windows you will use the MAC address of the device to connect.

On macOS you must use the Bluetooth ID of the device to connect.

Therefore, you must know the correct name and then MAC address or ID for that device in order to connect to it.

First, install the Bluetooth scanner command:

```shell
go install ./cmd/blescanner
```

Then, run the command:

```shell
blescanner
```

It should show the names of the various Bluetooth devices around you, including the MiP you want to connect to.

## The tutorial

The tutorial steps can be run either on your computer, or on a Bluetooth enabled microcontroller such as the Pimoroni Badger2040-W.

### step1

This first step tests that the MiP is connected correctly to your computer by turning on the chest LED.

#### Running on your computer

```shell
go run ./tutorial/step1/ [MAC address or Bluetooth ID]
```

Press "Control-C" to stop the program. 

#### Running on your microcontroller

```shell
tinygo flash -target badger2040-w -ldflags="-X main.DeviceAddress=[MAC address]" ./tutorial/step1/
```

### step2

Rolls forwards and backwards.

#### Running on your computer

```shell
go run ./step2/ [MAC address or Bluetooth ID]
```

Press "Control-C" to stop the program. 

#### Running on your microcontroller

```shell
tinygo flash -target badger2040-w -ldflags="-X main.DeviceAddress=[MAC address]" ./tutorial/step2/
```

### step3

Rolls in a square pattern.

#### Running on your computer

```shell
go run ./step3/ [MAC address or Bluetooth ID]
```

Press "Control-C" to stop the program. 

#### Running on your microcontroller

```shell
tinygo flash -target badger2040-w -ldflags="-X main.DeviceAddress=[MAC address]" ./tutorial/step3/
```

### step4

Rolls forwards and back, and then makes a sound.

#### Running on your computer

```shell
go run ./step4/ [MAC address or Bluetooth ID]
```

#### Running on your microcontroller

```shell
tinygo flash -target badger2040-w -ldflags="-X main.DeviceAddress=[MAC address]" ./tutorial/step4/
```

### step5

Gets ready and then does a little dance.

#### Running on your computer

```shell
go run ./step5/ [MAC address or Bluetooth ID]
```

#### Running on your microcontroller

```shell
tinygo flash -target badger2040-w -ldflags="-X main.DeviceAddress=[MAC address]" ./tutorial/step5/
```


### step6

Now take control of MiP and drive it around using a DS3 controller from your computer.

Plug in the DS3 controller to your computer. The controls are as follows:

* Right stick - direction

IMPORTANT NOTE: you must press the "P3" button when your program first runs for the "clone" DS3 joysticks we are using to fully turn on.

#### Running on your computer


```shell
go run ./step6/ [MAC address or Bluetooth ID]
```
