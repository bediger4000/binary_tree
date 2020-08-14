package main

import (
	"binary_tree/tree"
	"fmt"
	"log"
	"os"
	"strconv"
)

/*
  Given a binary search tree and a range [a, b] (inclusive),
  return the sum of the elements of the binary search tree within
  the range.
*/

func main() {
	a, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	b, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Summing node values [%d, %d]\n", a, b)

	root := tree.CreateNumeric(os.Args[3:])
	tree.Printf(os.Stdout, root)
	fmt.Println()
	if tree.BstProperty(root) {
		fmt.Printf("tree has binary search property\n")
	}

	var sum int
	visit1(root, a, b, &sum)
	fmt.Printf("Method 1 Sum %d\n", sum)
	fmt.Printf("visited %d nodes\n", visitCount)

	visitCount = 0
	sum = visit2(root, a, b)
	fmt.Printf("Method 2 Sum %d\n", sum)
	fmt.Printf("visited %d nodes\n", visitCount)
}

var visitCount int

func visit1(node *tree.NumericNode, a, b int, sum *int) {
	if node == nil {
		return
	}
	visitCount++
	if node.Data > b {
		return // minor optimization
	}
	if node.Data >= a && node.Data <= b {
		*sum += node.Data
	}
	visit1(node.Left, a, b, sum)
	visit1(node.Right, a, b, sum)
}

func visit2(node *tree.NumericNode, a, b int) int {
	if node == nil {
		return 0
	}
	visitCount++
	sum := 0
	if node.Data >= a && node.Data <= b {
		sum += node.Data
	}
	if a <= node.Data {
		sum += visit2(node.Left, a, b)
	}
	if b >= node.Data {
		sum += visit2(node.Right, a, b)
	}

	return sum
}

// NodeByValue should return a pointer to a node in
// the tree with data of the value, or the node with the
// next value bigger
func NodeByValue(node *tree.NumericNode, value int) *tree.NumericNode {
	if node.Data == value {
		return node
	}

	if value < node.Data {
		if node.Left != nil {
			return NodeByValue(node.Left, value)
		}
	}
	if node.Right != nil {
		return NodeByValue(node.Right, value)
	}
	return nil
}
