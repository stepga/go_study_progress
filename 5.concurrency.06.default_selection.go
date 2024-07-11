package main

import (
	"fmt"
	"time"
)

func main() {
	var tick <-chan time.Time = time.Tick(100 * time.Millisecond)
	var boom <-chan time.Time = time.After(500 * time.Millisecond)

	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			// XXX: use a `default` case within a `select` to prevent blocking
			time.Sleep(50 * time.Millisecond)
			fmt.Println("    + 50ms")
		}
	}
}
