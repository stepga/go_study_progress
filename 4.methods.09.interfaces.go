package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

// XXX: interface type: set of method signatures

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	var a Abser

	f := MyFloat(-math.Sqrt2)
	a = f // a MyFloat implements Abser
	fmt.Println(a.Abs())

	v := Vertex{3, 4}
	a = &v // a *Vertex implements Abser
	fmt.Println(a.Abs())

	// In the following line, v is a Vertex (not *Vertex)
	// and does NOT implement Abser.
	//a = v  // panic
}
