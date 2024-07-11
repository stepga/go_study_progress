package main

// XXX: go run 3.moretypes.18.excercise_slices.go | sed 's/IMAGE://g' | base64 -d  > foo; eog foo && rm foo

import "golang.org/x/tour/pic"

// XXX: go get golang.org/x/tour/pic

func Pic(dx, dy int) [][]uint8 {
	// DONE: Implement Pic. It should return a slice of length dy, each element
	// of which is a slice of dx 8-bit unsigned integers. When you run the
	// program, it will display your picture, interpreting the integers as
	// grayscale (well, bluescale) values.
	// The choice of image is up to you. Interesting functions include (x+y)/2,
	// x*y, and x^y.
	// (You need to use a loop to allocate each []uint8 inside the [][]uint8.)
	// (Use uint8(intValue) to convert between types.)
	//ret := make([][]uint8, dy)
	//return ret
	var counter uint8 = 0
	ret := make([][]uint8, dy)
	for i, _ := range ret {
		ret[i] = make([]uint8, dx)
		for j, _ := range ret[i] {
			ret[i][j] = counter
			counter++
		}
	}
	return ret
}

func main() {
	pic.Show(Pic)
}
