package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	var walker func(t *tree.Tree)
	walker = func(t *tree.Tree) {
		if t == nil {
			return
		}
		walker(t.Left)
		ch <- t.Value
		walker(t.Right)
	}
	walker(t)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	t1_ch := make(chan int, 10)
	t2_ch := make(chan int, 10)

	go Walk(t1, t1_ch)
	go Walk(t2, t2_ch)

	for {
		v1, ok1 := <-t1_ch
		v2, ok2 := <-t2_ch
		if ok1 != ok2 || v1 != v2 {
			return false
		}
		if !ok1 {
			return true
		}
	}
}

func main() {
	t1 := tree.New(10)
	t2 := tree.New(10)
	t3 := tree.New(100)
	t4 := tree.New(100)
	same := Same(t1, t2)
	fmt.Println("t1 is same as t2     PASS: ", same)
	same = Same(t1, t3)
	fmt.Println("t1 is not same as t3 PASS: ", !same)
	same = Same(t3, t4)
	fmt.Println("t3 is same as t4     PASS: ", same)
}
