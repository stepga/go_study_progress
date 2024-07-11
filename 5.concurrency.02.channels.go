package main

import "fmt"

// XXX: - channels are typed
// XXX: - you can send and receive values with the channel operator `<-`
// XXX:   ch <- v    // Send v to channel ch.
// XXX:   v := <-ch  // Receive from ch, and
// XXX:              // assign value to v.
// XXX:   (The data flows in the direction of the arrow.)
// XXX:
// XXX: - Like maps and slices, channels must be created before use:
// XXX:   ch := make(chan int)
// XXX:
// XXX: - By default, sends and receives block until the other side is ready.
// XXX:   This allows goroutines to synchronize without explicit locks or condition variables.

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	fmt.Println("pre sending sum: ", sum)
	c <- sum // send sum to c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)

	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
}
