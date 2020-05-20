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

	// construct BST from all the remaining command line values
	root := tree.CreateNumeric(os.Args[2:])

	// find parent of particular node by value
	parentNode := findParent(root, targetNodeValue)
	if parentNode == nil {
		log.Fatalf("Did not find parent node of value %d\n", targetNodeValue)
	}
	fmt.Printf("parent of node of interest has value %d\n", parentNode.Data)

	// find grandparent of particular node
	grandparentNode := findParent(root, parentNode.Data)
	if grandparentNode == nil {
		log.Fatalf("Did not find grandparent node of value %d\n", targetNodeValue)
	}
	fmt.Printf("grandparent of node of interest has value %d\n", grandparentNode.Data)

	// find children of grandparent - one of these is
	// parent of cousins
	var uncle *tree.Node
	if grandparentNode.Left != nil && grandparentNode.Left.Data == parentNode.Data {
		uncle = grandparentNode.Right
	}
	if uncle == nil && grandparentNode.Right != nil {
		uncle = grandparentNode.Left
	}

	if uncle == nil {
		fmt.Printf("No uncle node, so no cousins\n")
		return
	}

	// find at most 2 cousins' values
	fmt.Printf("uncle node has value %d\n", uncle.Data)
	if uncle.Left != nil {
		fmt.Printf("Cousin has value %d\n", uncle.Left.Data)
	}
	if uncle.Right != nil {
		fmt.Printf("Cousin has value %d\n", uncle.Right.Data)
	}

}

// Find performs binary search on argument node
// and its sub-trees. Since my trees are all Binary Search Trees,
// this can short-circuit a mere traversal of the entire tree.
// This is different than the example given, but I'm using what
// I've got available.
func findParent(node *tree.Node, value int) *tree.Node {
	if node == nil {
		return nil
	}
	if node.Left != nil {
		if value == node.Left.Data {
			return node
		}
	}
	if node.Right != nil {
		if value == node.Right.Data {
			return node
		}
	}
	if node.Data > value {
		return findParent(node.Left, value)
	}
	return findParent(node.Right, value)
}
