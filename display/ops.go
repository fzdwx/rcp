package display

import (
	"github.com/go-vgo/robotgo"
	"github.com/kbinani/screenshot"
	"image"
)

// CapAll union all display
func CapAll() (*Img, error) {
	if ActiveDisplayNums <= 0 {
		return nil, NotFound
	}
	var all = image.Rect(0, 0, 0, 0)

	for i := 0; i < ActiveDisplayNums; i++ {
		bounds := screenshot.GetDisplayBounds(i)
		all = bounds.Union(all)
	}

	img, err := screenshot.Capture(all.Min.X, all.Min.Y, all.Dx(), all.Dy())
	if err != nil {
		return nil, err
	}
	return newImg(img), nil
}

// CapMain captures whole region of first display.
func CapMain() (*Img, error) {
	return CapByDisPlay(1)
}

// CapByDisPlay captures whole region of display Index'th display.
func CapByDisPlay(displayIndex int) (*Img, error) {
	img, err := screenshot.CaptureDisplay(displayIndex)
	if err != nil {
		return nil, err
	}
	return newImg(img), nil
}

// Capture custom crop
func Capture(x, y, width, height int) (*Img, error) {
	robotgo.CaptureGo()
	img, err := screenshot.Capture(x, y, width, height)
	if err != nil {
		return nil, err
	}
	return newImg(img), nil
}
