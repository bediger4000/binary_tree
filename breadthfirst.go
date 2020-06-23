package main

/*
 * Traverse a tree breadth-first, iteratively.
 */

import (
	"binary_tree/tree"
	"fmt"
	"os"
)

func main() {
	root := tree.CreateNumeric(os.Args[1:])
	phrase := "has"
	if !tree.BstProperty(root) {
		phrase = "doesn't have"
	}
	fmt.Fprintf(os.Stderr, "tree %s binary search property\n", phrase)

	stack := &tree.Stack{}

	stack.Push(root)

	for !stack.Empty() {
		n := stack.Dequeue()
		fmt.Printf("%d ", n.(*tree.NumericNode).Data)
		stack.Push(n.(*tree.NumericNode).Left)
		stack.Push(n.(*tree.NumericNode).Right)
	}
	fmt.Println()
}
