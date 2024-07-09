package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for range_index, element_copy := range pow {
		fmt.Printf("2**%d = %d\n", range_index, element_copy)
	}
}
