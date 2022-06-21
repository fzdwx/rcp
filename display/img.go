package display

import (
	"bytes"
	"github.com/disintegration/imaging"
	"image"
)

type Img struct {
	source image.Image
}

// newImg new Img
func newImg(img image.Image) *Img {
	return &Img{source: img}
}

// Save saves the image to file with the specified filename.
// The format is determined from the filename extension:
// "jpg" (or "jpeg"), "png", "gif", "tif" (or "tiff") and "bmp" are supported.
//
// Examples:
//
//	// Save the image as PNG.
//	err := imaging.Save(img, "out.png")
//
//	// Save the image as JPEG.
//	err := imaging.Save(img, "out.jpg")
//
func (i Img) Save(path string) error {
	return imaging.Save(i.source, path)
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
