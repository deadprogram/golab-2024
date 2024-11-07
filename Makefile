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
	./build/videodrone ~/models/face_detection_yunet_2023mar.onnx

videodrone-gpu: builddir build/videodrone
	./build/videodrone ~/models/face_detection_yunet_2023mar.onnx cuda cuda

wasmvision-blur:
	wasmvision run -p blur

wasmvision-asciify:
	wasmvision run -p asciify

wasmvision-ollama:
	wasmvision run -p ollama -p mosaic -logging=error

