package main

import (
	"binary_tree/tree"
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
	root1 := tree.CreateNumericFromString(os.Args[1])
	root2 := tree.CreateNumericFromString(os.Args[2])

	merged := merge(root1, root2)
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
