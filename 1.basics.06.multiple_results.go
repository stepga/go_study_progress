package main

import "fmt"

func swap(x, y string) (string, string) {
	// XXX:               ^ the swap function returns two strings
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
