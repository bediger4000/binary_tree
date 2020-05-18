package tree

// Insert puts a node into a (possibly empty) binary tree.
// It preserves the binary-search-property, every node to the
// left of the current node has data value < current node's data value
//
// Usage strikes me as a little odd, but makes for a lot fewer lines
// of code overall:
// var root *Node
// for _, value := range values {
//    root = Insert(root, value)
// }
func Insert(node *Node, value int) *Node {

	if node == nil {
		return &Node{Data: value}
	}

	n := &(node.Left)
	if value >= node.Data {
		n = &(node.Right)
	}
	*n = Insert(*n, value)
	return node
}
