package tree

import "math"

// BstProperty returns true if tree has "Binary seach tree"
// property, false otherwise. Uses math.MinInt32 as minimum
// value for LHS of tree to be greater than, and math.MaxInt32
// as the value for RHS of tree to have values less than.
func BstProperty(root *Node) bool {
	if !bst(root.left, math.MinInt32, root.data) {
		return false
	}
	if !bst(root.right, root.data, math.MaxInt32) {
		return false
	}
	return true
}

// bst is the function that actually checks BST property for
// a given node somewhere in the tree. Needs to be called
// with max and min values so that it can do comparisons,
// so it's not actually suitable for users to call it.
func bst(node *Node, min int, max int) bool {

	if node == nil {
		return true
	}

	if !(node.data > min && node.data < max) {
		return false
	}
	if !bst(node.left, min, node.data) {
		return false
	}
	if !bst(node.right, node.data, max) {
		return false
	}

	return true
}
