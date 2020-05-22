package main

/*
 * Given a sorted array increasing order of unique integers,
 * create a binary search tree with minimal height.
 */

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"time"

	"binary_tree/tree"
)

func minHeightInsert(sortedArray []int) (root *tree.Node) {
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
		root.Left = minHeightInsert(sortedArray[0:middle])
		root.Right = minHeightInsert(sortedArray[middle+1:])

	}
	return root
}

func main() {

	rand.Seed(time.Now().UnixNano())

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
		fmt.Printf("/* In order traverse: ")
		tree.InorderTraverse(root)
		fmt.Printf(" */\n")
		tree.Draw(root)
		if tree.BstProperty(root) {
			fmt.Printf("/* This is a binary search tree */\n")
		}
	}
}
