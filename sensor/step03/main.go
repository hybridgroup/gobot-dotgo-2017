package main

import (
	"fmt"
	"os"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/firmata"
)

var button *gpio.GroveButtonDriver
var blue *gpio.GroveLedDriver
var green *gpio.GroveLedDriver

func TurnOff() {
	blue.Off()
}

func Reset() {
	TurnOff()
	fmt.Println("Sensors ready.")
}

func main() {
	board := firmata.NewAdaptor(os.Args[1])

	button = gpio.NewGroveButtonDriver(board, "2")
	blue = gpio.NewGroveLedDriver(board, "3")
	green = gpio.NewGroveLedDriver(board, "4")

	work := func() {
		Reset()

		button.On(gpio.ButtonPush, func(data interface{}) {
			TurnOff()
			fmt.Println("On!")
			blue.On()
		})

		button.On(gpio.ButtonRelease, func(data interface{}) {
			Reset()
		})

		gobot.Every(1*time.Second, func() {
			green.Toggle()
		})
	}

	robot := gobot.NewRobot("sensorStation",
		[]gobot.Connection{board},
		[]gobot.Device{button, blue, green},
		work,
	)

	robot.Start()
}
