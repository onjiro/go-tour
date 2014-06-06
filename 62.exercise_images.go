package main

import (
	"code.google.com/p/go-tour/pic"
	"image"
	"image/color"
)

type Image struct {
	Rectangle image.Rectangle
}

func (i *Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i *Image) Bounds() image.Rectangle {
	return i.Rectangle
}

func (i *Image) At(x, y int) color.Color {
	v := uint8(x * y)
	return &color.RGBA{v, v, 255, 255}
}

func main() {
	m := Image{image.Rect(0, 0, 100, 100)}
	pic.ShowImage(&m)
}
