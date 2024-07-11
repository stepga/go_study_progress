package main

import (
	"fmt"
)

// XXX: test for a closed channel: `v, ok := <-ch`
// `ok` is false if there are no more values to receive and the channel is closed.

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
	// XXX: ONLY the sender should close the channel
	//      sending on a closed channel will lead to a panic
}

func main() {
	c := make(chan int, 10)
	channel_capacity := cap(c)
	go fibonacci(channel_capacity, c)
	for v := range c {
		// XXX: receive values repeatedly from channel until it's closed
		fmt.Println(v)
	}
}
