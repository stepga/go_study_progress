package main

import "fmt"

type Vertex struct {
	X, Y float64
}

// XXX: functions with pointer argument _must_ take a pointer:
// var v Vertex
// ScaleFunc(v, 5)  // Compile error!
// ScaleFunc(&v, 5) // OK

// XXX: methods with pointer receivers can take both:
// var v Vertex
// v.Scale(5)  // OK
// p := &v
// p.Scale(10) // OK

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(2)
	// XXX: ^ as a convenience, Go interprets the statement v.Scale(5) as
	//     (&v).Scale(5) since the Scale method has a pointer receiver.
	ScaleFunc(&v, 10)

	p := &Vertex{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)

	fmt.Println(v, p)
}
