package main

import (
	"binary_tree/tree"
	"fmt"
	"os"
)

/* Daily Coding Problem: Problem #540 [Easy]

In Ancient Greece, it was common to write text with the first line going left
to right, the second line going right to left, and continuing to go back and
forth. This style was called "boustrophedon".

Given a binary tree, write an algorithm to print the nodes in boustrophedon order.

For example, given the following tree:

       1
     /   \
    /     \
  2         3
 / \       / \
4   5     6   7

You should return [1, 3, 2, 4, 5, 6, 7].
*/

// Stack constitutes a push-down stack of binary tree nodes
type Stack struct {
	array []*tree.Node
}

// Heading indicates which direction the program currently
// traverses a layer (nodes of the same depth) of a tree
type Heading int

// Right and Left - the two directions this program traverses
// layers of a binary tree.
const (
	Right Heading = 0
	Left  Heading = iota
)

func main() {
	root := tree.CreateNumeric(os.Args[1:])

	goingLeft := new(Stack)
	goingRight := new(Stack)

	stack := goingRight
	heading := Right

	stack.push(root)

	for !stack.empty() {

		switch heading {
		case Right:
			n := stack.pop()
			fmt.Printf("%d ", n.Data)
			goingLeft.push(n.Left)
			goingLeft.push(n.Right)
		case Left:
			n := stack.pop()
			fmt.Printf("%d ", n.Data)
			goingRight.push(n.Right)
			goingRight.push(n.Left)
		}

		if stack.empty() {
			switch heading {
			case Left:
				heading = Right
				stack = goingRight
			case Right:
				heading = Left
				stack = goingLeft
			}
		}
	}
	fmt.Println()
}

func (nq *Stack) push(n *tree.Node) {
	if n == nil {
		return
	}
	nq.array = append(nq.array, n)
}

func (nq *Stack) pop() (tail *tree.Node) {
	l := len(nq.array) - 1
	if l < 0 {
		return
	}
	tail = nq.array[l]
	nq.array = nq.array[:l]
	return
}

func (nq *Stack) empty() bool {
	if len(nq.array) == 0 {
		return true
	}
	return false
}
