package main

import (
	"flag"
	"fmt"
	"os"

	"binary_tree/tree"
)

func main() {
	graphvizOutput := flag.Bool("g", false, "GraphViz output on stdout")
	flag.Parse()

	if flag.NArg() == 0 {
		fmt.Printf("Need string represenation of tree on command line\n")
		return
	}

	stringrep := flag.Args()[0]

	root, err := tree.CreateNumericFromString(stringrep)
	if err != nil {
		fmt.Fprintf(os.Stderr, "parsing %q: %v\n", stringrep, err)
		return
	}
	fmt.Printf("/* Original:  %q\n", stringrep)
	fmt.Print(`   As parsed: "`)
	tree.Print(root)
	fmt.Print("\"\n */\n")
	if *graphvizOutput {
		tree.Draw(root)
	}
	if tree.BstProperty(root) {
		fmt.Printf("/* This is a binary search tree */\n")
	}
}
