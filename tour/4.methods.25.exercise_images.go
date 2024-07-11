package main

import "golang.org/x/tour/pic"
import "image"
import "image/color"

// XXX: go run 4.methods.25.exercise_images.go | sed 's/IMAGE://g' | base64 -d > foo && eog foo

type Image struct {
	width, height int
	color         uint8
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.width, i.height)
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) At(x, y int) color.Color {
	return color.RGBA{i.color + uint8(x), i.color + uint8(y), 255, 255}
}

func main() {
	m := Image{100, 100, 100}
	pic.ShowImage(m)
}
