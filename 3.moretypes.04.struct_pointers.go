package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9 // is equal to: (*p).X = 1e9
	// XXX: explicit dereferencing of pointer to structs is not necessary
	fmt.Println(v)
}
