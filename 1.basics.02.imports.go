package main

// XXX: good style: parenthesized import
import (
	"fmt"
	"math"
)

// XXX: ... or multiple import statements:
// import "fmt"
// import "math"

func main() {
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
}
