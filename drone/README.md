# Drone

The various Parrot Mindrones such as the Rolling Spider all use the same API.

## What you need

    - Parrot Minidrone
    - Dualshock 3 gamepad, or compatible
    - Personal computer with Go installed, and a Bluetooth 4.0 radio.
    - Linux or OS X

## Installation

```
go get -d -u gobot.io/x/gobot/...
```

## Running the code
When you run any of these examples, you will compile and execute the code on your computer. When you are running the program, you will be communicating with the robot using the Bluetooth Low Energy (LE) interface.

To compile/run the code, substitute the name of your Minidrone as needed.

### OS X

To run any of the Gobot BLE code on OS X, you must use the `GODEBUG=cgocheck=0` flag.

For example:

```
$ GODEBUG=cgocheck=0 go run drone/step01/main.go Mars_1234
```

### Linux

On Linux the BLE code will need to run as a root user account. The easiest way to accomplish this is probably to use `go build` to build your program, and then to run the requesting executable using `sudo`.

For example:

```
$ go build -o step01 drone/step01/main.go
$ sudo ./step01 Mars_1234
```

## Code

### step01/main.go

Takeoff and land.

### step02/main.go

Data.

### step03/main.go

Move.

### step04/main.go

Flips.

### step05/main.go

Freeflight.

### step06/main.go

Freeflight with data collection.

## License

Copyright (c) 2015-2017 The Hybrid Group. Licensed under the MIT license.
