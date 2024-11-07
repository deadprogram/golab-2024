package main

import (
	"image"
	"image/color"
	"os"

	"gocv.io/x/gocv"
)

func prepareModel() (*gocv.FaceDetectorYN, error) {
	model := os.Args[1]

	backend := gocv.NetBackendDefault
	if len(os.Args) > 2 {
		backend = gocv.ParseNetBackend(os.Args[2])
	}

	target := gocv.NetTargetCPU
	if len(os.Args) > 3 {
		target = gocv.ParseNetTarget(os.Args[3])
	}

	detector := gocv.NewFaceDetectorYNWithParams(model, "", image.Pt(200, 200), 0.9, 0.3, 5000, int(backend), int(target))

	return &detector, nil
}

var green = color.RGBA{0, 255, 0, 0}

func detectFaces(detector *gocv.FaceDetectorYN, img *gocv.Mat) {
	sz := img.Size()
	detector.SetInputSize(image.Pt(sz[1], sz[0]))

	faces := gocv.NewMat()
	defer faces.Close()

	detector.Detect(*img, &faces)
	if faces.Rows() < 1 {
		return
	}

	for r := 0; r < faces.Rows(); r++ {
		x0 := int(faces.GetFloatAt(r, 0))
		y0 := int(faces.GetFloatAt(r, 1))
		x1 := x0 + int(faces.GetFloatAt(r, 2))
		y1 := y0 + int(faces.GetFloatAt(r, 3))

		faceRect := image.Rect(x0, y0, x1, y1)

		gocv.Rectangle(img, faceRect, green, 1)
	}
}
