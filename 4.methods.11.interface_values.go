package main

import (
	"fmt"
	"math"
)

type Interface interface {
	M()
}

type TypeStruct struct {
	S string
}

func (t *TypeStruct) M() {
	fmt.Println(t.S)
}

type TypeFloat float64

func (f TypeFloat) M() {
	fmt.Println(f)
}

func describe(i Interface) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func main() {
	var i Interface

	i = &TypeStruct{"Hello"}
	describe(i)
	i.M()

	i = TypeFloat(math.Pi)
	describe(i)
	i.M()
}
