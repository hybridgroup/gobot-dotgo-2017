package main

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/ble"
	"gobot.io/x/gobot/platforms/joystick"
	"gobot.io/x/gobot/platforms/mqtt"
	"gobot.io/x/gobot/platforms/parrot/minidrone"
)

type pair struct {
	x float64
	y float64
}

var robot *gobot.Robot
var mqttAdaptor *mqtt.Adaptor

func ReportStatus(status string) {
	buf := new(bytes.Buffer)
	msg, _ := json.Marshal(status)
	binary.Write(buf, binary.LittleEndian, msg)
	mqttAdaptor.Publish("drones/"+robot.Name+"/status", buf.Bytes())
}

func main() {
	joystickConfig := os.Args[2]

	joystickAdaptor := joystick.NewAdaptor()
	joystick := joystick.NewDriver(joystickAdaptor,
		joystickConfig,
	)

	droneAdaptor := ble.NewClientAdaptor(os.Args[1])
	drone := minidrone.NewDriver(droneAdaptor)

	mqttAdaptor = mqtt.NewAdaptor(os.Args[3], "rover")
	mqttAdaptor.SetAutoReconnect(true)

	work := func() {
		offset := 32767.0
		rightStick := pair{x: 0, y: 0}
		leftStick := pair{x: 0, y: 0}

		joystick.On(joystick.Event("triangle_press"), func(data interface{}) {
			drone.TakeOff()
		})
		joystick.On(joystick.Event("square_press"), func(data interface{}) {
			drone.Stop()
		})
		joystick.On(joystick.Event("x_press"), func(data interface{}) {
			drone.Land()
		})
		joystick.On(joystick.Event("left_x"), func(data interface{}) {
			val := float64(data.(int16))
			if leftStick.x != val {
				leftStick.x = val
			}
		})
		joystick.On(joystick.Event("left_y"), func(data interface{}) {
			val := float64(data.(int16))
			if leftStick.y != val {
				leftStick.y = val
			}
		})
		joystick.On(joystick.Event("right_x"), func(data interface{}) {
			val := float64(data.(int16))
			if rightStick.x != val {
				rightStick.x = val
			}
		})
		joystick.On(joystick.Event("right_y"), func(data interface{}) {
			val := float64(data.(int16))
			if rightStick.y != val {
				rightStick.y = val
			}
		})

		gobot.Every(10*time.Millisecond, func() {
			pair := rightStick
			if pair.y < -10 {
				drone.Forward(minidrone.ValidatePitch(pair.y, offset))
			} else if pair.y > 10 {
				drone.Backward(minidrone.ValidatePitch(pair.y, offset))
			} else {
				drone.Forward(0)
			}

			if pair.x > 10 {
				drone.Right(minidrone.ValidatePitch(pair.x, offset))
			} else if pair.x < -10 {
				drone.Left(minidrone.ValidatePitch(pair.x, offset))
			} else {
				drone.Right(0)
			}
		})

		gobot.Every(10*time.Millisecond, func() {
			pair := leftStick
			if pair.y < -10 {
				drone.Up(minidrone.ValidatePitch(pair.y, offset))
			} else if pair.y > 10 {
				drone.Down(minidrone.ValidatePitch(pair.y, offset))
			} else {
				drone.Up(0)
			}

			if pair.x > 20 {
				drone.Clockwise(minidrone.ValidatePitch(pair.x, offset))
			} else if pair.x < -20 {
				drone.CounterClockwise(minidrone.ValidatePitch(pair.x, offset))
			} else {
				drone.Clockwise(0)
			}
		})

		drone.On(drone.Event("battery"), func(data interface{}) {
			fmt.Printf("battery: %d\n", data)
		})

		drone.On(minidrone.Hovering, func(data interface{}) {
			fmt.Println("hovering!")
			ReportStatus("hovering")
		})

		drone.On(minidrone.Landing, func(data interface{}) {
			fmt.Println("landing!")
			ReportStatus("landing")
		})

		drone.On(minidrone.Landed, func(data interface{}) {
			fmt.Println("landed.")
			ReportStatus("landed")
		})
	}

	robot = gobot.NewRobot("minidrone",
		[]gobot.Connection{joystickAdaptor, droneAdaptor, mqttAdaptor},
		[]gobot.Device{joystick, drone},
		work,
	)

	robot.Start()
}
