package main

import (
	"fmt"
	"os"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/firmata"
)

func main() {
	board := firmata.NewAdaptor(os.Args[1])

	button := gpio.NewGroveButtonDriver(board, "2")
	blue := gpio.NewGroveLedDriver(board, "3")

	work := func() {
		button.On(gpio.ButtonPush, func(data interface{}) {
			fmt.Println("On!")
			blue.On()
		})

		button.On(gpio.ButtonRelease, func(data interface{}) {
			fmt.Println("Off!")
			blue.Off()
		})
	}

	robot := gobot.NewRobot("sensorStation",
		[]gobot.Connection{board},
		[]gobot.Device{button, blue},
		work,
	)

	robot.Start()
}
