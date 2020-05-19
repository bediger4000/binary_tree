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
	fmt.Printf("tree %s binary search property\n", phrase)
	tree.InorderTraverse(root)
	fmt.Println()
}
