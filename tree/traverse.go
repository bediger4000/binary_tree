package tree

import "fmt"

// PreorderTraverse prints nodes' values on stdout in pre-order.
func PreorderTraverse(node Node) {
	if node.IsNil() {
		return
	}
	fmt.Printf(" %s", node)
	PreorderTraverse(node.LeftChild())
	PreorderTraverse(node.RightChild())
}

// InorderTraverseVisit does a traverse of a binary tree,
// callling a function in-order at every node
func InorderTraverseVisit(node Node, fn VisitorFunc) {
	for !node.IsNil() {
		InorderTraverseVisit(node.LeftChild(), fn)
		fn(node)
		node = node.RightChild()
	}
}

// AllorderTraverseVisit does a traverse of a binary tree,
// potentially callling a function pre-order, in-order and/or
// post-order at every node
func AllorderTraverseVisit(node Node, preorderfn, inorderfn, postorderfn VisitorFunc) {
	if node.IsNil() {
		return
	}
	if preorderfn != nil {
		preorderfn(node)
	}
	AllorderTraverseVisit(node.LeftChild(), preorderfn, inorderfn, postorderfn)
	if inorderfn != nil {
		inorderfn(node)
	}
	AllorderTraverseVisit(node.RightChild(), preorderfn, inorderfn, postorderfn)
	if postorderfn != nil {
		postorderfn(node)
	}
}

// PreorderTraverseVisit does a traverse of a binary tree,
// callling a function pre-order at every node
func PreorderTraverseVisit(node Node, fn VisitorFunc) {
	if node == nil {
		return
	}
	fn(node)
	PreorderTraverseVisit(node.LeftChild(), fn)
	PreorderTraverseVisit(node.RightChild(), fn)
}

// InorderTraverse prints nodes' values on stdout in order.
func InorderTraverse(node Node) {
	AllorderTraverseVisit(node, NullVisitor, printNode, NullVisitor)
}

// InorderTraverseTail prints nodes' values on stdout in order,
// but only recurses on left child nodes - tail call optimization.
func InorderTraverseTail(node Node) {
	for !node.IsNil() {
		InorderTraverseTail(node.LeftChild())
		fmt.Printf("%s ", node)
		node = node.RightChild()
	}
}

// printNode relies on node fitting interface fmt.Stringer, too.
func printNode(node Node) {
	fmt.Printf(" %s", node)
}
