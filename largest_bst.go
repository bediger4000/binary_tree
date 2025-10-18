package main

/*
	Given a tree, find the largest tree/subtree that is a BST.
	Given a tree, return the size of the largest tree/subtree that is a BST.
*/

import (
	"binary_tree/tree"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Count the number of univalue subtrees in a tree\n")
		fmt.Printf("Usage: %s '(1(0)(2))'\n", os.Args[0])
		return
	}
	root, err := tree.CreateNumericFromString(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "parsing %q: %v\n", os.Args[1], err)
		return
	}
	tree.Print(root)
	fmt.Println()
	n, _ := largestBST(root)
	fmt.Printf("Largest BST %d nodes\n", n)
}

func largestBST(node *tree.NumericNode) (biggest int, isBST bool) {
	if node == nil {
		return 0, true
	}

	// node not nil
	if node.Left == nil && node.Right == nil {
		// leaf node, BST of size 1
		return 1, true
	}

	leftCnt, leftBST := largestBST(node.Left)
	rightCnt, rightBST := largestBST(node.Right)

	if leftBST {
		biggest = leftCnt
	}

	if rightBST {
		if rightCnt > biggest {
			biggest = rightCnt
		}
	}

	if leftBST && rightBST {
		if (node.Left == nil || (node.Left != nil && node.Data > node.Left.Data)) &&
			(node.Right == nil || (node.Right != nil && node.Data < node.Right.Data)) {
			biggest = leftCnt + rightCnt + 1
			isBST = true
		}
	}

	return
}
