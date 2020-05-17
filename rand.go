package main

import "binary_tree/tree"

func main() {

	root := tree.RandomValueTree(100, 12, true)
	tree.Draw(root)
}
