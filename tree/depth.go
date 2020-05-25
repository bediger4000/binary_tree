package tree

// Depth used to trace which tree's node occurs
// at the maximum depth discovered so far.
type Depth struct {
	Node  *NumericNode
	Depth int
}

// FindDepth1 traverses a binary tree, saving the maximum depth
// it finds in a struct that's passed down the levels of recursion
// used to traverse the tree.
func FindDepth1(node *NumericNode, ply int, d *Depth) {
	if node == nil {
		return
	}

	FindDepth1(node.Left, ply+1, d)
	if ply > d.Depth {
		d.Depth = ply
		d.Node = node
	}
	FindDepth1(node.Right, ply+1, d)
}

// FindDepth2 traverses a binary tree, returning a pointer to the
// deepest node of the tree it finds.
// The flaw in this one is that when you call it on the tree's root,
// you have to pass a zero value for ply, otherwise the depth is offset.
func FindDepth2(node *NumericNode, ply int) (depth int, deepnode *NumericNode) {
	// past leaf node
	if node == nil {
		return -1, nil
	}

	// leaf node
	if node.Left == nil && node.Right == nil {
		return ply, node
	}

	// interior node
	ldepth, lnode := FindDepth2(node.Left, ply+1)
	rdepth, rnode := FindDepth2(node.Right, ply+1)

	if ldepth > rdepth {
		return ldepth, lnode
	}

	return rdepth, rnode
}
