package tree

// Equals returns true if two tree nodes have the same value,
// and each of their subtrees' nodes have the same values.
// Trees that don't have the same structure (missing node(s),
// or addtional node(s)) cause it to return false.
func Equals(t1, t2 *NumericNode) bool {
	if t1 == nil {
		return t2 == nil
	}
	if t2 == nil {
		return false
	}
	if t1.Data != t2.Data {
		return false
	}
	if !Equals(t1.Left, t2.Left) {
		return false
	}
	return Equals(t1.Right, t2.Right)
}
