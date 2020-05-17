package tree

// Node - binary tree element
type Node struct {
	data  int
	left  *Node
	right *Node
}

// Value returns whatever the node has as its data,
// but just the value of that data, mind you.
func (n *Node) Value() int {
	return n.data
}
