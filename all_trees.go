package main

/*
 * Find all binary tree configurations/shapes/topologies of N nodes
 * Algorithm from:
 * https://stackoverflow.com/questions/33397751/generating-all-binary-trees-of-n-nodes-in-lexigraphical-order
 *
 * Given an integer N, construct all possible binary trees with N nodes.
 * For N   0   1   2   3   4   5    6    7     8     9     10 ...
 *         1   1   2   5  14  42  132  429  1430  4862  16796
 * Catalan numbers, https://oeis.org/A000108
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
	if n == 0 {
		fmt.Printf("()\n")
		return
	}

	treeList := helper(1, n)
	for _, root := range treeList {
		tree.Print(root)
		fmt.Println()
	}
	fmt.Printf("Got back %d trees\n", len(treeList))
	fmt.Printf("Should have found %d trees\n", catalanNumber(n))
}

func helper(start, end int) []*tree.StringNode {
	if start > end {
		return []*tree.StringNode{nil}
	}

	var result []*tree.StringNode

	for i := start; i <= end; i++ {
		leftList := helper(start, i-1)
		rightList := helper(i+1, end)

		for _, l := range leftList {
			for _, r := range rightList {
				root := new(tree.StringNode)
				root.Data, root.Left, root.Right = "X", l, r
				result = append(result, root)
			}
		}
	}
	return result
}

func catalanNumber(n int) int {
	if n == 0 {
		return 1
	}
	n--
	return 2 * (2*n + 1) * catalanNumber(n) / (n + 2)
}
