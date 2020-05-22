package main

/*
 * Exercize func tree.CreateFromString, which parses a single string
 * like "(abc(ghi()(jkl))(def(pork)(beans)))"
 * and turns it into a binary tree.
 */

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

	root := tree.CreateFromString(stringrep)
	if outputGraphViz {
		tree.Draw(root)
		return
	}
	tree.Print(root)
	fmt.Println()
}
