.ONESHELL:

build/showvideo:
	go build -o ./build/showvideo ./demo/showvideo/main.go

showvideo: builddir build/showvideo
	./build/showvideo 2

clean:
	rm -rf build

builddir:
	mkdir -p build

build/videodrone:
	go build -o ./build/videodrone ./demo/videodrone/

videodrone-cpu: builddir build/videodrone
	./build/videodrone ~/models/yolov8n.onnx

videodrone-gpu: builddir build/videodrone
	./build/videodrone ~/models/yolov8n.onnx cuda cuda
