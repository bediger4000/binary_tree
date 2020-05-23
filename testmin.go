package main

/*
 * Given a sorted array increasing order of unique integers,
 * create a binary search tree with minimal height.
 *
 * Now test it.
 * Number of nodes in a complete binary tree of depth D:
 * N = 2^D - 1
 *
 * So D = log2(N + 1)
 *
 * Actual height (depth?) of a tree is a step function, not continuous.
 * Depth of a binary tree will be 1 more than D, for N not the number of
 * nodes in a complete tree.
 *
 * If depth of tree <= log2(N+1) + 1, where N is the number of nodes
 * in the tree, complete or not, it's a minimal height tree
 */

import (
	"binary_tree/tree"
	"fmt"
	"math"
	"math/rand"
	"os"
	"time"
)

// minHeightInsert take "middle" element of array, take it as root value,
// the sub-array with values less than middle element as array for left
// subtree, sub-array with values greater than middle element
// as array for right sub-tree. Deterministic.
func minHeightInsert(sortedArray []int) (root *tree.Node) {
	root = nil
	if sz := len(sortedArray); sz > 0 {
		middle := sz / 2
		root = &tree.Node{Data: sortedArray[middle]}
		root.Left = minHeightInsert(sortedArray[0:middle])
		root.Right = minHeightInsert(sortedArray[middle+1:])
	}
	return root
}

// minHeightInsert2 is a non-deterministic version of minHeightInsert
func minHeightInsert2(sortedArray []int) (root *tree.Node) {
	sz := len(sortedArray)
	switch sz {
	case 1:
		root = &tree.Node{Data: sortedArray[0]}
	case 2:
		// Arrays of size two can end up in 2 arrangements:
		if rand.Intn(2) == 0 {
			root = &tree.Node{Data: sortedArray[1]}
			root.Left = &tree.Node{Data: sortedArray[0]}
		} else {
			root = &tree.Node{Data: sortedArray[0]}
			root.Right = &tree.Node{Data: sortedArray[1]}
		}
	case 3:
		root = &tree.Node{Data: sortedArray[1]}
		root.Left = &tree.Node{Data: sortedArray[0]}
		root.Right = &tree.Node{Data: sortedArray[2]}
	default:
		middle := sz / 2
		// You've got a choice of "middle" for an even
		// array size.
		if (middle % 2) == 0 {
			middle -= rand.Intn(2)
		}
		root = &tree.Node{Data: sortedArray[middle]}
		root.Left = minHeightInsert2(sortedArray[0:middle])
		root.Right = minHeightInsert2(sortedArray[middle+1:])

	}
	return root
}

func main() {

	rand.Seed(time.Now().UnixNano() + int64(os.Getpid()))

	var sortedArray []int

	for r := 1; r < 66; r++ {
		sortedArray = append(sortedArray, r)

		// root := minHeightInsert(sortedArray)
		root := minHeightInsert2(sortedArray)

		depth, _ := tree.FindDepth2(root, 1)
		D := math.Log2(float64(len(sortedArray)) + 1.0)
		f := float64(depth)
		minht := false
		if f-D <= 1.0 {
			minht = true
		}

		bst := tree.BstProperty(root)

		fmt.Printf("%d\t%d\t%f\t%f\t%v\t%v\n",
			r, depth, D, D+1.0, bst, minht)
	}
}
