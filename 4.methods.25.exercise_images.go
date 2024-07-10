package main

import "golang.org/x/tour/pic"
import "image"
import "image/color"

// XXX: go run 4.methods.25.exercise_images.go | sed 's/IMAGE://g' | base64 -d > foo && eog foo

type Image struct {
	min_x, min_y, max_x, max_y int
	color_r                    [][]uint8
	color_g                    [][]uint8
	color_b                    [][]uint8
	color_a                    [][]uint8
}

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

func (i Image) At(x int, y int) color.Color {
	return color.RGBA{i.color_r[x][y], i.color_g[x][y], i.color_b[x][y], i.color_a[x][y]}
}

func (i Image) Bounds() image.Rectangle {
	min := image.Point{i.min_x, i.min_y}
	max := image.Point{i.max_x, i.max_y}
	return image.Rectangle{min, max}
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func main() {
	a := make([][]uint8, 200)
	for i := range a {
		a[i] = make([]uint8, 200)
		var count uint8 = 0
		for j := range a[i] {
			a[i][j] = count
			count++
		}
	}

	m := Image{0, 0, 200, 200, a, a, a, a}
	pic.ShowImage(m)
}
