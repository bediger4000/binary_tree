package main

import (
	"binary_tree/tree"
	"fmt"
	"log"
	"os"
	"strconv"
)

/*
Two nodes in a binary tree can be called cousins if they are on the same
level of the tree but have different parents. For example, in the
following diagram 4 and 6 are cousins.

    1
   / \
  2   3
 / \   \
4   5   6

Given a binary tree and a particular node, find all cousins of that
node.
*/

func main() {
	// Read value of "particular node"
	targetNodeValue, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Node of interest has value %d\n", targetNodeValue)

	// Construct BST from all the remaining command line values
	// Using a BST because by entering nodes in bread-first order,
	// I can get any shape tree I want.
	root := tree.CreateNumericFromString(os.Args[2])

	d := findDepth(root, targetNodeValue, 0)

	fmt.Printf("Particular node of value %d at depth %d\n",
		targetNodeValue, d)

	nodesAtDepth(root, targetNodeValue, d, 0)
}

// findDepth returns the depth in the tree (root has depth 0)
// of a node with data value of value (argument),
// or -1 if value not found
func findDepth(node *tree.NumericNode, value int, depth int) int {
	if node == nil {
		return -1
	}
	if node.Data == value {
		return depth
	}
	d := findDepth(node.Left, value, depth+1)
	if d > -1 {
		return d
	}
	return findDepth(node.Right, value, depth+1)
}

// nodesAtDepth prints the value of nodes at depth desiredDepth,
// but not the node with data value cousin. That's the "particular node"
// itself.
func nodesAtDepth(node *tree.NumericNode, cousin, desiredDepth, depth int) {
	if node == nil {
		return
	}
	if node.Data == cousin {
		return
	}
	if desiredDepth == depth {
		fmt.Printf("Cousin node %d\n", node.Data)
		return
	}
	nodesAtDepth(node.Left, cousin, desiredDepth, depth+1)
	nodesAtDepth(node.Right, cousin, desiredDepth, depth+1)
}
