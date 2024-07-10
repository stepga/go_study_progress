package main

import (
	"fmt"
	"image"
)

// XXX: package image defines the Image interface
//
// type Image interface {
//     ColorModel() color.Model
//     Bounds() Rectangle
//     At(x, y int) color.Color
// }

func main() {
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	r, g, b, a := m.At(0, 0).RGBA()
	fmt.Printf("%v, %v, %v, %v\n", r, g, b, a)
}
