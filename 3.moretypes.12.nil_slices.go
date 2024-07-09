package main

import "fmt"

func main() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		// XXX: nil slice: zero value == nil, BUT: length & capacity of 0
		fmt.Println("nil!")
	}
}
