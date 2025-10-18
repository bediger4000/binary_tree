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

type nodeHolder struct {
	left  *tree.NumericNode
	right *tree.NumericNode
	node  *tree.NumericNode
	next  *nodeHolder
}

func main() {

	root, err := tree.CreateNumericFromString(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "parsing %q: %v\n", os.Args[1], err)
		return
	}

	stack := &nodeHolder{node: root, left: root.Left, right: root.Right}

	var head, tail *tree.NumericNode

	for stack != nil {
		n := stack
		stack = stack.next
		fmt.Printf("%d -> ", n.node.Data)
		if n.right != nil {
			nn := &nodeHolder{node: n.right, left: n.right.Left, right: n.right.Right}
			nn.next = stack
			stack = nn
		}
		if n.left != nil {
			nn := &nodeHolder{node: n.left, left: n.left.Left, right: n.left.Right}
			nn.next = stack
			stack = nn
		}
		if head == nil {
			head = n.node
			tail = n.node
			continue
		}
		tail.Right = n.node
		tail = tail.Right
		tail.Left = nil
	}
	fmt.Println()

	for n := head; n != nil; n = n.Right {
		fmt.Printf("%d -> ", n.Data)
	}
	fmt.Println()
}
