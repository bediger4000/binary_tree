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

func main() {

	sortedArray := make([]int, 0)

	for _, str := range os.Args[1:] {
		val, err := strconv.Atoi(str)
		if err == nil {
			sortedArray = append(sortedArray, val)
		}
	}

	sort.Sort(sort.IntSlice(sortedArray))

	fmt.Printf("/* Sorted array: %v */\n", sortedArray)

	root := minHeightInsert(sortedArray)

	if root != nil {
		depth, _ := tree.FindDepth2(root, 0)
		D := math.Log2(float64(len(sortedArray)) + 1.0)
		f := float64(depth)
		if f >= D && f <= D+1.0 {
			fmt.Println("/* minimal height tree */")
		}

		fmt.Printf("/* In order traverse: ")
		tree.InorderTraverse(root)
		fmt.Printf(" */\n")
		tree.Draw(root)
		if tree.BstProperty(root) {
			fmt.Printf("/* This is a binary search tree */\n")
		}
	}
}
