package main

import (
	"code.google.com/p/go-tour/tree"
	"fmt"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	var f (func(node *tree.Tree))
	f = func(node *tree.Tree) {
		if node.Left != nil {
			f(node.Left)
		}

		ch <- node.Value

		if node.Right != nil {
			f(node.Right)
		}
	}
	f(t)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
//func Same(t1, t2 *tree.Tree) bool

func main() {
	ch := make(chan int)

	// walk and send values to channel
	go Walk(tree.New(1), ch)

	// show value sent to channel
	for i := range ch {
		fmt.Println(i)
	}
}
