package main

import (
	"fmt"

	"github.com/google/go-cmp/cmp"
	"github.com/stepga/go_study_progress/how_to_write_go_code/hello/morestrings"
)

func main() {
	fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}
