package main

/*
 * Given a sorted array increasing order of unique integers,
 * create a binary search tree with minimal height.
 */

import (
	"binary_tree/tree"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

func minHeightInsert(sortedArray []int) (root *tree.NumericNode) {
	root = nil
	if sz := len(sortedArray); sz > 0 {
		middle := sz / 2
		root = &tree.NumericNode{Data: sortedArray[middle]}
		root.Left = minHeightInsert(sortedArray[0:middle])
		root.Right = minHeightInsert(sortedArray[middle+1:])
	}
	return root
}

func main() {

	sortedArray := make([]int, 0)

	N := 1
	outputGraphViz := false
	if os.Args[1] == "-g" {
		N = 2
		outputGraphViz = true
	}

	for _, str := range os.Args[N:] {
		val, err := strconv.Atoi(str)
		if err == nil {
			sortedArray = append(sortedArray, val)
		}
	}

	sort.Sort(sort.IntSlice(sortedArray))

	fmt.Printf("/* Sorted array: %v */\n", sortedArray)

	root := minHeightInsert(sortedArray)

	if root != nil {
		depth, _ := tree.FindDepth2(root, 1)
		D := math.Log2(float64(len(sortedArray)) + 1.0)
		f := float64(depth)
		if f >= D && f <= D+1.0 {
			fmt.Printf("/* minimal height tree %.3f <= %.3f <= %.3f */\n", D, f, D+1.0)
		}

		fmt.Printf("/* In order traverse: ")
		tree.InorderTraverse(root)
		fmt.Printf(" */\n")
		if outputGraphViz {
			tree.Draw(root)
		}
		if tree.BstProperty(root) {
			fmt.Printf("/* This is a binary search tree */\n")
		}
	}
}
