package display

import (
	"bytes"
	"github.com/disintegration/imaging"
	"github.com/go-vgo/robotgo"
	"image"
)

type Img struct {
	source image.Image
}

// newImg new Img
func newImg(img image.Image) *Img {
	return &Img{source: img}
}

// SaveToPNG save to png
func (i Img) SaveToPNG(path string) error {
	return robotgo.SavePng(i.source, path)
}

// Encode  Img to bytes with png format
func (i Img) Encode() (*bytes.Buffer, error) {
	buffer := bytes.NewBuffer([]byte{})
	err := imaging.Encode(buffer, i.source, imaging.PNG)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}
