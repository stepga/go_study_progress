package main

import "fmt"

const Pi = 3.14

// XXX:  ^ constants are assigned via `=` not via `:=`

func main() {
	const World = "世界"
	//fmt.Printf("Address of Pi: %p\n", &Pi)
	//^ XXX: this would not succedd: cannot take address of constant
	//       as constants are not addressable
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	Pi := "Birth"
	fmt.Printf("Address of Pi: %p\n", &Pi)
	// XXX: ^ variable overlays constant
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
