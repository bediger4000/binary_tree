package tree

// Equals returns true if two tree nodes have the same value,
// and each of their subtrees' nodes have the same values.
// Trees that don't have the same structure (missing node(s),
// or addtional node(s)) cause it to return false.
func Equals(t1, t2 Node) bool {
	// The both-nil and one-not-nil tests will answer
	// incorrectly if t1 and t2 are not of the same type.
	if t1.IsNil() {
		return t2.IsNil()
	}
	if t2.IsNil() {
		return false
	}

	switch ty1 := t1.(type) {
	case *NumericNode:
		switch ty2 := t2.(type) {
		case *NumericNode:
			if ty1.Data != ty2.Data {
				return false
			}
		default:
			return false
		}
	case *StringNode:
		switch ty2 := t2.(type) {
		case *StringNode:
			if ty1.Data != ty2.Data {
				return false
			}
		default:
			return false
		}
	}
	if !Equals(t1.LeftChild(), t2.LeftChild()) {
		return false
	}
	return Equals(t1.RightChild(), t2.RightChild())
}
