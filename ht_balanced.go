package main

import (
	"binary_tree/tree"
	"fmt"
	"os"
)

/*
Daily Coding Problem: Problem #502 [Easy]

This problem was asked by PayPal.

Given a binary tree, determine whether or not it is height-balanced. A
height-balanced binary tree can be defined as one in which the heights
of the two subtrees of any node never differ by more than one.
*/
// There's no time or complexity requirement, so what's to prevent you
// from traversing the tree and finding all the subtree heights?

func main() {
	root := tree.CreateFromString(os.Args[1])

	phrase := " not"
	if Balanced(root) {
		phrase = ""
	}

	fmt.Printf("input tree is%s balanced\n", phrase)
}

// Balanced decides whether or not its argument is height-balanced. A
// height-balanced binary tree can be defined as one in which the
// heights of the two subtrees of any node never differ by more than
// one.
func Balanced(node tree.Node) bool {
	if node.IsNil() {
		return true
	}

	leftDepth := tree.FindDepth(node.LeftChild(), 0)
	fmt.Printf("Left depth %d\n", leftDepth)
	rightDepth := tree.FindDepth(node.RightChild(), 0)
	fmt.Printf("Right depth %d\n", rightDepth)

	depthDiff := leftDepth - rightDepth

	if depthDiff >= -1 && depthDiff <= 1 {
		return true
	}

	return false
}
