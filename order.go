package main

import (
	"binary_tree/tree"
	"fmt"
	"os"
)

func main() {
	root := tree.CreateNumeric(os.Args[1:])
	phrase := "has"
	if !tree.BstProperty(root) {
		phrase = "doesn't have"
	}
	fmt.Fprintf(os.Stderr, "tree %s binary search property\n", phrase)
	tree.InorderTraverse(root)
	fmt.Println()
	tree.Print(root)
	fmt.Println()
}
