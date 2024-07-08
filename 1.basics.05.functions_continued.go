package main

import "fmt"

// XXX: omit the type for two or more consecutive named function params
func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
