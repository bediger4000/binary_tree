package tree

import (
	"fmt"
	"io"
	"os"
)

// Draw outputs GraphViz declarations for a single binary tree
// on standard out.
func Draw(root *Node) {
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
func DrawPrefixed(out io.Writer, node *Node, prefix string) {
	fmt.Fprintf(out, "%s%p [label=\"%d\"];\n", prefix, node, node.data)
	if node.left != nil {
		DrawPrefixed(out, node.left, prefix)
		fmt.Fprintf(out, "%s%p -> %s%p;\n", prefix, node, prefix, node.left)
	} else {
		fmt.Fprintf(out, "%s%pL [shape=\"point\"];\n", prefix, node)
		fmt.Fprintf(out, "%s%p -> %s%pL;\n", prefix, node, prefix, node)
	}
	if node.right != nil {
		DrawPrefixed(out, node.right, prefix)
		fmt.Fprintf(out, "%s%p -> %s%p;\n", prefix, node, prefix, node.right)
	} else {
		fmt.Fprintf(out, "%s%pR [shape=\"point\"];\n", prefix, node)
		fmt.Fprintf(out, "%s%p -> %s%pR;\n", prefix, node, prefix, node)
	}
}
