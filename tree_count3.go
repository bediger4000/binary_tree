package main

/*
 * Given a complete binary tree, count the number of nodes in faster than O(n)
 * time. Recall that a complete binary tree has every level filled except the
 * last, and the nodes in the last level are filled starting from the left.
 * "Complete" means: every level, except possibly the last, is completely
 * filled, and all nodes in the last level are as far left as possible. It
 * can have between 1 and 2h nodes at the last level h.
 *
 * Binary sub-tree search version
 */

import (
	"binary_tree/tree"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	root := generateHeap(n)
	nodeCount := countNodes(root)

	n, nodesTouched := binarySubtreeCount(root, -1, -1)

	fmt.Printf("%d nodes in tree, found %d, %d nodes accessed\n", nodeCount, n, nodesTouched)
}

// binarySubtreeCount returns count of nodes in subtree rooted at formal
// argument named node, and the number of nodes accesses to get that count.
// lDepth, rDepth are either -1 (calculate left and/or right depth) or the
// left and/or right depth for formal argment node.
func binarySubtreeCount(node *tree.NumericNode, lDepth, rDepth int) (int, int) {
	if node == nil {
		return 0, 0
	}

	touched := 0
	if lDepth < 0 {
		// find left depth from left child to avoid re-counting current node
		lDepth = tree.LeftDepth(node.Left)
		touched += lDepth
		lDepth++
	}
	if rDepth < 0 {
		// find right depth from right child to avoid re-counting current node
		rDepth = tree.RightDepth(node.Right)
		touched += rDepth
		rDepth++
	}

	if lDepth == rDepth {
		// full binary tree rooted at node, 2^D-1 nodes in size
		return (1 << lDepth) - 1, touched
	}

	// lDepth > rDepth
	nLeft, touchedLeft := binarySubtreeCount(node.Left, lDepth-1, -1)
	nRight, touchedRight := binarySubtreeCount(node.Right, -1, rDepth-1)

	return nLeft + nRight + 1, touched + touchedLeft + touchedRight
}

func countNodes(node *tree.NumericNode) int {
	if node == nil {
		return 0
	}
	l := countNodes(node.Left)
	r := countNodes(node.Right)
	return r + l + 1
}

func generateHeap(nodeCount int) *tree.NumericNode {

	root := &tree.NumericNode{Data: nodeCount}
	nodeCount--

	q := &tree.Stack{}
	q.Push(root)

	for q != nil && nodeCount > 0 {
		node := q.Dequeue().(*tree.NumericNode)

		node.Left = &tree.NumericNode{Data: nodeCount}
		nodeCount--
		if nodeCount == 0 {
			break
		}
		q.Push(node.Left)

		node.Right = &tree.NumericNode{Data: nodeCount}
		nodeCount--
		if nodeCount == 0 {
			break
		}
		q.Push(node.Right)
	}

	return root
}
