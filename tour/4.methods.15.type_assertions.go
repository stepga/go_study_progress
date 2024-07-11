package main

import "fmt"

func main() {
	var i interface{} = "hello"

	s := i.(string)
	// XXX: This statement asserts that the interface value `i` holds the
	//      concrete type `string` and assigns the underlying `string` value to
	//      the variable `s`.
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	// XXX: if the type assertion fails `f` will be the zero value of type `float64`
	fmt.Println(f, ok)

	//f = i.(float64) // panic!
	// XXX: ^ If `i` does not hold a `float64`, the statement will trigger a panic.
	//fmt.Println(f)
}
