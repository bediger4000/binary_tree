package tree

import (
	"fmt"
	"io"
	"os"
)

type drawable interface {
	leftChild() drawable
	rightChild() drawable
	isNil() bool
}

func (n *Node) leftChild() drawable {
	return n.left
}
func (n *Node) rightChild() drawable {
	return n.right
}
func (n *Node) isNil() bool {
	return n == nil
}

func (n *StringNode) leftChild() drawable {
	return n.Left
}
func (n *StringNode) rightChild() drawable {
	return n.Right
}
func (n *StringNode) isNil() bool {
	return n == nil
}

// Draw outputs GraphViz declarations for a single binary tree
// on standard out.
func Draw(root drawable) {
	fmt.Fprintf(os.Stdout, "digraph g {\n")
	DrawPrefixed(os.Stdout, root, "N")
	fmt.Fprintf(os.Stdout, "\n}\n")
}

// DrawPrefixed is a fairly general function to write GraphViz
// notation for a binary tree to some output. Some applications
// need a prefix to distinguish 2 disconnected trees in the same
// digraph that might have the same values.
// Leaf nodes get point-shaped left and right child nodes to keep
// the rendered shapes of trees looking right. Nodes without one or the
// child nodes also get a point-shaped pseudo-child node for the
// same reason.
func DrawPrefixed(out io.Writer, node drawable, prefix string) {
	if node.isNil() {
		return
	}
	fmt.Fprintf(out, "%s%p [label=\"%s\"];\n", prefix, node, node)
	left := node.leftChild()
	if left.isNil() {
		fmt.Fprintf(out, "%s%pL [shape=\"point\"];\n", prefix, node)
		fmt.Fprintf(out, "%s%p -> %s%pL;\n", prefix, node, prefix, node)
	} else {
		DrawPrefixed(out, left, prefix)
		fmt.Fprintf(out, "%s%p -> %s%p;\n", prefix, node, prefix, left)
	}
	right := node.rightChild()
	if right.isNil() {
		fmt.Fprintf(out, "%s%pR [shape=\"point\"];\n", prefix, node)
		fmt.Fprintf(out, "%s%p -> %s%pR;\n", prefix, node, prefix, node)
	} else {
		DrawPrefixed(out, right, prefix)
		fmt.Fprintf(out, "%s%p -> %s%p;\n", prefix, node, prefix, right)
	}
}
