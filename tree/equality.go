package tree

// Equals returns true if two tree nodes have the same value,
// and each of their subtrees' nodes have the same values.
// Trees that don't have the same structure (missing node(s),
// or addtional node(s)) cause it to return false.
func Equals(t1, t2 *Node) bool {
	if t1 == nil {
		return t2 == nil
	}
	if t2 == nil {
		return false
	}
	if t1.data != t2.data {
		return false
	}
	if !Equals(t1.left, t2.left) {
		return false
	}
	return Equals(t1.right, t2.right)
}
