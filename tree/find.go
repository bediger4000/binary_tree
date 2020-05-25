package tree

// Find performs binary search on argument node
// and its sub-trees.
func Find(node *NumericNode, value int) *NumericNode {
	if node == nil || node.Data == value {
		return node
	}
	if node.Data > value {
		return Find(node.Left, value)
	}
	return Find(node.Right, value)
}
