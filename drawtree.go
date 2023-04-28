package main

import (
	"binary_tree/tree"
	"os"
)

func main() {
	root := tree.CreateFromString(os.Args[1])
	tree.Draw(root)
}
