package main

import "fmt"

// DONE: fibonacci is a function that returns a function that returns an int.
func fibonacci() func() int {
	var previous, current int = 0, 0

	return func() int {
		if previous == 0 {
			previous = 1
		} else {
			current_old := current
			current = current + previous
			previous = current_old
		}
		return current
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
