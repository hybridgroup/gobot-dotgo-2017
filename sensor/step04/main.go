package main

import (
	"fmt"
	"os"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/firmata"
	"gobot.io/x/gobot/platforms/mqtt"
)

var button *gpio.GroveButtonDriver
var blue *gpio.GroveLedDriver
var green *gpio.GroveLedDriver

func TurnOff() {
	blue.Off()
	green.Off()
}

func Reset() {
	TurnOff()
	fmt.Println("Sensor ready.")
	green.On()
}

func main() {
	board := firmata.NewAdaptor(os.Args[1])

	button = gpio.NewGroveButtonDriver(board, "2")
	blue = gpio.NewGroveLedDriver(board, "3")
	green = gpio.NewGroveLedDriver(board, "4")

	mqttAdaptor := mqtt.NewAdaptor(os.Args[2], "sensor")
	mqttAdaptor.SetAutoReconnect(true)

	heartbeat := mqtt.NewDriver(mqttAdaptor, "basestation/heartbeat")

	work := func() {
		Reset()

		button.On(gpio.ButtonPush, func(data interface{}) {
			TurnOff()
			fmt.Println("On!")
		})

		button.On(gpio.ButtonRelease, func(data interface{}) {
			Reset()
		})

		heartbeat.On(mqtt.Data, func(data interface{}) {
			fmt.Println("heartbeat")
			blue.Toggle()
		})
	}

	robot := gobot.NewRobot("sensorStation",
		[]gobot.Connection{board, mqttAdaptor},
		[]gobot.Device{button, blue, green, heartbeat},
		work,
	)

	robot.Start()
}
