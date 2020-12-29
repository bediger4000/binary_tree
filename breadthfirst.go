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

	root := tree.CreateNumericFromString(os.Args[1])

	stack := &tree.Stack{}

	stack.Push(root)

	for !stack.Empty() {
		n := stack.Dequeue().(*tree.NumericNode)
		fmt.Printf("%d ", n.Data)
		stack.Push(n.Left)
		stack.Push(n.Right)
	}

	fmt.Println()
}
