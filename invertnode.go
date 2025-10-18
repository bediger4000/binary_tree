package main

// invert (left-right mirror) a tree made with interface Node
// just to see if it can be done. It can.

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

	str := os.Args[1]
	var root tree.Node
	var err error
	if str == "-n" {
		str = os.Args[2]
		root, err = tree.CreateNumericFromString(str)
	} else {
		root, err = tree.CreateFromString(str)
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "parsing %q: %v\n", str, err)
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

// invert is a type-independent recursive function that (destructively) mirrors
// a tree left-to-right.
func invert(node tree.Node) {
	if node.IsNil() {
		return
	}
	invert(node.LeftChild())
	invert(node.RightChild())
	left := node.LeftChild()
	right := node.RightChild()
	node.SetLeftChild(right)
	node.SetRightChild(left)
}
