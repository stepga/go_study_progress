package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// XXX: methods with value receivers take either a value or a pointer as the receiver when they are called
// var v Vertex
// fmt.Println(v.Abs()) // OK
// p := &v
// fmt.Println(p.Abs()) // OK

func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// XXX: Functions that take a value argument must take a value of that specific type
// var v Vertex
// fmt.Println(AbsFunc(v))  // OK
// fmt.Println(AbsFunc(&v)) // Compile error!

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(AbsFunc(v))

	p := &Vertex{4, 3}
	fmt.Println(p.Abs())
	fmt.Println(AbsFunc(*p))
}
