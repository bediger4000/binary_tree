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
			root = tree.Insert(root, val)
		} else {
			fmt.Printf("Problem with %q: %s\n", str, err)
		}
	}

	var d tree.Depth

	if root != nil {
		tree.FindDepth1(root, 0, &d)
		fmt.Printf("Max depth %d, node value at depth %d\n", d.Depth, d.Node.Value())
	}
}
