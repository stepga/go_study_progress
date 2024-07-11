package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	// XXX:    ^ An array has a fixed size.

	var s []int = primes[1:4]
	// XXX: A slice is a dynamically-sized, flexible view into the elements of
	//      an array. In practice, slices are much more common than arrays.
	// XXX:                ^   array[low : high], excluding `high`

	fmt.Println(primes)
	fmt.Println("Slice [1:4]")
	fmt.Println(s)
}
