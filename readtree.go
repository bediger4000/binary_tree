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

	root, err := tree.CreateFromString(stringrep)
	if err != nil {
		fmt.Fprintf(os.Stderr, "parsing %q: %v\n", stringrep, err)
		return
	}
	if outputGraphViz {
		tree.Draw(root)
		return
	}
	tree.Print(root)
	fmt.Println()
}
