package main

import (
	"fmt"
	"time"

	"gobot.io/x/gobot/v2"
	"gobot.io/x/gobot/v2/platforms/dji/tello"
)

var (
	robot *gobot.Robot

	// drone
	drone      = tello.NewDriver("8890")
	flightData *tello.FlightData
)

func startDrone() {
	drone.On(tello.FlightDataEvent, func(data interface{}) {
		// TODO: protect flight data from race condition
		flightData = data.(*tello.FlightData)
	})

	drone.On(tello.ConnectedEvent, func(data interface{}) {
		fmt.Println("Connected")
		drone.StartVideo()
		drone.SetVideoEncoderRate(tello.VideoBitRateAuto)
		drone.SetExposure(0)
		gobot.Every(100*time.Millisecond, func() {
			drone.StartVideo()
		})
	})

	drone.On(tello.VideoFrameEvent, func(data interface{}) {
		pkt := data.([]byte)
		if _, err := ffmpegIn.Write(pkt); err != nil {
			fmt.Println(err)
		}
	})

	robot = gobot.NewRobot("tello",
		[]gobot.Connection{joyAdaptor},
		[]gobot.Device{drone, stick},
	)

	robot.Start()
}
