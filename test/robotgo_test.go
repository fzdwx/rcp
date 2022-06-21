package test

import (
	"fmt"
	"github.com/axgle/mahonia"
	"github.com/go-vgo/robotgo"
	"github.com/kbinani/screenshot"
	hook "github.com/robotn/gohook"
	"github.com/vcaesar/imgo"
	"image"
	"image/png"
	"os"
	"testing"
)

// 测试
func Test_screen(t *testing.T) {
	x, y := robotgo.GetMousePos()
	fmt.Println("pos: ", x, y)

	color := robotgo.GetPixelColor(100, 200)
	fmt.Println("color----", color)

	sx, sy := robotgo.GetScreenSize()
	fmt.Println("get screen size: ", sx, sy)

	bit := robotgo.CaptureScreen(0, 0, sx, sy)
	defer robotgo.FreeBitmap(bit)
	//robotgo.SaveBitmap(bit, "cutBm.png")

	img := robotgo.ToImage(bit)
	imgo.Save("test.png", img)
}

// 截屏
func Test_SCREENSHOT(t *testing.T) {
	capture, err := screenshot.Capture(0, 0, 1900, 800)
	if err != nil {
		fmt.Println(err)
	}
	save(capture, "screen.png")
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

// 获取当前窗口的title
func Test_get_title(t *testing.T) {
	var enc = mahonia.NewDecoder("gbk")
	for {
		fmt.Println(enc.ConvertString(robotgo.GetTitle()))
	}
}

// 事件监听
func Test_event_listen(t *testing.T) {
	fmt.Println("--- Please press ctrl + shift + q ---")
	ok := hook.AddEvents("q", "ctrl", "shift")
	if ok {
		fmt.Println("add events...")
	}

	fmt.Println("--- Please press w---")
	ok = hook.AddEvents("w")
	if ok {
		fmt.Println("add events")
	}

	// start hook
	s := hook.Start()
	// end hook
	defer hook.End()

	for ev := range s {
		fmt.Println("hook: ", ev)
	}
}
