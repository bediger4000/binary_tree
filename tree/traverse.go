package tree

import "fmt"

// PreorderTraverse prints nodes' values on stdout in pre-order.
func PreorderTraverse(node Node) {
	if node.isNil() {
		return
	}
	fmt.Printf(" %s", node)
	PreorderTraverse(node.leftChild())
	PreorderTraverse(node.rightChild())
}

// InorderTraverseVisit does a traverse of a binary tree,
// callling a function in-order at every node
func InorderTraverseVisit(node Node, fn VisitorFunc) {
	for !node.isNil() {
		InorderTraverseVisit(node.leftChild(), fn)
		fn(node)
		node = node.rightChild()
	}
}

// AllorderTraverseVisit does a traverse of a binary tree,
// callling a function pre-order, in-order and post-order at every node
func AllorderTraverseVisit(node Node, preorderfn, inorderfn, postorderfn VisitorFunc) {
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
func PreorderTraverseVisit(node Node, fn VisitorFunc) {
	if node == nil {
		return
	}
	fn(node)
	InorderTraverseVisit(node.leftChild(), fn)
	InorderTraverseVisit(node.rightChild(), fn)
}

// InorderTraverse prints nodes' values on stdout in order.
func InorderTraverse(node Node) {
	AllorderTraverseVisit(node, NullVisitor, printNode, NullVisitor)
}

// InorderTraverseTail prints nodes' values on stdout in order,
// but only recurses on left child nodes - tail call optimization.
func InorderTraverseTail(node Node) {
	for !node.isNil() {
		InorderTraverseTail(node.leftChild())
		fmt.Printf("%s ", node)
		node = node.rightChild()
	}
}

// printNode relies on node fitting interface fmt.Stringer, too.
func printNode(node Node) {
	fmt.Printf(" %s", node)
}
