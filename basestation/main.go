// TO RUN:
//  go run ./basestation/main.go <SERVER>
//
// EXAMPLE:
//	go run ./basestation/main.go ssl://iot.eclipse.org:8883
//
package main

import (
	"fmt"
	"os"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/mqtt"
)

func main() {
	mqttAdaptor := mqtt.NewAdaptor(os.Args[1], "basestation")
	mqttAdaptor.SetAutoReconnect(true)

	heartbeat := mqtt.NewDriver(mqttAdaptor, "basestation/heartbeat")

	work := func() {
		data := []byte("hi")
		gobot.Every(1*time.Second, func() {
			fmt.Println("Heartbeat")
			heartbeat.Publish(data)
		})
	}

	robot := gobot.NewRobot("basestation",
		[]gobot.Connection{mqttAdaptor},
		[]gobot.Device{heartbeat},
		work,
	)

	robot.Start()
}
