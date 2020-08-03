package main

import (
	"binary_tree/tree"
	"fmt"
	"os"
)

func main() {
	root := tree.CreateNumeric(os.Args[1:])

	var d tree.Depth

	if root != nil {
		tree.FindDepth1(root, 0, &d)
		fmt.Printf("Max depth %d, node value at depth %d\n", d.Depth, d.Node.Value())
	}
}
