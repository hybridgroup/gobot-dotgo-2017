package main

import (
	"os"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/ble"
	// comment the following line, and uncomment subsequent for SPRK+
	"gobot.io/x/gobot/platforms/sphero/ollie"
	// "gobot.io/x/gobot/platforms/sphero/sprkplus"
)

func main() {
	bleAdaptor := ble.NewClientAdaptor(os.Args[1])
	// comment the following line, and uncomment subsequent for SPRK+
	rover := ollie.NewDriver(bleAdaptor)
	// rover := sprkplus.NewDriver(bleAdaptor)

	work := func() {
		gobot.Every(1*time.Second, func() {
			r := uint8(gobot.Rand(255))
			g := uint8(gobot.Rand(255))
			b := uint8(gobot.Rand(255))
			rover.SetRGB(r, g, b)
		})
	}

	robot := gobot.NewRobot("roverBot",
		[]gobot.Connection{bleAdaptor},
		[]gobot.Device{rover},
		work,
	)

	robot.Start()
}
