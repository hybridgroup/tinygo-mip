# Tutorial

Este tutorial contiene una serie de pequeñas actividades y ejercicios que te ayudarán a aprender cómo controlar el robot MiP usando Bluetooth.

## Encontrar la dirección MAC o el Bluetooth ID para el robot MiP

Necesitas la dirección MAC o el Bluetooth ID del robot MiP para conectarte a él.

En Linux y Windows utilizarás la dirección MAC del dispositivo para conectarte.

En macOS deberá utilizar el ID Bluetooth del dispositivo para conectarse.

Por lo tanto, debes conocer el nombre correcto y luego la dirección MAC o ID de ese dispositivo para poder conectarte a él.

Para averiguar la dirección MAC única o el ID Bluetooth de un dispositivo, puedes utilizar el escáner Bluetooth que se encuentra en el directorio de herramientas de esta repo.

Primero, cambia el directorio actual al directorio `tools`:

```shell
cd tools
```

A continuación, ejecuta el comando del escáner de Bluetooth:

```shell
go run ./blescanner
```

## El tutorial

Los pasos del tutorial se pueden ejecutar tanto en tu ordenador como en cualquier microcontrolador con capacidades Bluetooth, como por ejemplo el Pimoroni Badger2040-W.

### step1

El primer paso del tutorial comprueba que el robot MiP se conecta correctamente con tu ordenador encendiendo el LED que tiene en el pecho.

#### Ejecutando el código en tu ordenador

```shell
go run ./tutorial/step1/ [dirección MAC o Bluetooth ID]
```

Presiona "Control+C" para terminar el programa.

#### Ejecutando el código en tu microcontrolador

```shell
tinygo flash -target badger2040-w -ldflags="-X main.DeviceAddress=[MAC address]" ./tutorial/step1/
```

### step2

Rueda hacia adelante y hacia atrás.

#### Ejecutando el código en tu ordenador

```shell
go run ./step2/ [dirección MAC o Bluetooth ID]
```

Presiona "Control+C" para terminar el programa.

#### Ejecutando el código en tu microcontrolador

```shell
tinygo flash -target badger2040-w -ldflags="-X main.DeviceAddress=[MAC address]" ./tutorial/step2/
```

### step3

Rueda formando un padrón en cuadrado.

#### Ejecutando el código en tu ordenador

```shell
go run ./step3/ [dirección MAC o Bluetooth ID]
```

Presiona "Control+C" para terminar el programa.

#### Ejecutando el código en tu microcontrolador

```shell
tinygo flash -target badger2040-w -ldflags="-X main.DeviceAddress=[MAC address]" ./tutorial/step3/
```

### step4

Rueda hacia adelante, hacia atrás y entonces hace un sonido.

#### Ejecutando el código en tu ordenador

```shell
go run ./step4/ [dirección MAC o Bluetooth ID]
```

#### Ejecutando el código en tu microcontrolador

```shell
tinygo flash -target badger2040-w -ldflags="-X main.DeviceAddress=[MAC address]" ./tutorial/step4/
```

### step5

Se prepara y realiza un pequeño baile.

#### Ejecutando el código en tu ordenador

```shell
go run ./step5/ [dirección MAC o Bluetooth ID]
```

#### Ejecutando el código en tu microcontrolador

```shell
tinygo flash -target badger2040-w -ldflags="-X main.DeviceAddress=[MAC address]" ./tutorial/step5/
```


### step6

Ahora toma el control del robot MiP y manéjalo usando un controlador DS3 desde tu ordenador.

Conecta el controlador DS3 en tu ordenador, los controles son:

* Joystick derecho - dirección

NOTA IMPORTANTE: debes pulsar el botón «P3» cuando tu programa se ejecute por primera vez para que los joysticks DS3 «clónicos» que estamos utilizando se enciendan completamente.


#### Ejecutando el código en tu ordenador

```shell
go run ./step6/ [dirección MAC o Bluetooth ID]
```

