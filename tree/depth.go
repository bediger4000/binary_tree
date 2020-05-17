package tree

type Depth struct {
	node  *Node
	depth int
}

// FindDepth1 traverses a binary tree, saving the maximum depth
// it finds in a struct that's passed down the levels of recursion
// used to traverse the tree.
func FindDepth1(node *Node, ply int, d *Depth) {
	if node == nil {
		return
	}

	FindDepth1(node.left, ply+1, d)
	if ply > d.depth {
		d.depth = ply
		d.node = node
	}
	FindDepth1(node.right, ply+1, d)
}

// FindDepth2 traverses a binary tree, returning a pointer to the
// deepest node of the tree it finds.
func FindDepth2(node *Node, ply int) (depth int, deepnode *Node) {
	// past leaf node
	if node == nil {
		return -1, nil
	}

	// leaf node
	if node.left == nil && node.right == nil {
		return ply, node
	}

	// interior node
	ldepth, lnode := FindDepth2(node.left, ply+1)
	rdepth, rnode := FindDepth2(node.right, ply+1)

	if ldepth > rdepth {
		return ldepth, lnode
	}

	return rdepth, rnode
}
