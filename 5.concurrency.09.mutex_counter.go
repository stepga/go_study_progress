package main

import (
	"fmt"
	"sync"
	"time"
)

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key]++
	c.mu.Unlock()
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer c.mu.Unlock()
	return c.v[key]
}

func main() {
	count_goal := 90000000
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < count_goal; i++ {
		go c.Inc("somekey")
	}

	for {
		if c.Value("somekey") == count_goal {
			break
		}
		time.Sleep(time.Microsecond * 50)
	}
	fmt.Println("Done: ", c.Value("somekey"))
}
