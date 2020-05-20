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

	goingLeft := new(tree.Stack)
	goingRight := new(tree.Stack)

	stack := goingRight
	heading := Right

	stack.Push(root)

	for !stack.Empty() {

		switch heading {
		case Right:
			n := stack.Pop()
			fmt.Printf("%d ", n.Data)
			goingLeft.Push(n.Left)
			goingLeft.Push(n.Right)
		case Left:
			n := stack.Pop()
			fmt.Printf("%d ", n.Data)
			goingRight.Push(n.Right)
			goingRight.Push(n.Left)
		}

		if stack.Empty() {
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
