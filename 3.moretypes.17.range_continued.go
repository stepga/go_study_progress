package main

import "fmt"

func main() {
	pow := make([]int, 10)
	fmt.Println("Filling pow slice:")
	for i, _ := range pow {
		// XXX: `for i, _ := range(...)` == `for i := range(...)`
		pow[i] = 1 << uint(i) // == 2**i
		fmt.Printf("i %d: %d\n", i, pow[i])
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}
