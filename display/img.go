package display

import (
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
func (m Img) SaveToPNG(path string) error {
	return robotgo.SavePng(m.source, path)
}

func (m *Img) Resize() *Img {
	i := m.source.Bounds().Dx() / 4
	m.source = imaging.Resize(m.source, i, 0, imaging.Lanczos)
	return m
}
