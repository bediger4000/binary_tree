package main

import (
	"binary_tree/tree"
)

// Return:
// negative if value1 < value2
//  0       if value1 == value2
// positive if value1 > value2
func imposeOrder(value1, value2 string, valueorder map[string]int) int {
	return valueorder[value1] - valueorder[value2]
}

// insert as a different function than tree.Insert necessary
// because (a) StringNode has string data, (b) the comparison
// happens based on some given order.
// It would be possible to write a more general tree.Insert
// that has a comparsion function: func comparison(n1, n2 comparable) int
// that returned negative/zero/positive based on a function of
// interface comparable
func insert(node *tree.StringNode, value string, valueorder map[string]int) *tree.StringNode {

	if node == nil {
		return &tree.StringNode{Data: value}
	}

	n := &(node.Left)
	if imposeOrder(value, node.Data, valueorder) > 0 {
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
