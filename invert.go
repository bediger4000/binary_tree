package main

/*
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

	root := tree.RandomValueTree(100, 15, true)

	if root != nil {
		fmt.Printf("digraph g1 {\n")
		fmt.Printf("subgraph cluster_0 {\n\tlabel=\"before\"\n")
		tree.DrawPrefixed(os.Stdout, root, "a")
		fmt.Printf("\n}\n")
		root.Invert()
		fmt.Printf("subgraph cluster_1 {\n\tlabel=\"after\"\n")
		tree.DrawPrefixed(os.Stdout, root, "b")
		fmt.Printf("\n}\n")
		fmt.Printf("\n}\n")

	}
}
