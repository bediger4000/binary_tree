package main

import (
	"binary_tree/tree"
	"bytes"
	"flag"
	"fmt"
)

func main() {

	n := flag.Int("n", 12, "number of nodes")
	max := flag.Int("m", 100, "maximum node value")
	graphvizOutput := flag.Bool("g", false, "GraphViz output on stdout")
	flag.Parse()

	hdr, trailer := "", ""
	root := tree.RandomValueTree(*max, *n, true)
	if *graphvizOutput {
		hdr, trailer = "/* ", " */"
		tree.Draw(root)
	}
	var buffer bytes.Buffer
	tree.Printf(&buffer, root)
	fmt.Printf("%s%s%s\n", hdr, buffer.String(), trailer)
}
