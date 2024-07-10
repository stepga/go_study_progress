package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %g", float64(e))
}

func sqrt_base(x float64) float64 {
	var z, last_z float64 = 1, 0

	if x == 0 {
		return 0
	}

	for round := 0; ; round++ {
		z -= (z*z - x) / (2 * z)
		if last_z != 0.0 && last_z-z < 0.1 {
			break
		}
		last_z = z
	}
	return z
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	return sqrt_base(x), nil
}

func main() {
	if sqrt, err := Sqrt(2); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Sqrt(2):", sqrt)
	}

	if sqrt, err := Sqrt(-2.1); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Sqrt(2):", sqrt)
	}

	fmt.Println(Sqrt(-2.2))
}
