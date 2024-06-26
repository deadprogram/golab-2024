/*
How to run

	go run ./demo/videodrone
*/

package main

import (
	"fmt"
	"io"

	"gocv.io/x/gocv"
)

func main() {
	net, outputNames, err := prepareDNNModel()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer net.Close()

	go startJoystick()
	go startDrone()
	go startVideo()

	// now handle video frames from ffmpeg stream in main thread, to be macOS/Windows friendly
	for {
		if _, err := io.ReadFull(ffmpegOut, buf); err != nil {
			fmt.Println(err)
			continue
		}
		img, _ := gocv.NewMatFromBytes(frameY, frameX, gocv.MatTypeCV8UC3, buf)
		if img.Empty() {
			continue
		}

		detect(net, &img, outputNames)

		window.IMShow(img)
		if window.WaitKey(1) >= 0 {
			go robot.Stop()

			ffmpeg.Process.Kill()
			window.Close()

			break
		}
	}
}
