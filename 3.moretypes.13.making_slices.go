package main

import "fmt"

func main() {
	a := make([]int, 5) // len(a)=5
	// XXX: make allocates a zeroed array and returns a slice to that array
	printSlice("a", a)

	b := make([]int, 0, 5) // len(b)=0, cap(b)=5
	// XXX:             ^ capacity passed to make
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
