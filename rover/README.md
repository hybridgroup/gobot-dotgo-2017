# Rover

The Sphero SPRK+ and Sphero Ollie, and Sphero BB-8 all use the same API.

## Installation

```
go get -d -u gobot.io/x/gobot/...
```

## Running the code
When you run any of these examples, you will compile and execute the code on your computer. When you are running the program, you will be communicating with the robot using the Bluetooth Low Energy (LE) interface.

To compile/run the code, substitute the name of your SPRK+, Ollie or BB-8 as needed.

### OS X

To run any of the Gobot BLE code on OS X, you must use the `GODEBUG=cgocheck=0` flag.

For example:

```
$ GODEBUG=cgocheck=0 go run step01/main.go BB-128E
```

### Linux

On Linux the BLE code will need to run as a root user account. The easiest way to accomplish this is probably to use `go build` to build your program, and then to run the requesting executable using `sudo`.

For example:

```
$ go build -o step01 rover/step01/main.go
$ sudo ./step01 BB-123E
```

## Code

### step01/main.go

Change colors of the built-in LED.

### step02/main.go

Rolls around at random.

### step03/main.go

Gets collision notifications from robot.

### step04/main.go

Receive heartbeat data from base station.

### step05/main.go

Control robot using joystick.

### step06/main.go

Control robot using joystick to collect data and send to base station.

## License

Copyright (c) 2015-2017 The Hybrid Group. Licensed under the MIT license.
