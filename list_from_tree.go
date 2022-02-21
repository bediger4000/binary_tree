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

	root := tree.CreateNumericFromString(os.Args[1])

	ch := make(chan *tree.NumericNode, 0)

	go traverse(root, ch)

	var head, tail *tree.NumericNode

	for node := range ch {
		if head == nil {
			head = node
			tail = node
			tail.Right = nil
			continue
		}
		tail.Right = node
		tail = node
		tail.Right = nil
	}

	for node := head; node != nil; node = node.Right {
		fmt.Printf("%d -> ", node.Data)
	}
	fmt.Println()
}

func traverse(root *tree.NumericNode, ch chan *tree.NumericNode) {

	realTraverse(root, ch)
	close(ch)
}

func realTraverse(node *tree.NumericNode, ch chan *tree.NumericNode) {
	if node == nil {
		return
	}
	// This program won't work if it doesn't save node.Left and node.Right here.
	// The main goroutine changes node.Left and node.Right after reading node
	// from the channel.
	left := node.Left
	right := node.Right
	ch <- node
	realTraverse(left, ch)
	realTraverse(right, ch)
}
