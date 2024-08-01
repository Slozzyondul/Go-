package binarytrees1

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	walkHelper(t, ch)
	close(ch)
}

// walkHelper is a recursive helper function for Walk.
func walkHelper(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	walkHelper(t.Left, ch)
	ch <- t.Value
	walkHelper(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for v1 := range ch1 {
		v2, ok := <-ch2
		if !ok || v1 != v2 {
			return false
		}
	}
	_, ok := <-ch2
	return !ok
}

func EquivalentBinaryTree() {
	// Test the Walk function
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	for v := range ch {
		fmt.Println(v)
	}

	// Test the Same function
	fmt.Println("Same(tree.New(1), tree.New(1)):", Same(tree.New(1), tree.New(1)))
	fmt.Println("Same(tree.New(1), tree.New(2)):", Same(tree.New(1), tree.New(2)))
}
