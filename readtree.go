package main

import (
	"binary_tree/tree"
	"fmt"
	"os"
)

func main() {
	outputGraphViz := false
	stringrep := os.Args[1]
	if stringrep == "-g" {
		outputGraphViz = true
		stringrep = os.Args[2]
	}

	root := tree.CreateByParsing(stringrep)
	if outputGraphViz {
		tree.Draw(root)
		return
	}
	tree.PrintStringTree(root)
	fmt.Println()
}
