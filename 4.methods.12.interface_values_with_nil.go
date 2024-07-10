package main

import "fmt"

type Interface interface {
	M()
}

type Type struct {
	S string
}

func (t *Type) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func describe(i Interface) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func main() {
	var i Interface

	var t *Type
	i = t
	describe(i)
	i.M()

	i = &Type{"hello"}
	describe(i)
	i.M()
}
