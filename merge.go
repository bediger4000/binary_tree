package main

import (
	"binary_tree/tree"
	"flag"
	"fmt"
	"os"
)

/*
Daily Coding Problem: Problem #422 [Easy]

This problem was asked by Salesforce.

Write a program to merge two binary trees.
Each node in the new tree should hold a value equal to the sum of the values of
the corresponding nodes of the input trees.

If only one input tree has a node in a given position,
the corresponding node in the new tree should match that input node.
*/

func main() {
	graphViz := flag.Bool("g", false, "GraphViz dot output on stdout")
	flag.Parse()

	root1, err := tree.CreateNumericFromString(flag.Args()[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "parsing %q: %v\n", flag.Args()[0], err)
		return
	}

	root2, err := tree.CreateNumericFromString(flag.Args()[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "parsing %q: %v\n", flag.Args()[1], err)
		return
	}

	merged := merge(root1, root2)

	if *graphViz {

		fmt.Printf("digraph g1 {\n")
		fmt.Printf("subgraph cluster_0 {\n\tlabel=\"1st\"\n")
		tree.DrawPrefixed(os.Stdout, root1, "a")
		fmt.Printf("\n}\n")
		fmt.Printf("subgraph cluster_1 {\n\tlabel=\"2nd\"\n")
		tree.DrawPrefixed(os.Stdout, root2, "b")
		fmt.Printf("\n}\n")
		fmt.Printf("subgraph cluster_2 {\n\tlabel=\"merged\"\n")
		tree.DrawPrefixed(os.Stdout, merged, "m")
		fmt.Printf("\n}\n")
		fmt.Printf("\n}\n")

		return
	}
	tree.Print(merged)
}

func merge(node1, node2 *tree.NumericNode) *tree.NumericNode {
	if node1 == nil && node2 == nil {
		return nil
	}
	if node1 != nil && node2 == nil {
		return &tree.NumericNode{
			Data:  node1.Data,
			Left:  merge(node1.Left, nil),
			Right: merge(node1.Right, nil),
		}
	}
	if node1 == nil && node2 != nil {
		return &tree.NumericNode{
			Data:  node2.Data,
			Left:  merge(nil, node2.Left),
			Right: merge(nil, node2.Right),
		}
	}
	return &tree.NumericNode{
		Data:  node1.Data + node2.Data,
		Left:  merge(node1.Left, node2.Left),
		Right: merge(node1.Right, node2.Right),
	}
}
