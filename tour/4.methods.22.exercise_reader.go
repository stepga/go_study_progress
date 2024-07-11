package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// DONE: Add a Read([]byte) (int, error) method to MyReader.

func (m MyReader) Read(b []byte) (int, error) {
	b[0] = byte('A')
	return 1, nil
	// XXX or:
	//for x := range b {
	//	b[x] = 'A'
	//}
	//return len(b), nil
}

func main() {
	reader.Validate(MyReader{})
}
