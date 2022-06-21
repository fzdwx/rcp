package ro

import (
	"github.com/kbinani/screenshot"
	"image"
)

// All union all display
func All() (*Img, error) {
	if ActiveDisplayNums <= 0 {
		return nil, DisplayNotFound
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
