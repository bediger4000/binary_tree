package tree

import "math"

// BstProperty returns true if tree has "Binary seach tree"
// property, false otherwise. Uses math.MinInt32 as minimum
// value for LHS of tree to be greater than, and math.MaxInt32
// as the value for RHS of tree to have values less than.
func BstProperty(root *NumericNode) bool {
	return bst(root, math.MinInt32, math.MaxInt32)
}

// bst is the function that actually checks BST property for
// a given node somewhere in the tree. Needs to be called
// with max and min values so that it can do comparisons,
// so it's not actually suitable for users to call it.
func bst(node *NumericNode, min, max int) bool {

	if node == nil {
		return true
	}

	if !(node.Data > min && node.Data < max) {
		return false
	}
	if !bst(node.Left, min, node.Data) {
		return false
	}
	if !bst(node.Right, node.Data, max) {
		return false
	}

	return true
}
