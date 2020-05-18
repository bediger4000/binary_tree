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

func (n *StringNode) Value(s string) {
	n.Data = s
}

func (n *Node) String() string {
	return fmt.Sprintf("%d", n.Data)
}

func (n *StringNode) String() string {
	return n.Data
}
