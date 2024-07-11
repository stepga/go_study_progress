package main

import (
	"fmt"
	"time"
)

func say(s string) {
	// XXX: execution of `say` happens in new goroutine
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s, i)
	}
}

func main() {
	go say("world")
	// XXX: evaluation of `say` and `"world"` happens in goroutine `main`
	say("hello")
	// XXX: ATTENTION: goroutines run in the same address space
	//      -> synchronize shared memory (see `sync` package)
}
