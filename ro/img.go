package ro

import (
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

func (m Img) SaveToPNG(path string) error {
	return robotgo.SavePng(m.source, path)
}
