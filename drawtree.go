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
			root = tree.Insert(root, val)
		} else {
			fmt.Fprintf(os.Stderr, "Problem with %q: %s\n", str, err)
		}
	}

	tree.Draw(root)
}
