package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	var float_var_explicit float64 = math.Sqrt(float64(x*x + y*y))
	float_var_implicit := math.Sqrt(21)
	var z uint = uint(float_var_explicit)
	fmt.Println(x, y, z, float_var_explicit, float_var_implicit)
}
