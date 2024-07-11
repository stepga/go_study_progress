package main

import (
	"fmt"
	"github.com/stepga/go_study_progress/how_to_write_go_code/hello/morestrings"
)

func main() {
	fmt.Println(morestrings.ReverseRunes("hello"))
	fmt.Println(morestrings.ReverseRunes("olleh"))
}
