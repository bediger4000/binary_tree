package tree

import "fmt"

// Node is an element of a binary tree with a numeric value
type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

// StringNode is an element of a binary tree with a string value
type StringNode struct {
	Data  string
	Left  *StringNode
	Right *StringNode
}

// Value returns whatever the node has as its data,
// but just the value of that data, mind you.
func (n *Node) Value() int {
	return n.Data
}

// Value returns the string that's the data a StringNode
// instance carries around. Since string is a value type,
// I think this is OK - we're not giving out data that can
// be changed out from under us. It's a getter, though.
func (n *StringNode) Value(s string) {
	n.Data = s
}

func (n *Node) String() string {
	return fmt.Sprintf("%d", n.Data)
}

func (n *StringNode) String() string {
	return n.Data
}

// VisitorFunc types the *TraverseVisit functions in tree/traverse.go
type VisitorFunc func(node *Node)

// NullVisitor does nothing, but it can be used to make
// tree.AllorderTraverseVisit into pre-, post-, in-order
// or a combination. A placeholder.
func NullVisitor(node *Node) {
}
