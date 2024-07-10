package main

// XXX: find implementations of `Read(b []byte)`:
//      https://cs.opensource.google/search?q=Read%5C(%5Cw%2B%5Cs%5C%5B%5C%5Dbyte%5C)&ss=go%2Fgo
//
// XXX: find implementations of `Write` in `net/http`:
//      go doc "net/http" Write
//      ...

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:%d] = %q\n", n, b[:n])
		if err == io.EOF {
			break
		}
	}
}
