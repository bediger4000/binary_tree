package main

/*
 * Given the root to a binary search tree,
 * find the second largest node in the tree.
 */

import (
	"binary_tree/tree"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("find 2nd largest value in a binary search tree\n")
		fmt.Printf("Usage: %s 1 3 6 ...'\n", os.Args[0])
		return
	}
	root := tree.CreateNumeric(os.Args[1:])

	if root.Left == nil && root.Right == nil {
		fmt.Printf("Single-node tree does not have a 2nd largest value\n")
		return
	}

	f, v := findSecondLargest(root)
	secondLargest := v
	if f {
		// handle the case of a 2-node tree, root and its left child
		secondLargest = root.Data
	}
	fmt.Printf("Second largest value in tree: %d\n", secondLargest)
}

func findSecondLargest(node *tree.NumericNode) (bool, int) {
	if node.Right == nil {
		// This is the largest-valued node
		// If it does not have a left child, the parent
		// has the 2nd largest value. Otherwise, the left
		// child has the 2nd largest value.
		if node.Left != nil {
			return false, node.Left.Data
		}
		return true, node.Data
	}

	f, v := findSecondLargest(node.Right)
	if f {
		// node.Right does not have a left child,
		// so this node has 2nd largest value
		v = node.Data
	}
	return false, v
}
