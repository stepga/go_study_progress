package main

import "fmt"

type I interface {
	M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func main() {
	var i I
	describe(i)
	//i.M() // BANG!
	// XXX: ^ a nil interface value: it holds no concrete value nor a type
}
