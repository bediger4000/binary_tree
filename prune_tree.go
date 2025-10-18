package main

/*
This question was asked by BufferBox.

Given a binary tree where all nodes are either 0 or 1, prune the tree so
that subtrees containing all 0s are removed.

For example, given the following tree:

   0
  / \
 1   0
    / \
   1   0
  / \
 0   0

should be pruned to:

   0
  / \
 1   0
    /
   1

We do not remove the tree at the root or its left child because it still
has a 1 as a descendant.
*/

import (
	"binary_tree/tree"
	"fmt"
	"os"
)

func prune(node *tree.StringNode) *tree.StringNode {
	if node == nil {
		return nil
	}

	node.Left = prune(node.Left)
	node.Right = prune(node.Right)
	if node.Left == nil && node.Right == nil {
		if node.Data == "0" {
			return nil
		}
		return node
	}
	return node
}

func main() {
	root, err := tree.CreateFromString(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "parsing %q: %v\n", os.Args[1], err)
		return
	}
	fmt.Printf("digraph g1 {\n")
	fmt.Printf("subgraph cluster_0 {\n\tlabel=\"before\"\n")
	tree.DrawPrefixed(os.Stdout, root, "orig")
	fmt.Printf("\n}\n")
	pruned := prune(root)
	fmt.Printf("subgraph cluster_1 {\n\tlabel=\"after\"\n")
	tree.DrawPrefixed(os.Stdout, pruned, "prune")
	fmt.Printf("\n}\n")
	fmt.Printf("\n}\n")
}
