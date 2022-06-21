package main

import (
	"github.com/go-vgo/robotgo"
	"image"
	"image/png"
	"os"
)

func main() {
	//capture, err := screenshot.Capture(0, 0, 1900, 800)
	//robotgo.
	//if err != nil {
	//	fmt.Println(err)
	//}
	//save(capture, "screen.png")

	bitmap := robotgo.CaptureScreen(0, 0, 1920, 800)

	robotgo.Save(robotgo.ToImage(bitmap), "test.png")
}

// save *image.RGBA to filePath with PNG format.
func save(img *image.RGBA, filePath string) {
	file, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	png.Encode(file, img)
}
