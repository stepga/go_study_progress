package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	var z, last_z float64 = 1, 0

	for round := 0; ; round++ {
		z -= (z*z - x) / (2 * z)
		fmt.Printf("sqrt(%.2f) %d: %.2f\n", x, round, z)
		if last_z != 0.0 && last_z-z < 0.1 {
			break
		}
		last_z = z
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(23213212131))
	fmt.Println(Sqrt(232132121312321313131))
}
