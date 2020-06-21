package main

/*
 * Given a complete binary tree, count the number of nodes in faster than O(n)
 * time. Recall that a complete binary tree has every level filled except the
 * last, and the nodes in the last level are filled starting from the left.
 * "Complete" means: every level, except possibly the last, is completely
 * filled, and all nodes in the last level are as far left as possible. It
 * can have between 1 and 2h nodes at the last level h.
 */

import (
	"binary_tree/tree"
	"fmt"
	"math"
	"os"
)

func main() {
	root := tree.CreateNumeric(os.Args[1:])
	nodeCount := len(os.Args[1:])

	leftdepth := tree.LeftDepth(root)
	rightdepth := tree.RightDepth(root)
	nodesTouched := leftdepth + rightdepth

	fmt.Printf("Left depth %d, right depth %d\n", leftdepth, rightdepth)

	if leftdepth == rightdepth {
		fmt.Printf("Full tree, depth %d\n", leftdepth)
		theoretical := math.Pow(2.0, float64(leftdepth)) - 1.0
		fmt.Printf("%d nodes in tree\ntheoretical node count %.1f\n", nodeCount, theoretical)
		fmt.Printf("%d nodes in tree, %d nodes accessed\n", nodeCount, nodesTouched)
		return
	}

	leftturns := newTurnsZero(leftdepth - 1)
	rightturns := newTurnsOnes(leftdepth - 1)

	// saving this for later
	count := int(math.Pow(2.0, float64(rightdepth))) - 1

	var mid turns

	for {
		fmt.Printf("\nLeft  %v, depth %d\n", leftturns, leftdepth)
		fmt.Printf("Right %v, depth %d\n", rightturns, rightdepth)
		mid = addAndHalf(leftturns, rightturns)
		middepth := tree.ProbeDepth(root, mid)
		nodesTouched += middepth
		fmt.Printf("Mid   %v, depth %d\n", mid, middepth)

		if equal(mid, leftturns) {
			fmt.Printf("LFound it: %v and %v\n", leftturns, mid)
			break
		}
		if equal(mid, rightturns) {
			fmt.Printf("RFound it: %v and %v\n", rightturns, mid)
			break
		}

		if middepth < leftdepth {
			fmt.Printf("right becomes mid\n")
			rightturns = mid
			rightdepth = middepth
			continue
		}
		fmt.Printf("left becomes mid\n")
		leftturns = mid
		leftdepth = middepth
	}

	n := mid.toNumber()
	fmt.Printf("last node at depth %d has number %b\n", leftdepth, n)
	// n is a number with binary digits of mid.
	// n does not equal the number of nodes in a complete
	// tree. n does number the leaf nodes of the tree
	// left to right, 0-indexed.
	fmt.Printf("%d nodes in tree, %d nodes accessed\n", count+n+1, nodesTouched)
}

// turns represent left/right child node choices when
// descending a binary tree. Also treated as a binary
// number with each digit an element of the array.
// When treated as a number, the most significant digit
// is at index 0, least significant digit is at max index.
type turns []int

func newTurnsZero(size int) turns {
	t := make([]int, size)
	return t
}

func newTurnsOnes(size int) turns {
	t := make([]int, size)
	for i := range t {
		t[i] = 1
	}
	return t
}

// toNumber makes a single int out of the array
// of 1s and 0s that comprises an instance of turns
func (t turns) toNumber() int {
	var n int
	for _, digit := range t {
		n = (n << 1) | digit
	}
	return n
}

// equal could be thought of as comparing two instances
// of turns on an element-by-element basis, and returning
// true if all those elements numerically equate.
func equal(t1, t2 turns) bool {
	l := len(t1) - 1
	for i := range t1 {
		idx := l - i
		if t1[idx] != t2[idx] {
			return false
		}
	}
	return true
}

// addAndHalf add two instances of turns together,
// keeping the carry digit - one extra digit of precision,
// because it then shifts right one bit. If you don't keep
// the carry bit, you get wrong answers for the purposes of
// this program.
func addAndHalf(t1, t2 turns) turns {
	rt := newTurnsZero(len(t1))
	l := len(t1) - 1
	carry := 0
	for i := range t1 {
		idx := l - i
		n := t1[idx] + t2[idx] + carry
		carry = 0
		if n > 1 {
			carry = 1
		}
		rt[idx] = n % 2
	}
	tmp := rt[0]
	rt[0] = carry
	for i := 1; i < len(rt); i++ {
		rt[i], tmp = tmp, rt[i]
	}
	return rt
}
