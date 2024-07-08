package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
	// XXX: ^  naked return: a return statement without arguments returns
	//         the named return values
	// XXX: use only in short function (or never ;-) ...)
}

func main() {
	a, b := split(17)
	fmt.Printf("Split 17 into: %d, %d\n", a, b)
}
