package main

import (
	"binary_tree/tree"
	"fmt"
	"os"
)

func main() {
	root, err := tree.CreateNumericFromString(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "parsing %q: %v\n", os.Args[1], err)
		return
	}

	tree.Print(root)
	if !tree.BstProperty(root) {
		fmt.Printf(" is NOT a valid binary search tree\n")
		return
	} else {
		fmt.Println()
	}
	largestValue(root)
}

func largestValue(node *tree.NumericNode) bool {
	if node.Right == nil {
		fmt.Printf("Largest value %d\n", node.Data)
		if node.Left != nil {
			fmt.Printf("Second largest value %d\n", node.Left.Data)
			return false
		}
		return true
	}
	foundit := largestValue(node.Right)
	if foundit {
		fmt.Printf("Second largest value %d\n", node.Data)
	}
	return false
}
