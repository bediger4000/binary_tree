package main

import (
	"binary_tree/tree"
	"os"
)

func main() {
	root := tree.CreateNumeric(os.Args[1:])
	tree.Draw(root)
}
