/*
How to run

	go run ./demo/videodrone [model] [backend] [target]
*/

package main

import (
	"fmt"
	"io"
	"log"

	"gocv.io/x/gocv"
)

func main() {
	net, err := prepareModel()
	if err != nil {
		log.Fatal(err)
	}

	go startJoystick()
	go startDrone()
	go startVideo()

	window := gocv.NewWindow("Tello")

	// now handle video frames from ffmpeg stream in main thread, to be macOS/Windows friendly
	for {
		if _, err := io.ReadFull(ffmpegOut, buf); err != nil {
			fmt.Println(err)
			continue
		}
		img, err := gocv.NewMatFromBytes(frameY, frameX, gocv.MatTypeCV8UC3, buf)
		switch {
		case err != nil:
			fmt.Println(err)
			continue
		case img.Empty():
			continue
		}

		detectFaces(net, &img)

		window.IMShow(img)

		// press any key to close and exit
		if window.WaitKey(1) >= 0 {
			go robot.Stop()

			ffmpeg.Process.Kill()
			window.Close()

			break
		}
	}
}
