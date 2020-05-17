package tree

import "fmt"

// PreorderTraverse prints nodes' values on stdout in pre-order.
func PreorderTraverse(node *Node) {
	if node == nil {
		return
	}
	fmt.Printf(" %d", node.data)
	PreorderTraverse(node.left)
	PreorderTraverse(node.right)
}

// InorderTraverse prints nodes' values on stdout in order.
func InorderTraverse(node *Node) {
	if node == nil {
		return
	}
	InorderTraverse(node.left)
	fmt.Printf(" %d", node.data)
	InorderTraverse(node.right)
}
