package interfaces1

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

// Image struct to represent our custom image
type Image struct {
	width, height int
}

// Bounds method returns the domain for which At can return non-zero color.
func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.width, img.height)
}

// ColorModel returns the Image's color model.
func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

// At returns the color of the pixel at (x, y).
func (img Image) At(x, y int) color.Color {
	v := uint8((x + y) / 2)
	return color.RGBA{v, v, 255, 255}
}

func ImageInterface1() {
	m := Image{width: 256, height: 256}
	pic.ShowImage(m)
}
