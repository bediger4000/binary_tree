package main

import (
	"binary_tree/tree"
	"fmt"
	"os"
)

// The first element of preorder array has the value of the
// node to be constructed.
// inorder has that value somewhere in it, and the index of
// that value divides inorder into values of nodes in LH sub-tree,
// and values of nodes in RH sub-tree.
// The only tricky parts are:
// when to stop recursing, and how to cut off enough of preorder
// array so that a recursive call to insert() has the correct preorder
// array, especially for the RH sub-tree.
func insert(preorder, inorder []string) *tree.StringNode {
	// preorder is always longer than it needs to be,
	// not trimming off RH sub-tree values from it.
	if len(inorder) == 0 {
		return nil
	}
	node := &tree.StringNode{Data: preorder[0]}

	for idx, str := range inorder {
		if str != node.Data {
			continue
		}
		// Found idx, the index into inorder slice of this node's
		// value. Create left and right subtrees, then exit the loop
		node.Left = insert(preorder[1:], inorder[:idx])
		// have to trim as many nodes as exist in the LH sub-tree,
		// plus 1 for the current node from preorder.
		node.Right = insert(preorder[idx+1:], inorder[idx+1:])
		break
	}
	return node
}

func main() {

	var orig, root *tree.StringNode
	var err error

	preorder := []string{"a", "b", "d", "e", "c", "f", "g"}
	inorder := []string{"d", "b", "e", "a", "f", "c", "g"}

	if len(os.Args) > 1 {
		// A lisp-like string rep of a tree on command line.
		if orig, err = tree.CreateFromString(os.Args[1]); err != nil {
			fmt.Fprintf(os.Stderr, "problem parsing %q: %v\n", os.Args[1], err)
		}
		// Create pre- and in-order arrays from tree
		iorder := make(StringSlice, 0)
		porder := make(StringSlice, 0)
		tree.AllorderTraverseVisit(orig, porder.appendString, iorder.appendString, tree.NullVisitor)
		preorder = porder
		inorder = iorder
	}

	fmt.Println("/*")
	root = insert(preorder, inorder)
	fmt.Println("*/")

	if root != nil {
		if orig == nil {
			tree.Draw(root)
			return
		}
		fmt.Printf("digraph g1 {\n")
		fmt.Printf("subgraph cluster_0 {\n\tlabel=\"original\"\n")
		tree.DrawPrefixed(os.Stdout, orig, "o")
		fmt.Printf("\n}\n")
		fmt.Printf("subgraph cluster_1 {\n\tlabel=\"reconstructed\"\n")
		tree.DrawPrefixed(os.Stdout, root, "r")
		fmt.Printf("\n}\n")
		fmt.Printf("\n}\n")

		phrase := ""
		if !tree.Equals(root, orig) {
			phrase = "not "
		}
		fmt.Printf("/* original and reconstructed %sequal */\n", phrase)

	}
}

// StringSlice exists to allow using obj.appendString as a tree.VisitorFunc
// type: don't have to use globals to collect arrays of strings during tree
// traverses.
type StringSlice []string

func (ss *StringSlice) appendString(node tree.Node) {
	(*ss) = append(*ss, node.String())
}
