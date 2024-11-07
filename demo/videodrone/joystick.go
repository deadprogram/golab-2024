package main

import (
	"fmt"
	"sync/atomic"
	"time"

	"gobot.io/x/gobot/v2"
	"gobot.io/x/gobot/v2/platforms/dji/tello"
	"gobot.io/x/gobot/v2/platforms/joystick"
)

var (
	joyAdaptor                   = joystick.NewAdaptor("0")
	stick                        = joystick.NewDriver(joyAdaptor, "dualshock4")
	leftX, leftY, rightX, rightY atomic.Value

	offset = 32767.0
)

type pair struct {
	x float64
	y float64
}

func init() {
	leftX.Store(float64(0.0))
	leftY.Store(float64(0.0))
	rightX.Store(float64(0.0))
	rightY.Store(float64(0.0))
}

func startJoystick() {
	stick.On(joystick.CirclePress, func(data interface{}) {
		drone.Forward(0)
		drone.Up(0)
		drone.Clockwise(0)
	})
	stick.On(joystick.SquarePress, func(data interface{}) {
		fmt.Println("battery:", flightData.BatteryPercentage)
	})
	stick.On(joystick.TrianglePress, func(data interface{}) {
		drone.ThrowTakeOff()
		println("Takeoff")
	})
	stick.On(joystick.XPress, func(data interface{}) {
		drone.PalmLand()
		println("Land")
	})
	stick.On(joystick.CirclePress, func(data interface{}) {
		drone.PalmLand()
		println("Land")
	})
	stick.On(joystick.LeftX, func(data interface{}) {
		val := float64(data.(int))
		leftX.Store(val)
	})

	stick.On(joystick.LeftY, func(data interface{}) {
		val := float64(data.(int))
		leftY.Store(val)
	})

	stick.On(joystick.RightX, func(data interface{}) {
		val := float64(data.(int))
		rightX.Store(val)
	})

	stick.On(joystick.RightY, func(data interface{}) {
		val := float64(data.(int))
		rightY.Store(val)
	})
	gobot.Every(50*time.Millisecond, func() {
		rightStick := getRightStick()

		switch {
		case rightStick.y < -10:
			drone.Forward(tello.ValidatePitch(rightStick.y, offset))
		case rightStick.y > 10:
			drone.Backward(tello.ValidatePitch(rightStick.y, offset))
		default:
			drone.Forward(0)
		}

		switch {
		case rightStick.x > 10:
			drone.Right(tello.ValidatePitch(rightStick.x, offset))
		case rightStick.x < -10:
			drone.Left(tello.ValidatePitch(rightStick.x, offset))
		default:
			drone.Right(0)
		}
	})

	gobot.Every(50*time.Millisecond, func() {
		leftStick := getLeftStick()
		switch {
		case leftStick.y < -10:
			drone.Up(tello.ValidatePitch(leftStick.y, offset))
		case leftStick.y > 10:
			drone.Down(tello.ValidatePitch(leftStick.y, offset))
		default:
			drone.Up(0)
		}

		switch {
		case leftStick.x > 20:
			drone.Clockwise(tello.ValidatePitch(leftStick.x, offset))
		case leftStick.x < -20:
			drone.CounterClockwise(tello.ValidatePitch(leftStick.x, offset))
		default:
			drone.Clockwise(0)
		}
	})
}

func getLeftStick() pair {
	s := pair{x: 0, y: 0}
	s.x = leftX.Load().(float64)
	s.y = leftY.Load().(float64)
	return s
}

func getRightStick() pair {
	s := pair{x: 0, y: 0}
	s.x = rightX.Load().(float64)
	s.y = rightY.Load().(float64)
	return s
}
