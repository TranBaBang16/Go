package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
	w int
	h int
}

func (im Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, im.w, im.h)
}

func (im Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (im Image) Pic() uint8 {
	return uint8((im.w + im.h) / 2)
}

func (im Image) At(x, y int) color.Color {
	x = im.w
	y = im.h
	v := im.Pic()
	return color.RGBA{v, v, 255, 255}
}

func main() {
	m := Image{24, 113}
	pic.ShowImage(m)
}
