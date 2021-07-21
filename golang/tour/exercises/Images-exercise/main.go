package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
	W int
	H int
}

/*
type Image interface {
	// ColorModel returns the Image's color model.
	ColorModel() color.Model
	// Bounds returns the domain for which At can return non-zero color.
	// The bounds do not necessarily contain the point (0, 0).
	Bounds() Rectangle
	// At returns the color of the pixel at (x, y).
	// At(Bounds().Min.X, Bounds().Min.Y) returns the upper-left pixel of the grid.
	// At(Bounds().Max.X-1, Bounds().Max.Y-1) returns the lower-right one.
	At(x, y int) color.Color
}
*/

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.W, i.H)
}

func (i Image) At(x, y int) color.Color {
	var v uint8
	tmp := x + y
	switch {
	case tmp%12 == 0:
		v = uint8((x + y) / 2)
	case tmp%5 == 0:
		v = uint8(x * y)
	case tmp%3 == 0:
		v = uint8(x ^ y)
	}
	return color.RGBA{v, v, 255, 255}
}

func main() {
	m := Image{64, 64}
	pic.ShowImage(m)
}
