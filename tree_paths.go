package main

/*
Daily Coding Problem: Problem #110

Given a binary tree, return all paths from the root to leaves.

For example, given the tree

   1
  / \
 2   3
    / \
   4   5

it should return [[1, 2], [1, 3, 4], [1, 3, 5]].
*/

import (
	"binary_tree/tree"
	"fmt"
	"os"
)

// ValueAccumulator is an embellishment.
type ValueAccumulator []int

func (values *ValueAccumulator) collect(node tree.Node) {
	*values = append(*values, node.(*tree.NumericNode).Data)
}

// PathAccumulator keeps track of state (currentPath) during
// a traverse of a tree, and holds copies of the path as it
// existed when arriving at leaf nodes.
type PathAccumulator struct {
	currentPath []int
	paths       [][]int
}

// before, called in pre-order, add the node's
// value to the PathAccumulator's current path.
// If this is a leaf node, it copies the current path
// and appends that copy to the accumulator's list of paths.
func (pa *PathAccumulator) before(n tree.Node) {
	node := n.(*tree.NumericNode)
	pa.currentPath = append(pa.currentPath, node.Data)

	if node.Left == nil && node.Right == nil {
		path := make([]int, len(pa.currentPath))
		copy(path, pa.currentPath)
		pa.paths = append(pa.paths, path)
	}
}

// after, called in post-order, trims the node's
// value off the PathAccumulator's current path.
func (pa *PathAccumulator) after(node tree.Node) {
	pa.currentPath = pa.currentPath[:len(pa.currentPath)-1]
}

func main() {
	root := tree.CreateNumeric(os.Args[1:])

	if root != nil {
		var values ValueAccumulator
		tree.InorderTraverseVisit(root, values.collect)
		fmt.Printf("/* All values: %v */\n", values)

		p := &PathAccumulator{}
		tree.AllorderTraverseVisit(root, p.before, tree.NullVisitor, p.after)
		fmt.Printf("/* Paths:\n")
		for _, p := range p.paths {
			fmt.Printf("%v\n", p)
		}
		fmt.Printf("*/\n")
	}
}
