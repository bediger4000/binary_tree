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

// InorderTraverseVisit does a traverse of a binary tree,
// callling a function in-order at every node
func InorderTraverseVisit(node *Node, fn VisitorFunc) {
	if node == nil {
		return
	}
	InorderTraverseVisit(node.Left, fn)
	fn(node)
	InorderTraverseVisit(node.Right, fn)
}

// AllorderTraverseVisit does a traverse of a binary tree,
// callling a function pre-order, in-order and post-order at every node
func AllorderTraverseVisit(node *Node, preorderfn, inorderfn, postorderfn VisitorFunc) {
	if node == nil {
		return
	}
	preorderfn(node)
	AllorderTraverseVisit(node.Left, preorderfn, inorderfn, postorderfn)
	inorderfn(node)
	AllorderTraverseVisit(node.Right, preorderfn, inorderfn, postorderfn)
	postorderfn(node)
}

// PreorderTraverseVisit does a traverse of a binary tree,
// callling a function pre-order at every node
func PreorderTraverseVisit(node *Node, fn VisitorFunc) {
	if node == nil {
		return
	}
	fn(node)
	InorderTraverseVisit(node.Left, fn)
	InorderTraverseVisit(node.Right, fn)
}

// InorderTraverse prints nodes' values on stdout in order.
func InorderTraverse(node *Node) {
	AllorderTraverseVisit(node, NullVisitor, printNode, NullVisitor)
}
func printNode(node *Node) {
	fmt.Printf(" %d", node.Data)
}
