package tree

import "fmt"

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

// Value returns the string that's the data a StringNode
// instance carries around. Since string is a value type,
// I think this is OK - we're not giving out data that can
// be changed out from under us. It's a getter, though.
func (n *StringNode) Value(s string) {
	n.Data = s
}

func (n *NumericNode) String() string {
	return fmt.Sprintf("%d", n.Data)
}

func (n *StringNode) String() string {
	return n.Data
}

// VisitorFunc types the *TraverseVisit functions in tree/traverse.go
type VisitorFunc func(node Node)

// NullVisitor does nothing, but it can be used to make
// tree.AllorderTraverseVisit into pre-, post-, in-order
// or a combination. A placeholder.
func NullVisitor(node Node) {
}

type Node interface {
	leftChild() Node
	rightChild() Node
	isNil() bool
}

func (n *NumericNode) leftChild() Node {
	return n.Left
}
func (n *NumericNode) rightChild() Node {
	return n.Right
}
func (n *NumericNode) isNil() bool {
	return n == nil
}

func (n *StringNode) leftChild() Node {
	return n.Left
}
func (n *StringNode) rightChild() Node {
	return n.Right
}
func (n *StringNode) isNil() bool {
	return n == nil
}
