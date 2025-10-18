package main

import (
	"binary_tree/tree"
	"fmt"
	"os"
)

func main() {
	root, err := tree.CreateFromString(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "problem with %q: %v\n", os.Args[1], err)
		return
	}
	tree.Draw(root)
}
