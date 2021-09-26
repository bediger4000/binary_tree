package main

/*
 * Print lisp-like serialized tree with heap property.
 * Creates binary tree of *tree.NumericNode, then prints it.
 * With -g flag, prints GraphViz representation on stdout.
 *
 * Node data like this:
 *            1
 *           / \
 *          2   3
 *         / \
 *        4   5
 *
 * Root is node with data 1, final node is node with data N.
 */

import (
	"binary_tree/tree"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	var drawGraph bool
	x := 1
	if os.Args[x] == "-g" {
		drawGraph = true
		x = 2
	}

	n, err := strconv.Atoi(os.Args[x])
	if err != nil {
		log.Fatal(err)
	}

	root := createHeap(n)

	if drawGraph {
		fmt.Print("/* ")
	}
	tree.Print(root)
	if drawGraph {
		fmt.Print(" */")
	}
	fmt.Println()

	if drawGraph {
		tree.Draw(root)
	}
}

func createHeap(n int) (root *tree.NumericNode) {
	stack := &tree.Stack{}

	count := n
	n--
	root = &tree.NumericNode{Data: (count - n)}

	stack.Push(root)

	for n > 0 {
		node := stack.Dequeue().(*tree.NumericNode)
		n--
		node.Left = &tree.NumericNode{Data: count - n}
		stack.Push(node.Left)
		if n > 0 {
			n--
			node.Right = &tree.NumericNode{Data: count - n}
			stack.Push(node.Right)
		}
	}

	return root
}
