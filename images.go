package main

import (
	"image"
	"image/color"
)

// Image implements Image.image interface
type Image struct {
	width, height int
}

// Bounds return image.Rectangle
func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.width, i.height)
}

// ColorModel return image.color.Model
func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

// At return color of the pixcel at (x, y).
func (i Image) At(x, y int) color.Color {
	return color.RGBA{uint8(x * y), uint8(x * y), 255, 255}
}
