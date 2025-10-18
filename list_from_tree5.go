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

	flatten(root)

	for n := root; n != nil; n = n.Right {
		fmt.Printf("%d -> ", n.Data)
	}
	fmt.Println()
}

func flatten(curr *tree.NumericNode) {
	for curr != nil {
		if curr.Left != nil {
			prev := curr.Left
			for prev.Right != nil {
				prev = prev.Right
			}
			prev.Right = curr.Right
			curr.Right = curr.Left
			curr.Left = nil
		}
		curr = curr.Right
	}
}
