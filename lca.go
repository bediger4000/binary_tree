package main

import (
	"binary_tree/tree"
	"fmt"
	"os"
)

/*
 * Daily Coding Problem: Problem #545 [Hard]
 * Daily Coding Problem: Problem #112
 *
 * Given a binary tree, find the lowest common ancestor (LCA) of two
 * given nodes in the tree. Assume that each node in the tree also has a
 * pointer to its parent.
 *
 * According to the definition of LCA on Wikipedia: The lowest
 * common ancestor is defined between two nodes v and w as the lowest
 * node in T that has both v and w as descendants (where we allow a node
 * to be a descendant of itself).
 */

func main() {
	/* O(n) solution, n number of nodes in tree:
	 * 1. find path from root to v-valued node.
	 * 2. find path from root to w-valued node.
	 * 3. Compare paths. The last node common to both paths
	 *    is LCA.
	 * Does not use the link back to the parent node.
	 */
	valueV := os.Args[1]
	valueW := os.Args[2]
	stringrep := os.Args[3]
	fmt.Printf("Find LCA of %q and %q in tree ", valueV, valueW)
	root, err := tree.CreateFromString(stringrep)
	if err != nil {
		fmt.Fprintf(os.Stderr, "problem parsing %q: %v\n", stringrep, err)
		return
	}
	tree.Print(root)
	fmt.Println()

	var apath []string
	pathToV := path(root, valueV, apath)
	fmt.Printf("Path to %q: %v\n", valueV, pathToV)

	var bpath []string
	pathToW := path(root, valueW, bpath)
	fmt.Printf("Path to %q: %v\n", valueW, pathToW)

	lenV := len(pathToV)
	lenW := len(pathToW)

	if lenV == 0 || lenW == 0 {
		fmt.Println("No common ancestor, do both appear in tree?")
		return
	}

	lca := pathToV[0]
	lenMin := lenV
	if lenW < lenMin {
		lenMin = lenW
	}

	for i := 1; i < lenMin; i++ {
		if pathToV[i] != pathToW[i] {
			break
		}
		lca = pathToV[i]
	}

	fmt.Printf("Last common ancestor %q\n", lca)
}

// path returns either nil (no path from root to node with label)
// or an array of strings that denote labels of nodes between
// (and including) root and the labeled node.
// I think we have to traverse the tree, because the problem doesn't
// says the tree has any order imposed.
// That will take O(n/2) (n nodes in tree) on average to find the
// given value since it's unordered.
func path(node *tree.StringNode, label string, soFar []string) []string {
	if node == nil {
		return nil
	}

	if node.Data == label {
		finalpath := make([]string, len(soFar))
		copy(finalpath, soFar)
		finalpath = append(finalpath, node.Data)
		return finalpath
	}

	if node.Left == nil && node.Right == nil {
		return nil
	}

	mypath := make([]string, len(soFar))
	copy(mypath, soFar)
	mypath = append(mypath, node.Data)
	fp := path(node.Left, label, mypath)
	if len(fp) > 0 {
		return fp
	}
	return path(node.Right, label, mypath)
}
