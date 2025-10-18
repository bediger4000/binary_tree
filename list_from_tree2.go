package main

import (
	"binary_tree/tree"
	"fmt"
	"os"
)

/*
LC 114 [Medium] Flatten Binary Tree to Linked List

Question: Given a binary tree, flatten it to a linked list in-place.

For example, given the following tree:

    1
   / \
  2   5
 / \   \
3   4   6

The flattened tree should look like:

1
 \
  2
   \
    3
     \
      4
       \
        5
         \
          6
*/

func main() {

	root, err := tree.CreateNumericFromString(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "parsing %q: %v\n", os.Args[1], err)
		return
	}

	head := &tree.NumericNode{}

	traverse(root, &head)

	for n := head.Right; n != nil; n = n.Right {
		fmt.Printf("%d -> ", n.Data)
	}
	fmt.Println()
}

func traverse(node *tree.NumericNode, last **tree.NumericNode) **tree.NumericNode {
	if node == nil {
		return last
	}

	(*last).Right = node
	(*last).Left = nil

	left := node.Left
	right := node.Right

	return traverse(right, traverse(left, &node))
}
