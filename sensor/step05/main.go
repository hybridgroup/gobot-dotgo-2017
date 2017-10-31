package main

import (
	"fmt"
	"os"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/firmata"
	"gobot.io/x/gobot/platforms/mqtt"
)

var button *gpio.GroveButtonDriver
var blue *gpio.GroveLedDriver
var green *gpio.GroveLedDriver
var buzzer *gpio.GroveBuzzerDriver
var touch *gpio.GroveTouchDriver

func Alert() {
	TurnOff()
	buzzer.Tone(gpio.C4, gpio.Half)
	time.Sleep(1 * time.Second)
	Reset()
}

func TurnOff() {
	green.Off()
	blue.Off()
}

func Reset() {
	TurnOff()
	fmt.Println("Sensors ready.")
	green.On()
}

func main() {
	board := firmata.NewAdaptor(os.Args[1])

	button = gpio.NewGroveButtonDriver(board, "2")
	blue = gpio.NewGroveLedDriver(board, "3")
	green = gpio.NewGroveLedDriver(board, "4")
	buzzer = gpio.NewGroveBuzzerDriver(board, "6")
	touch = gpio.NewGroveTouchDriver(board, "8")

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

		touch.On(gpio.ButtonPush, func(data interface{}) {
			Alert()
		})
	}

	robot := gobot.NewRobot("sensorStation",
		[]gobot.Connection{board, mqttAdaptor},
		[]gobot.Device{button, blue, green, heartbeat, buzzer, touch},
		work,
	)

	robot.Start()
}
