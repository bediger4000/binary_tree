package tree

// Insert puts a node into a (possibly empty) binary tree.
// It preserves the binary-search-property, every node to the
// left of the current node has data value < current node's data value
//
// Usage strikes me as a little odd, but makes for a lot fewer lines
// of code overall:
// var root *NumericNode
// for _, value := range values {
//    root = Insert(root, value)
// }
func Insert(node *NumericNode, value int) *NumericNode {

	if node == nil {
		return &NumericNode{Data: value}
	}

	// Skip duplicates - is this correct?
	if node.Data == value {
		return node
	}

	n := &(node.Left)
	if value >= node.Data {
		n = &(node.Right)
	}
	*n = Insert(*n, value)
	return node
}
