package main

/*
 * A unival tree (which stands for "universal value") is a tree where all nodes
 * under it have the same value.
 *
 * Given the root to a binary tree, count the number of unival subtrees.
 */

import (
	"binary_tree/tree"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Count the number of univalue subtrees in a tree\n")
		fmt.Printf("Usage: %s ''(0 (1) (0 (1 (1) (1)) (0)))'\n", os.Args[0])
		return
	}
	root := tree.CreateNumericFromString(os.Args[1])
	cnt := countUnival(root)
	fmt.Printf("%d unival subtrees\n", cnt)
}

func countUnival(node *tree.NumericNode) int {
	if node == nil {
		return 0
	}
	subtreecount := countUnival(node.Left)
	subtreecount += countUnival(node.Right)
	if (node.Left == nil || node.Data == node.Left.Data) &&
		(node.Right == nil || node.Data == node.Right.Data) {
		subtreecount++
	}

	return subtreecount
}
