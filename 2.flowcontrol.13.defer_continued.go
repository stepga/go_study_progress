package main

import "fmt"

func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Printf("defer put on stack: %d\n", i)
		// XXX: a deferred function’s arguments are evaluated when the defer statement is evaluated.
		// XXX: deferred function calls are executed in Last In First Out order after the surrounding function returns.
		// XXX: deferred functions may read and assign to the returning function’s named return values.
		// XXX: see also: `panic` and `recover`
	}

	fmt.Println("done")
}
