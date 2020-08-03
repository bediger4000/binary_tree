package tree

import (
	"fmt"
)

// NumericNode is an element of a binary tree with a numeric value
type NumericNode struct {
	Data  int
	Left  *NumericNode
	Right *NumericNode
}

// StringNode is an element of a binary tree with a string value
type StringNode struct {
	Data  string
	Left  *StringNode
	Right *StringNode
}

// Value returns whatever the node has as its data,
// but just the value of that data, mind you.
func (n *NumericNode) Value() int {
	return n.Data
}

/*
// Value returns the string that's the data a StringNode
// instance carries around. Since string is a value type,
// I think this is OK - we're not giving out data that can
// be changed out from under us. It's a getter, though.
func (n *StringNode) Value(s string) {
	n.Data = s
}
*/

func (n *NumericNode) String() string {
	return fmt.Sprintf("%d", n.Data)
}

func (n *StringNode) String() string {
	return n.Data
}

// VisitorFunc types the *TraverseVisit functions in tree/traverse.go
type VisitorFunc func(node Node)

// NodeCreatorFn instances turn a string into a single Node,
// but the Node is actually of some user-defined type that fits
// interface Node
type NodeCreatorFn func(stringrep string) Node

// NullVisitor does nothing, but it can be used to make
// tree.AllorderTraverseVisit into pre-, post-, in-order
// or a combination. A placeholder.
func NullVisitor(node Node) {
}

// Node makes it possible to use the same code for StringNode
// and NumericNode pointers when printing or traversing a tree.
type Node interface {
	LeftChild() Node
	RightChild() Node
	SetLeftChild(Node)
	SetRightChild(Node)
	IsNil() bool
}

func (n *NumericNode) LeftChild() Node {
	return n.Left
}
func (n *NumericNode) RightChild() Node {
	return n.Right
}
func (n *NumericNode) SetLeftChild(node Node) {
	if node == nil {
		return
	}
	n.Left = node.(*NumericNode)
}
func (n *NumericNode) SetRightChild(node Node) {
	if node == nil {
		return
	}
	n.Right = node.(*NumericNode)
}
func (n *NumericNode) IsNil() bool {
	return n == nil
}

func (n *StringNode) LeftChild() Node {
	return n.Left
}
func (n *StringNode) RightChild() Node {
	return n.Right
}
func (n *StringNode) SetLeftChild(node Node) {
	if node == nil {
		return
	}
	n.Left = node.(*StringNode)
}
func (n *StringNode) SetRightChild(node Node) {
	if node == nil {
		return
	}
	n.Right = node.(*StringNode)
}
func (n *StringNode) IsNil() bool {
	return n == nil
}
