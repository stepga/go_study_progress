package main

import "fmt"

func main() {
	var i interface{}
	// XXX: ^  An empty interface may hold values of any type. (Every type implements at least zero methods.)

	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
