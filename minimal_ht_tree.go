package main

/*
 * Given a sorted array increasing order of unique integers,
 * create a binary search tree with minimal height.
 */

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"time"
)

type TreeNode struct {
	data  int
	left  *TreeNode
	right *TreeNode
}

func insert(node *TreeNode, value int) *TreeNode {

	if node == nil {
		return &TreeNode{data: value}
	}

	n := &(node.left)
	if value >= node.data {
		n = &(node.right)
	}
	*n = insert(*n, value)
	return node
}

func minHeightInsert(sortedArray []int) (root *TreeNode) {
	sz := len(sortedArray)
	switch sz {
	case 1:
		root = &TreeNode{data: sortedArray[0]}
	case 2:
		// Arrays of size two can end up in 2 arrangements:
		if rand.Intn(2) == 0 {
			root = &TreeNode{data: sortedArray[1]}
			root.left = &TreeNode{data: sortedArray[0]}
		} else {
			root = &TreeNode{data: sortedArray[0]}
			root.right = &TreeNode{data: sortedArray[1]}
		}
	case 3:
		root = &TreeNode{data: sortedArray[1]}
		root.left = &TreeNode{data: sortedArray[0]}
		root.right = &TreeNode{data: sortedArray[2]}
	default:
		middle := sz / 2
		// You've got a choice of "middle" for an even
		// array size.
		if (middle % 2) == 0 {
			middle -= rand.Intn(2)
		}
		root = &TreeNode{data: sortedArray[middle]}
		root.left = minHeightInsert(sortedArray[0:middle])
		root.right = minHeightInsert(sortedArray[middle+1:])

	}
	return root
}

func inorderTraverse(node *TreeNode) {
	if node == nil {
		return
	}
	inorderTraverse(node.left)
	fmt.Printf("%d ", node.data)
	inorderTraverse(node.right)
}

func drawTree(node *TreeNode) {
	fmt.Printf("Node%p [label=\"%d\"];\n", node, node.data)
	if node.left != nil {
		drawTree(node.left)
		fmt.Printf("Node%p -> Node%p;\n", node, node.left)
	} else {
		fmt.Printf("Node%pL [shape=\"point\"];\n", node)
		fmt.Printf("Node%p -> Node%pL;\n", node, node)
	}
	if node.right != nil {
		drawTree(node.right)
		fmt.Printf("Node%p -> Node%p;\n", node, node.right)
	} else {
		fmt.Printf("Node%pR [shape=\"point\"];\n", node)
		fmt.Printf("Node%p -> Node%pR;\n", node, node)
	}
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
		inorderTraverse(root)
		fmt.Printf(" */\n")
		fmt.Printf("digraph g {\n")
		drawTree(root)
		fmt.Printf("\n}\n")
		if bstProperty(root) {
			fmt.Printf("/* This is a binary search tree */\n")
		}
	}
}

// Return true if tree has "Binary seach tree"
// property, false otherwise.
func bstProperty(root *TreeNode) bool {
	if !bst(root.left, math.MinInt32, root.data) {
		return false
	}
	if !bst(root.right, root.data, math.MaxInt32) {
		return false
	}
	return true
}

// function that actually checks BST property for
// a given node somewhere in the tree.
func bst(node *TreeNode, min int, max int) bool {

	if node == nil {
		return true
	}

	if !(node.data > min && node.data < max) {
		return false
	}
	if !bst(node.left, min, node.data) {
		return false
	}
	if !bst(node.right, node.data, max) {
		return false
	}

	return true
}
