package main

import (
	"binary_tree/tree"
	"log"
	"os"
	"strconv"
)

func main() {

	n := 12
	if len(os.Args) > 1 {
		var err error
		n, err = strconv.Atoi(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
	}

	root := tree.RandomValueTree(100, n, true)
	tree.Draw(root)
}
