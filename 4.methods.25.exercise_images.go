package main

import "golang.org/x/tour/pic"

type Image struct{}

// TODO: Define your own Image type, implement the necessary methods, and call pic.ShowImage.
//
// from package image:
//type Image interface {
//	// ColorModel returns the Image's color model.
//	ColorModel() color.Model
//	// Bounds returns the domain for which At can return non-zero color.
//	// The bounds do not necessarily contain the point (0, 0).
//	Bounds() Rectangle
//	// At returns the color of the pixel at (x, y).
//	// At(Bounds().Min.X, Bounds().Min.Y) returns the upper-left pixel of the grid.
//	// At(Bounds().Max.X-1, Bounds().Max.Y-1) returns the lower-right one.
//	At(x, y int) color.Color
//}

func main() {
	m := Image{}
	pic.ShowImage(m)
}
