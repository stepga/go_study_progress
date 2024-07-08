package main

import "fmt"

func add(x int, y int) int {
	// XXX:^      ^        add() takes 2 params of type int
	//                   ^ and returns int
	return x + y
}

func main() {
	// XXX: ^  Go's main() takes no arguments
	//         (c.f. `int main(int argc, char *argv[])`)
	fmt.Println(add(42, 13))

	// XXX: distinction between type and expression syntax makes it easy to
	//      invoke closures
	sum := func(a, b int) int { return a + b }(3, 4)
	fmt.Printf("Sum from closure: %d\n", sum)
}

// XXX: TYPE syntax arrays/slices: brackets left of type
// var a []int
// EXPRESSION syntax arrays/slices: brackets right of type
// x = a[1]

// XXX: TYPE syntax pointers: brackets left of type
// var p *int
// EXPRESSION syntax arrays/slices: brackets right of type
// x = *p
