package main

import (
	"binary_tree/tree"
	"fmt"
	"math"
	"os"
)

type PathKeeper struct {
	path   []int
	minSum int
}

func main() {
	root := tree.CreateNumericFromString(os.Args[1])
	pk := &PathKeeper{minSum: math.MaxInt32}

	minPathSumTraverse(root, 0, []int{}, 0, pk)

	fmt.Printf("min path sum %d\n", pk.minSum)
	fmt.Printf("min sum path %v\n", pk.path)
}

func minPathSumTraverse(node *tree.NumericNode, depth int, path []int, sum int, pk *PathKeeper) {
	// return on nil node to avoid cluttering recursing code
	// with checks on nil pointers.
	if node == nil {
		return
	}
	// check sum and path on leaf nodes only.
	if node.Left == nil && node.Right == nil {
		sum += node.Data
		path = append(path, node.Data)
		if sum < pk.minSum {
			pk.path = make([]int, len(path))
			copy(pk.path, path)
			pk.minSum = sum
		}
		path = path[:depth]
		return
	}
	path = append(path, node.Data)
	minPathSumTraverse(node.Left, depth+1, path, sum+node.Data, pk)
	minPathSumTraverse(node.Right, depth+1, path, sum+node.Data, pk)
	path = path[:depth]
}
