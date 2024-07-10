package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	// XXX:         ^  `error` built-in interface:
	// type error interface {
	//     Error() string
	// }
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run(should_fail bool) error {
	if !should_fail {
		return nil
	}
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	if err := run(false); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("success: no error")
	}

	if err := run(true); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("success: no error")
	}
}
