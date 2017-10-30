package main

import (
	"math"
	"os"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/ble"
	"gobot.io/x/gobot/platforms/joystick"
	"gobot.io/x/gobot/platforms/parrot/minidrone"
)

type pair struct {
	x float64
	y float64
}

func main() {
	pwd, _ := os.Getwd()
	joystickConfig := pwd + "/dualshock3.json"

	joystickAdaptor := joystick.NewAdaptor()
	joystick := joystick.NewDriver(joystickAdaptor,
		joystickConfig,
	)

	droneAdaptor := ble.NewClientAdaptor(os.Args[1])
	drone := minidrone.NewDriver(droneAdaptor)

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
				drone.Forward(validatePitch(pair.y, offset))
			} else if pair.y > 10 {
				drone.Backward(validatePitch(pair.y, offset))
			} else {
				drone.Forward(0)
			}

			if pair.x > 10 {
				drone.Right(validatePitch(pair.x, offset))
			} else if pair.x < -10 {
				drone.Left(validatePitch(pair.x, offset))
			} else {
				drone.Right(0)
			}
		})

		gobot.Every(10*time.Millisecond, func() {
			pair := leftStick
			if pair.y < -10 {
				drone.Up(validatePitch(pair.y, offset))
			} else if pair.y > 10 {
				drone.Down(validatePitch(pair.y, offset))
			} else {
				drone.Up(0)
			}

			if pair.x > 20 {
				drone.Clockwise(validatePitch(pair.x, offset))
			} else if pair.x < -20 {
				drone.CounterClockwise(validatePitch(pair.x, offset))
			} else {
				drone.Clockwise(0)
			}
		})
	}

	robot := gobot.NewRobot("minidrone",
		[]gobot.Connection{joystickAdaptor, droneAdaptor},
		[]gobot.Device{joystick, drone},
		work,
	)

	robot.Start()
}

func validatePitch(data float64, offset float64) int {
	value := math.Abs(data) / offset
	if value >= 0.1 {
		if value <= 1.0 {
			return int((float64(int(value*100)) / 100) * 100)
		}
		return 100
	}
	return 0
}
