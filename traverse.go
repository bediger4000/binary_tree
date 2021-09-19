package main

/*
 * See if generalized AllorderTraverseVisit can do in-order traversal.
 * It can.
 */

import (
	"binary_tree/tree"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var root *tree.NumericNode
	for i := range os.Args[1:] {
		n, err := strconv.Atoi(os.Args[i+1])
		if err == nil {
			root = tree.Insert(root, n)
		}
	}
	inordertraverse(root)
	fmt.Println()
}

// inordertraverse sets up tree.AllorderTraverseVisit for in-order
// traverse of a tree.
func inordertraverse(root *tree.NumericNode) {
	tree.AllorderTraverseVisit(root, nil, printnode, nil)
}

func printnode(node tree.Node) {
	if node.IsNil() {
		return
	}
	fmt.Printf("%s, ", node)
}
