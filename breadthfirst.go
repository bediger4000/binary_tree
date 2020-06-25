package main

/*
 * Daily Coding Problem: Problem #107
 * This problem was asked by Microsoft.
 *
 * Print the nodes in a binary tree level-wise. For example, the
 * following should print 1, 2, 3, 4, 5.
 *
 *     1
 *    / \
 *   2   3
 *  / \
 * 4   5
 *
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
