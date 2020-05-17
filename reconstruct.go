package main

import (
	"binary_tree/tree"
	"fmt"
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

	n := &(node.left)
	if imposeOrder(value, node.Data, valueorder) != -1 {
		n = &(node.right)
	}
	*n = insert(*n, value, valueorder)
	return node
}

func drawTree(node *TreeNode) {
	fmt.Printf("Node%p [label=\"%s\"];\n", node, node.data)
	if node.left != nil {
		drawTree(node.left)
		fmt.Printf("Node%p -> Node%p;\n", node, node.left)
	} else {
		fmt.Printf("Node%pL [shape=\"point\"];\n", node)
		fmt.Printf("Node%p -> Node%pL;\n", node, node)
	}
	if node.right != nil {
		drawTree(node.right)
		fmt.Printf("Node%p -> Node%p;\n", node, node.right)
	} else {
		fmt.Printf("Node%pR [shape=\"point\"];\n", node)
		fmt.Printf("Node%p -> Node%pR;\n", node, node)
	}
}

func main() {

	var root *TreeNode

	preorder := []string{"a", "b", "c", "d", "e", "f", "g"}
	inorder := []string{"d", "b", "e", "a", "f", "c", "g"}

	for _, v := range preorder {
		root = insert(root, v, inorder)
	}

	if root != nil {
		fmt.Printf("digraph g {\n")
		drawTree(root)
		fmt.Printf("\n}\n")
	}
}
