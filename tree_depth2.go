package main

import (
	"binary_tree/tree"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var root *tree.Node

	for _, str := range os.Args[1:] {
		val, err := strconv.Atoi(str)

		if err == nil {
			fmt.Printf("insert %d\n", val)
			// tree.Insert creates a binary search tree
			root = tree.Insert(root, val)
		} else {
			fmt.Printf("Problem with %q: %s\n", str, err)
		}
	}

	if root != nil {
		depth, node := tree.FindDepth2(root, 0)
		fmt.Printf("Max depth %d, node value at depth %d\n", depth, node.Value())
	}
}
