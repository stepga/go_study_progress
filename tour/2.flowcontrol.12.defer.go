package main

import "fmt"

func helper() string {
	fmt.Println("helper")
	return "world"
}

func main() {
	defer fmt.Println(helper())
	// XXX a defer statemnt defers/delays the execution of a function until the
	//     surrounding function returns

	fmt.Println("hello")
}
