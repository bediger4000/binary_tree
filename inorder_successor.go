package main

import (
	"binary_tree/tree"
	"fmt"
	"log"
	"os"
	"strconv"
)

/*
Daily Coding Problem: Problem #133

This problem was asked by Amazon.

Given a node in a binary tree, return the next bigger element, also
known as the inorder successor.

For example, the inorder successor of 22 is 30.

   10
  /  \
 5    30
     /  \
   22    35

Node   Inorder Successor
  5         10
 10         22
 30         35
 22         30
 35          -

*/

type FoundIt int

const (
	FoundValue     FoundIt = 0
	FoundSuccessor FoundIt = iota
	NotFound       FoundIt = iota
)

func main() {
	given, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	root := tree.CreateNumeric(os.Args[2:])

	successorValue, clue := searchFor(root, given, NotFound)
	switch clue {
	case FoundValue:
		fmt.Printf("%d does not have an in-order successor\n", given)
	case FoundSuccessor:
		fmt.Printf("Inorder successor of %d is %d\n", given, successorValue)
	case NotFound:
		fmt.Printf("Did not find node with value %d\n", given)
	}
}

func searchFor(node *tree.NumericNode, given int, what FoundIt) (int, FoundIt) {
	if node == nil {
		return 0, what
	}

	if what == FoundValue {
		// this node is the Right child of the given value node.
		// it's the inorder successor
		return node.Data, FoundSuccessor
	}

	v, hint := searchFor(node.Left, given, what)

	if hint == FoundSuccessor {
		return v, FoundSuccessor
	}

	if hint == FoundValue {
		// this node is the parent of the given value node,
		// it's the inorder successor.
		return node.Data, FoundSuccessor
	}

	if node.Data == given {
		// this node has the given value, whatever node gets visited next is
		// the inorder successor
		what = FoundValue
	}

	return searchFor(node.Right, given, what)
}
