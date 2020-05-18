package main

import (
	"binary_tree/tree"
	"fmt"
	"os"
)

/* Daily Coding Problem: Problem #540 [Easy]

In Ancient Greece, it was common to write text with the first line going left
to right, the second line going right to left, and continuing to go back and
forth. This style was called "boustrophedon".

Given a binary tree, write an algorithm to print the nodes in boustrophedon order.

For example, given the following tree:

       1
     /   \
    /     \
  2         3
 / \       / \
4   5     6   7

You should return [1, 3, 2, 4, 5, 6, 7].
*/

type NodeQueue struct {
	array []*tree.Node
}

func main() {
	root := tree.CreateNumeric(os.Args[1:])

	queue := new(NodeQueue)

	queue.push(root)

	for !queue.empty() {
		n := queue.dequeue()
		fmt.Printf("%d ", n.Data)
		if n.Left != nil {
			queue.push(n.Left)
		}
		if n.Right != nil {
			queue.push(n.Right)
		}
	}
	fmt.Println()
}

func (nq *NodeQueue) push(n *tree.Node) {
	nq.array = append(nq.array, n)
}

func (nq *NodeQueue) dequeue() (head *tree.Node) {
	if len(nq.array) == 0 {
		return
	}
	head = nq.array[0]
	nq.array = nq.array[1:]
	return
}

func (nq *NodeQueue) empty() bool {
	if len(nq.array) == 0 {
		return true
	}
	return false
}
