package main

/*

This problem was asked by Yahoo.

Recall that a full binary tree is one in which each node is either a
leaf node,
or has two children.
Given a binary tree,
convert it to a full one by removing nodes with only one child.

For example, given the following tree:

			 0
		  /     \
		1         2
	  /            \
	3                 4
	  \             /   \
		5          6     7

You should convert it to:

		 0
	  /     \
	5         4
			/   \
		   6     7

*/

import (
	"binary_tree/tree"
	"fmt"
	"os"
)

func main() {
	drawGraph := false
	stringRep := os.Args[1]
	if os.Args[1] == "-g" {
		drawGraph = true
		stringRep = os.Args[2]
	}

	root, err := tree.CreateNumericFromString(stringRep)
	if err != nil {
		fmt.Fprintf(os.Stderr, "parsing %q: %v\n", stringRep, err)
		return
	}

	root = convert(root)

	if drawGraph {
		fmt.Print("/* ")
	}
	tree.Print(root)
	if drawGraph {
		fmt.Print(" */")
	}
	fmt.Println()

	if drawGraph {
		tree.Draw(root)
	}
}

func convert(node *tree.NumericNode) *tree.NumericNode {
	if node == nil {
		return nil
	}

	node.Left = convert(node.Left)
	node.Right = convert(node.Right)

	if node.Left != nil {
		if node.Right == nil {
			// left non-nil, right nil
			return node.Left
		}
		// both left and right child non-nil
		return node
	}
	if node.Right != nil {
		if node.Left == nil {
			// right non-nil, left nil
			return node.Right
		}
	}
	// both left and right nil
	return node
}
