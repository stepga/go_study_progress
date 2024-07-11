package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		// XXX: ^ `if` statement can start with a short statement like `for`
		return v
	}
	// XXX: Variables declared by the `if` statement are only in scope until the
	//      end of the if.
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
		pow(3, 3, 30),
	)
}
