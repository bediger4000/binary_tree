package tree

func ProbeDepth(node Node, turns []int) int {
	if node.isNil() {
		return 0
	}

	child := node.leftChild()
	if child.isNil() {
		return 1
	}
	if turns[0] == 1 {
		child = node.rightChild()
	}
	if child.isNil() {
		return 1
	}

	return ProbeDepth(child, turns[1:]) + 1
}
