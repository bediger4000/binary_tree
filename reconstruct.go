package main

import (
	"binary_tree/tree"
)

// Return:
// -1 if value1 < value2
//  0 if value1 == value2
//  1 if value1 > value2
func imposeOrder(value1, value2 string, valueorder map[string]int) int {
	if value1 == value2 {
		return 0
	}

	if valueorder[value1] < valueorder[value2] {
		return -1
	}

	return 1
}

func insert(node *tree.StringNode, value string, valueorder map[string]int) *tree.StringNode {

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

	ordering := make(map[string]int)
	for i, str := range inorder {
		ordering[str] = i
	}

	for _, v := range preorder {
		root = insert(root, v, ordering)
	}

	if root != nil {
		tree.Draw(root)
	}
}
