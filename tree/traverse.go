package tree

import "fmt"

// PreorderTraverse prints nodes' values on stdout in pre-order.
func PreorderTraverse(node drawable) {
	if node.isNil() {
		return
	}
	fmt.Printf(" %s", node)
	PreorderTraverse(node.leftChild())
	PreorderTraverse(node.rightChild())
}

// InorderTraverseVisit does a traverse of a binary tree,
// callling a function in-order at every node
func InorderTraverseVisit(node drawable, fn VisitorFunc) {
	if node.isNil() {
		return
	}
	InorderTraverseVisit(node.leftChild(), fn)
	fn(node)
	InorderTraverseVisit(node.rightChild(), fn)
}

// AllorderTraverseVisit does a traverse of a binary tree,
// callling a function pre-order, in-order and post-order at every node
func AllorderTraverseVisit(node drawable, preorderfn, inorderfn, postorderfn VisitorFunc) {
	if node.isNil() {
		return
	}
	preorderfn(node)
	AllorderTraverseVisit(node.leftChild(), preorderfn, inorderfn, postorderfn)
	inorderfn(node)
	AllorderTraverseVisit(node.rightChild(), preorderfn, inorderfn, postorderfn)
	postorderfn(node)
}

// PreorderTraverseVisit does a traverse of a binary tree,
// callling a function pre-order at every node
func PreorderTraverseVisit(node drawable, fn VisitorFunc) {
	if node == nil {
		return
	}
	fn(node)
	InorderTraverseVisit(node.leftChild(), fn)
	InorderTraverseVisit(node.rightChild(), fn)
}

// InorderTraverse prints nodes' values on stdout in order.
func InorderTraverse(node drawable) {
	AllorderTraverseVisit(node, NullVisitor, printNode, NullVisitor)
}

// printNode relies on node fitting interface fmt.Stringer, too.
func printNode(node drawable) {
	fmt.Printf(" %s", node)
}
