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

	fmt.Printf("Left depth %d, right depth %d\n", leftdepth, rightdepth)

	if leftdepth == rightdepth {
		fmt.Printf("Full tree, depth %d\n", leftdepth)
		theoretical := math.Pow(2.0, float64(leftdepth)) - 1.0
		fmt.Printf("%d nodes, theoretical node count %.1f\n", nodeCount, theoretical)
		return
	}

	leftturns := newTurns(leftdepth - 1)
	rightturns := allOnes(leftdepth - 1)

	for {
		fmt.Printf("\nLeft  %v, depth %d\n", leftturns, leftdepth)
		fmt.Printf("Right %v, depth %d\n", rightturns, rightdepth)
		mid := addAndHalf(leftturns, rightturns)
		middepth := tree.ProbeDepth(root, mid)
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
}

type turns []int

func newTurns(size int) turns {
	t := make([]int, size)
	return t
}

func allOnes(size int) turns {
	t := make([]int, size)
	for i := range t {
		t[i] = 1
	}
	return t
}

func (t turns) increment() bool {
	carry := true
	for i := 0; carry && i < len(t); i++ {
		t[i]++
		if t[i] == 1 {
			carry = false
		} else {
			t[i] = 0
		}
	}
	return carry
}

// halve t as a binary number by right shifting 1
func (t turns) half() {
	i := len(t) - 1
	last := t[i]
	t[i] = 0
	for i--; i >= 0; i-- {
		last, t[i] = t[i], last
	}
}

// halve t in the other direction
func (t turns) half2() {
	tmp := t[0]
	t[0] = 0
	for i := 1; i < len(t); i++ {
		t[i], tmp = tmp, t[i]
	}
}

func equal(t1, t2 turns) bool {
	for i := range t1 {
		if t1[i] != t2[i] {
			return false
		}
	}
	return true
}

func add(t1, t2 turns) turns {
	rt := newTurns(len(t1))
	carry := 0
	for i := range t1 {
		n := t1[i] + t2[i] + carry
		carry = 0
		if n > 1 {
			carry = 1
		}
		n = n % 2
		rt[i] = n
	}
	return rt
}

func addAndHalf(t1, t2 turns) turns {
	rt := newTurns(len(t1))
	carry := 0
	for i := range t1 {
		n := t1[i] + t2[i] + carry
		carry = 0
		if n > 1 {
			carry = 1
		}
		rt[i] = n % 2
	}
	tmp := rt[0]
	rt[0] = 0
	for i := 1; i < len(rt); i++ {
		rt[i], tmp = tmp, rt[i]
	}
	rt[len(rt)-1] = carry
	return rt
}
