package main

import (
	"binary_tree/tree"
	"fmt"
	"os"
)

/*
 * Daily Coding Problem: Problem #80
 * Given the root of a binary tree, return a deepest node. For example,
 * in the following tree, return d.
 *
 *     a
 *    / \
 *   b   c
 *  /
 * d
 *
 * A goroutine running a recursive func traverses the tree,
 * sends back leaf nodes and their depths.
 * main goroutine reads from channel, and keeps track of a
 * the deepest leaf node that the recursive func finds.
 */

type deepNode struct {
	node  tree.Node
	depth int
}

func main() {
	root := tree.CreateNumericFromString(os.Args[1])

	if root != nil {
		ch := make(chan *deepNode, 5)
		go findDepth(root, ch)
		maxDepth := 0
		var deepest tree.Node
		for d := range ch {
			if d.depth > maxDepth {
				maxDepth = d.depth
				deepest = d.node
			}
		}
		fmt.Printf("Max depth %d, node value at depth %v\n", maxDepth, deepest)
	}
}

// findDepth calls recursive function, then closes channel
// when that function returns.
func findDepth(root tree.Node, ch chan *deepNode) {
	realFindDepth(root, 0, ch)
	close(ch)
	return
}

// realFindDepth recursively visits all tree nodes, keeps track of
// depth, puts all leaf nodes and their depth on a channel.
func realFindDepth(node tree.Node, depth int, ch chan *deepNode) {
	if node.IsNil() {
		return
	}

	if node.LeftChild().IsNil() && node.RightChild().IsNil() {
		ch <- &deepNode{node: node, depth: depth}
		return
	}

	realFindDepth(node.LeftChild(), depth+1, ch)
	realFindDepth(node.RightChild(), depth+1, ch)
}
