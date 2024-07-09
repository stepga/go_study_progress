package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		// XXX:  A closure is a function value that references variables from
		//       outside its body.
		sum += x
		// XXX: The function may access and assign to the referenced variables
		return sum
		// XXX:^ Each closure is "bound" to its own `sum` variable.
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
