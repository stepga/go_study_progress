package main

import "fmt"

var i, j int = 1, 2

// XXX:        ^ one initializer per variable

func main() {
	var c, python, java = true, false, "no!"
	// XXX:             ^ type can be omitted if initializers are present
	fmt.Println(i, j, c, python, java)
}
