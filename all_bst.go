package main

import (
	"binary_tree/tree"
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
)

/*
 * Given an integer N, construct all possible binary search trees with N nodes.
 * This isn't *all* binary trees, just valid BSTs
 * I'm assuming "N nodes" means N nodes, each with a unique integer value.
 */

func main() {
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	if n == 0 {
		fmt.Printf("()\n")
		return
	}

	c := make(chan string, 0)

	go createBSTs(c, n)

	found := make(map[string]bool)
	count, uniqCount := 0, 0

	for treeRep := range c {
		count++
		if found[treeRep] {
			continue
		}
		uniqCount++
		found[treeRep] = true
		fmt.Printf("%s\n", treeRep)
	}
	fmt.Printf("Found %d unique BSTs from %d candidates\n", uniqCount, count)
	cn := catalanNumber(n)
	fmt.Printf("should have found %d unique BSTs\n", cn)
}

func createBSTs(c chan string, n int) {
	// create array of N unique integers
	ary := make([]int, n)
	for i := 1; i <= n; i++ {
		ary[i-1] = i
	}

	// send off real recursive function
	generate(c, len(ary), ary)

	// When it's done close the channel
	close(c)
}

// generate comprises a Go transliteration of whatever
// pseudocode in which Wikipedia shows Heap's Algorithm.
// When Heap's algorithm gets to a permutation (k == 1),
// construct a BST using the integers in the array in order,
// serialize the tree to a string, send string to main() via channel
func generate(c chan string, k int, a []int) {
	if k == 1 {
		// construct a BST from permutation in a
		var root *tree.NumericNode
		for x := range a {
			root = tree.Insert(root, a[x])
		}

		// turn tree into a string, send down channel
		buf := &bytes.Buffer{}
		tree.Printf(buf, root)
		c <- buf.String()

		return
	}

	generate(c, k-1, a)

	for i := 0; i < k-1; i++ {
		if (k % 2) == 0 {
			a[i], a[k-1] = a[k-1], a[i]
		} else {
			a[0], a[k-1] = a[k-1], a[0]
		}
		generate(c, k-1, a)
	}
}

func catalanNumber(n int) int {
	if n == 0 {
		return 1
	}
	n--
	return 2 * (2*n + 1) * catalanNumber(n) / (n + 2)
}
