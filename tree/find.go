package tree

// Find performs binary search on argument node
// and its sub-trees.
func Find(node *Node, value int) *Node {
	if node == nil || node.Data == value {
		return node
	}
	if node.Data > value {
		return Find(node.Left, value)
	}
	return Find(node.Right, value)
}
