package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func rot13char(c byte) byte {
	if c >= 'a' && c <= 'm' || c >= 'A' && c <= 'M' {
		return c + 13
	} else if c >= 'n' && c <= 'z' || c >= 'N' && c <= 'Z' {
		return c - 13
	}
	return c
}

func (rotr rot13Reader) Read(b []byte) (int, error) {
	n, e := rotr.r.Read(b)
	for i := 0; i < n; i++ {
		b[i] = rot13char(b[i])
	}
	return n, e
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	fmt.Println("Reader before:", r.r)
	io.Copy(os.Stdout, &r)
	fmt.Printf("\nReader after: %v\n\n", r.r)

	buffer := make([]byte, 5)
	s = strings.NewReader("Lbh penpxrq gur pbqr!")
	r = rot13Reader{s}
	for n, e := r.Read(buffer); e == nil; n, e = r.Read(buffer) {
		fmt.Printf("%d read bytes: '%s'\n", n, string(buffer[:n]))
	}

	fmt.Println()

	buffer = make([]byte, 5)
	s = strings.NewReader("Lbh penpxrq gur pbqr!")
	r = rot13Reader{s}
	for {
		n, e := r.Read(buffer)
		if e != nil {
			if e == io.EOF {
				fmt.Println("... end of input")
				return
			} else {
				fmt.Println("!!! error:", e)
				return
			}
		}
		fmt.Printf("%d read bytes: '%s'\n", n, string(buffer[:n]))
	}
}
