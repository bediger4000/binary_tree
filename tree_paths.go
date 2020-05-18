package main

/*
Given a binary tree, return all paths from the root to leaves.

For example, given the tree

   1
  / \
 2   3
    / \
   4   5

it should return [[1, 2], [1, 3, 4], [1, 3, 5]].
*/

import (
	"fmt"
	"os"
	"strconv"
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

func inorderTraverse(node *TreeNode, values *[]int) {
	if node == nil {
		return
	}
	inorderTraverse(node.left, values)
	*values = append(*values, node.data)
	inorderTraverse(node.right, values)
}

// 3 9 6 5 7
func pathTraverse(node *TreeNode, path []int, paths *[][]int) {
	if node == nil {
		return
	}

	// path is a slice: sometimes append()
	// doesn't actually reallocate the backing
	// store of the slice, so you get duplicate
	// paths, and miss one path. Make a copy.
	mypath := make([]int, len(path))
	copy(mypath, path)

	mypath = append(mypath, node.data)

	if node.left == nil && node.right == nil {
		*paths = append(*paths, mypath)
		return
	}

	pathTraverse(node.left, mypath, paths)
	pathTraverse(node.right, mypath, paths)
}

func main() {
	var root *TreeNode

	for _, str := range os.Args[1:] {
		val, err := strconv.Atoi(str)

		if err == nil {
			// fmt.Printf("insert %d\n", val)
			root = insert(root, val)
		} else {
			fmt.Printf("Problem with %q: %s\n", str, err)
		}
	}

	if root != nil {
		/*
			fmt.Printf("digraph g {\n")
			drawTree(root)
			fmt.Printf("}\n")
		*/
		var values []int
		inorderTraverse(root, &values)
		fmt.Printf("/* All values: %v */\n", values)
		var path []int
		var paths [][]int
		pathTraverse(root, path, &paths)
		fmt.Printf("/*\n")
		for _, p := range paths {
			fmt.Printf("%v\n", p)
		}
		fmt.Printf("*/\n")
	}
}
