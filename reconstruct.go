package main

import (
	"binary_tree/tree"
)

// Return:
// -1 if value1 < value2
//  0 if value1 == value2
//  1 if value1 > value2
func imposeOrder(value1, value2 string, valueorder []string) int {
	if value1 == value2 {
		return 0
	}

	for _, val := range valueorder {
		if val == value1 {
			return -1
		}
		if val == value2 {
			return 1
		}
	}

	// Should probably panic here
	return 0 // almost certainly wrong.
}

func insert(node *tree.StringNode, value string, valueorder []string) *tree.StringNode {

	if node == nil {
		return &tree.StringNode{Data: value}
	}

	n := &(node.Left)
	if imposeOrder(value, node.Data, valueorder) != -1 {
		n = &(node.Right)
	}
	*n = insert(*n, value, valueorder)
	return node
}

func main() {

	var root *tree.StringNode

	preorder := []string{"a", "b", "c", "d", "e", "f", "g"}
	inorder := []string{"d", "b", "e", "a", "f", "c", "g"}

	for _, v := range preorder {
		root = insert(root, v, inorder)
	}

	if root != nil {
		tree.Draw(root)
	}
}
