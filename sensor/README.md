# Sensor Station

## Installation

```
go get -d -u gobot.io/x/gobot/...
```

## Connecting the Arduino to your computer

Plug the Arduino into your computer using the USB cable provided in your starter kit. The Firmata firmware that we use for controlling the Arduino from your computer has already been flashed onto the Arduino board.

## Running the code
When you run any of these examples, you will compile and execute the code on your computer.

To compile/run the code:

```
$ go run step1.go /dev/ttyACM0
```

If using Mac OS X then the Arduino will probably use a device name like `/dev/tty.usbmodem1421`. Perform a directory listing of `/dev/`; the Arduino is likely a device named using the pattern `/dev/tty.usbmodem`.

Substitute the name of the program and the name of your serial port as needed.

The Gobot program will use the serial interface to communicate with the connected Arduino that is running the Firmata firmware. Your Arduino already has Firmata installed for you. If you need to reload Firmata on your Arduino you can use Gort (http://gort.io)

## Code

### step1.go - Blue LED

![Gobot](../../images/sensor/arduino/step1.jpg)

Connect the blue LED to D3.

Run the code.

You should see the blue LED blink.

### step2.go - Blue LED, Button

![Gobot](../../images/sensor/arduino/step2.jpg)

Connect the button to D2.

Run the code.

When you press the button, the blue LED should turn on.

### step3.go - Blue LED, Button, Green LED

![Gobot](../../images/sensor/arduino/step3.jpg)

Connect the Green LED to D4.

Run the code.

The green LED should light up. When you press the button, the blue LED should turn on, and the green LED should turn off.

### step4.go - Blue LED, Button, Green LED, Gobot API

This step has us playing with the Gobot API. No additional hardware needs to be connected.

Run the code.

You can now point your web browser to `http://<IP of your device>:3000` and try out the [Robeaux](https://github.com/hybridgroup/robeaux) web interface.

### step5.go - Blue LED, Button, Green LED, Gobot API, Buzzer, Touch

![Gobot](../../images/sensor/arduino/step5.jpg)

Connect the buzzer to D6, and connect the touch sensor to D8.

Run the code.

When your finger touches the capacitive touch sensor, the buzzer should sound.

### step6.go - Blue LED, Button, Green LED, Gobot API, Buzzer, Touch, Dial

![Gobot](../../images/sensor/arduino/step6.jpg)

Connect the rotary dial to A0.

Run the code.

Turning the dial will display the current analog reading on your console.

### step7.go - Blue LED, Button, Green LED, Gobot API, Buzzer, Touch, Dial, Temperature, Red LED

![Gobot](../../images/sensor/arduino/step7.jpg)

Connect the temperature sensor to A1, and the red LED to D5

Run the code.

By default, if the temperature exceeds 15 degrees, then the red LED will light up.
In case the room is warmer than 15 degrees, change the default temperature in step7.go.

In order to increase the temperature of the sensor, hold it between your fingers and wait for the LED to light up.
To turn the LED off, let go of the temperature sensor (note: the temperature will drop much slower than it increased).

### step8.go - Blue LED, Button, Green LED, Gobot API, Buzzer, Touch, Dial, Temperature, Red LED, Sound

![Gobot](../../images/sensor/arduino/step8.jpg)

Connect the sound sensor to A2.

Run the code.

When a sound is detected, the blue LED will light up, the sound sensor reading will be displayed on your computer's console.

### step9.go - Blue LED, Button, Green LED, Gobot API, Buzzer, Touch, Dial, Temperature, Red LED, Sound, Light

![Gobot](../../images/sensor/arduino/step9.jpg)

Connect the light sensor to A3.

Run the code.

When a light is detected, the blue LED will light up, and the light sensor reading will be displayed on your computer's console.

### step10.go - Blue LED, Button, Green LED, Gobot API, Buzzer, Touch, Dial, Temperature, Red LED, Sound, Light, LCD

![Gobot](../../images/sensor/arduino/step10.jpg)

Connect the LCD to any of the I2C ports on the Grove shield.

Run the code.

The LCD display will show informative messages, and also change the backlight color to match the color of whichever of the 3 LEDs is lit.

## License

Copyright (c) 2015-2017 The Hybrid Group. Licensed under the MIT license.
