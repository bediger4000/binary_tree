package main

/*
This problem was asked by Apple.

Given the root of a binary tree,
find the most frequent subtree sum.
The subtree sum of a node is the sum of all values under a node,
including the node itself.

For example, given the following tree:

      5
     / \
    2  -5

Return 2 as it occurs twice: once as the left leaf,
and once as the sum of 2 + 5 - 5.

*/

import (
	"binary_tree/tree"
	"fmt"
	"os"
)

func main() {
	root := tree.CreateNumericFromString(os.Args[1])

	ch := make(chan int)
	sumsCount := make(map[int]int)

	go findSums(root, ch)

	// read subtree sums from channel, count them.
	for sum := range ch {
		sumsCount[sum]++
	}

	for sum, count := range sumsCount {
		fmt.Printf("sum %d appears %d times\n", sum, count)
	}
}

func findSums(root *tree.NumericNode, ch chan int) {
	// it's just harder to screw up if you have a separate
	// recursive function to do things, and then close the
	// channel when the tree traverse finishes.
	// The recursive function really can't tell when it's done,
	// so it's hard for it to decide when to close the channel.
	recursiveFindSum(root, ch)
	close(ch)
}

func recursiveFindSum(node *tree.NumericNode, ch chan int) int {
	// single check for having recursed far enough
	if node == nil {
		return 0
	}

	leftSum := recursiveFindSum(node.Left, ch)
	rightSum := recursiveFindSum(node.Right, ch)
	sum := node.Data + leftSum + rightSum
	ch <- sum  // to count subtree sums
	return sum // for caller to use
}
