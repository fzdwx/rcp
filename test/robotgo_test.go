package test

import (
	"fmt"
	"github.com/axgle/mahonia"
	"github.com/go-vgo/robotgo"
	"github.com/vcaesar/imgo"
	"testing"
)

func Test_screen(t *testing.T) {
	x, y := robotgo.GetMousePos()
	fmt.Println("pos: ", x, y)

	color := robotgo.GetPixelColor(100, 200)
	fmt.Println("color----", color)

	sx, sy := robotgo.GetScreenSize()
	fmt.Println("get screen size: ", sx, sy)

	bit := robotgo.CaptureScreen(10, 10, 30, 30)
	defer robotgo.FreeBitmap(bit)

	img := robotgo.ToImage(bit)
	imgo.Save("test.png", img)
}

func Test_get_title(t *testing.T) {
	var enc = mahonia.NewDecoder("gbk")
	for {
		fmt.Println(enc.ConvertString(robotgo.GetTitle()))
	}
}
