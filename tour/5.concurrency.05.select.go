package main

import "fmt"

func fibonacci(channel_val, channel_quit chan int) {
	x, y := 0, 1
	for {
		select {
		case channel_val <- x:
			x, y = y, x+y
		case <-channel_quit:
			fmt.Println("channel_quit")
			return
		}
		// XXX: - `select` lets a goroutine wait on multiple communication operations
		// XXX: - `select` blocks until one of its cases can run, then it executes that case.
		//        It chooses one at random if multiple are ready
	}
}

func main() {
	channel_val := make(chan int)
	channel_quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-channel_val)
		}
		channel_quit <- 0
	}()
	fibonacci(channel_val, channel_quit)
}
