package main

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"golang.org/x/tour/tree"
)

// XXX:
// type Tree struct {
//     Left  *Tree
//     Value int
//     Right *Tree
// }

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	ch <- t.Value
	Walk(t.Left, ch)
	Walk(t.Right, ch)
}

func EqualUnsortedIntArray(a1, a2 []int) bool {
	less := func(a, b int) bool { return a < b }
	return cmp.Diff(a1, a2, cmpopts.SortSlices(less)) == ""
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	var t1_values []int
	var t2_values []int
	t1_ch := make(chan int, 10)
	t2_ch := make(chan int, 10)

	go Walk(t1, t1_ch)
	go Walk(t2, t2_ch)

	for {
		select {
		case v := <-t1_ch:
			t1_values = append(t1_values, v)
		case v := <-t2_ch:
			t2_values = append(t2_values, v)
		default:
			if len(t1_values) == 10 && len(t2_values) == 10 {
				return EqualUnsortedIntArray(t1_values, t2_values)
			}
		}
	}
}

func main() {
	t1 := tree.New(10)
	t2 := tree.New(10)
	t3 := tree.New(100)
	same := Same(t1, t2)
	fmt.Println("t1 is same as t2     PASS: ", same == true)
	same = Same(t1, t3)
	fmt.Println("t1 is not same as t3 PASS: ", same == false)
}
