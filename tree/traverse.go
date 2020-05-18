package tree

import "fmt"

// PreorderTraverse prints nodes' values on stdout in pre-order.
func PreorderTraverse(node *Node) {
	if node == nil {
		return
	}
	fmt.Printf(" %d", node.Data)
	PreorderTraverse(node.Left)
	PreorderTraverse(node.Right)
}

// InorderTraverse prints nodes' values on stdout in order.
func InorderTraverse(node *Node) {
	if node == nil {
		return
	}
	InorderTraverse(node.Left)
	fmt.Printf(" %d", node.Data)
	InorderTraverse(node.Right)
}
