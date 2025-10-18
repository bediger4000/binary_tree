package main

/*
 */

import (
	"binary_tree/tree"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Determine whether the input tree is a valid binary search tree\n")
		fmt.Printf("Usage: %s '(0 (-1) (2))'\n", os.Args[0])
		return
	}
	root, err := tree.CreateNumericFromString(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "parsing %q: %v\n", os.Args[1], err)
		return
	}

	tree.Print(root)
	if tree.BstProperty(root) {
		fmt.Printf(" is a valid binary search tree\n")
		return
	}
	fmt.Printf(" is NOT a valid binary search tree\n")

}
