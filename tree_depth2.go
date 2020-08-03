package main

import (
	"binary_tree/tree"
	"fmt"
	"os"
)

/*
 * Daily Coding Problem: Problem #80
 * Given the root of a binary tree, return a deepest node. For example,
 * in the following tree, return d.
 *
 *     a
 *    / \
 *   b   c
 *  /
 * d
 */

func main() {
	root := tree.CreateNumeric(os.Args[1:])

	if root != nil {
		depth, node := tree.FindDepth2(root, 0)
		fmt.Printf("Max depth %d, node value at depth %d\n", depth, node.Value())
	}
}
