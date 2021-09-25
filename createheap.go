package main

/*
 * Print lisp-like serialized tree with heap property.
 * Creates binary tree of *tree.NumericNode, then prints it.
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

	stack := &tree.Stack{}

	count := n
	n--
	root := &tree.NumericNode{Data: (count - n)}

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
