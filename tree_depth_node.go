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
	root := tree.CreateNumericFromString(os.Args[1])

	if root != nil {
		depth, node := findDepth(root)
		fmt.Printf("Max depth %d, node value at depth %s\n", depth, node)
	}
}

// findDepth using more-or-less generic tree.AllorderTraverseVisit,
// and 3 closures using local variables as the visitor functions.
// No globals so this is thread safe.
func findDepth(root tree.Node) (maxDepth int, deepNode tree.Node) {

	var depth int

	// you could also implement a struct with depth and a Node element,
	// "detach" a function using the "object.Method" syntax to get the
	// correct function signature, and set that as the visitor functions.
	preorderfun := func(node tree.Node) {
		depth += 1
	}
	postorderfun := func(node tree.Node) {
		depth -= 1
	}
	inorderfun := func(node tree.Node) {
		if depth > maxDepth {
			maxDepth = depth
			deepNode = node
		}
	}

	tree.AllorderTraverseVisit(root, preorderfun, inorderfun, postorderfun)

	return
}
