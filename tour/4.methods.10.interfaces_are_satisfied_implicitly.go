package main

import "fmt"

type InterfaceType interface {
	M()
}

type StructType struct {
	S string
}

// This method means type StructType implements the interface InterfaceType,
// but we don't need to explicitly declare that it does so.
func (t StructType) M() {
	fmt.Println(t.S)
}

func main() {
	var i InterfaceType = StructType{"hello"}
	i.M()
}
