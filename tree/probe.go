package tree

func ProbeDepth(node Node, turns []int) int {
	if node.IsNil() {
		return 0
	}

	child := node.LeftChild()
	if child.IsNil() {
		return 1
	}
	if turns[0] == 1 {
		child = node.RightChild()
	}
	if child.IsNil() {
		return 1
	}

	return ProbeDepth(child, turns[1:]) + 1
}
