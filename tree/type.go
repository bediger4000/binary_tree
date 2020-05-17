package tree

import "fmt"

// Node is an element of a binary tree with a numeric value
type Node struct {
	data  int
	left  *Node
	right *Node
}

// StringNode is an element of a binary tree with a string value
type StringNode struct {
	Data  string
	Left  *Node
	Right *Node
}

// Value returns whatever the node has as its data,
// but just the value of that data, mind you.
func (n *Node) Value() int {
	return n.data
}

func (n *Node) String() string {
	return fmt.Sprintf("%d", n.data)
}

func (n *StringNode) String() string {
	return n.Data
}
