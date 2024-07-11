package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

// XXX: go get golang.org/x/tour/wc

func WordCount(s string) map[string]int {
	// DONE: Implement WordCount. It should return a map of the counts of each
	// “word” in the string s. The wc.Test function runs a test suite against the
	// provided function and prints success or failure.
	// You might find strings.Fields helpful.
	ret := make(map[string]int)
	for _, key := range strings.Fields(s) {
		v, ok := ret[key]
		if ok {
			ret[key] = v + 1
		} else {
			ret[key] = 1
		}
	}
	return ret
}

func main() {
	wc.Test(WordCount)
}
