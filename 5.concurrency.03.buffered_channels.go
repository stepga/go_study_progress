package main

import "fmt"

func fill(ch chan int, value int) {
	fmt.Printf("about to put %d into channel\n", value)
	ch <- value
	fmt.Printf("done: %d\n", value)
}

func main() {
	ch := make(chan int, 2)
	// XXX:              ^ channel's buffer length
	ch <- 1
	// XXX ^: sends to a buffered channel blocks only if the buffer is full
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	//fmt.Println(<-ch) // XXX: fatal error: all goroutines are asleep - deadlock!

	fill(ch, 1)
	fill(ch, 2)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
