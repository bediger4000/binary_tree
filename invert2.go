package main

/*
 * Daily Coding Problem: Problem #83
 * Daily Coding Problem: Problem #596 [Medium]
 *
Invert a binary tree.

For example, given the following tree:

    a
   / \
  b   c
 / \  /
d   e f
should become:

  a
 / \
 c  b
 \  / \
  f e  d

*/

import (
	"fmt"
	"os"

	"binary_tree/tree"
)

func main() {

	root, err := tree.CreateFromString(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "parsing %q: %v\n", os.Args[1], err)
		return
	}

	if root != nil {
		fmt.Printf("digraph g1 {\n")
		fmt.Printf("subgraph cluster_0 {\n\tlabel=\"before\"\n")
		tree.DrawPrefixed(os.Stdout, root, "a")
		fmt.Printf("\n}\n")

		invert(root)

		fmt.Printf("subgraph cluster_1 {\n\tlabel=\"after\"\n")
		tree.DrawPrefixed(os.Stdout, root, "b")
		fmt.Printf("\n}\n")
		fmt.Printf("\n}\n")

	}
}

func invert(node *tree.StringNode) {
	if node == nil {
		return
	}
	invert(node.Left)
	invert(node.Right)
	node.Left, node.Right = node.Right, node.Left
}
