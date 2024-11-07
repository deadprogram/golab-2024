package main

import (
	"fmt"
	"os/exec"
	"strconv"
)

const (
	frameX    = 640
	frameY    = 480
	frameSize = frameX * frameY * 3
)

var (
	// ffmpeg command to decode video stream from drone
	ffmpeg = exec.Command("ffmpeg", "-hwaccel", "auto", "-hwaccel_device", "opencl", "-i", "pipe:0",
		"-nostats", "-flags", "low_delay", "-probesize", "32", "-fflags", "nobuffer+fastseek+flush_packets",
		"-analyzeduration", "0", "-af", "aresample=async=1:min_comp=0.1:first_pts=0",
		"-pix_fmt", "bgr24", "-s", strconv.Itoa(frameX)+"x"+strconv.Itoa(frameY), "-f", "rawvideo", "pipe:1")
	ffmpegIn, _  = ffmpeg.StdinPipe()
	ffmpegOut, _ = ffmpeg.StdoutPipe()

	buf = make([]byte, frameSize)
)

func startVideo() {
	if err := ffmpeg.Start(); err != nil {
		fmt.Println(err)
		return
	}
}
