package main

import (
	"binary_tree/tree"
	"bytes"
	"fmt"
	"log"
	"os"
)

/*
Given the root to a binary tree,
implement serialize(root),
which serializes the tree into a string,
and deserialize(s),
which deserializes the string back into the tree.

For example, given the following Node class

class Node:
    def __init__(self, val, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

The following test should pass:

node = Node('root', Node('left', Node('left.left')), Node('right'))
assert deserialize(serialize(node)).left.left.val == 'left.left'

*/

func main() {

	root := &tree.StringNode{
		Data: "root",
		Left: &tree.StringNode{
			Data: "left",
			Left: &tree.StringNode{
				Data: "left.left",
			},
		},
		Right: &tree.StringNode{
			Data: "right",
		},
	}

	buf := &bytes.Buffer{}
	tree.Printf(buf, root)
	treeAsString := buf.String()
	fmt.Printf("Tree as string: %q\n", treeAsString)
	newroot, err := tree.CreateFromString(treeAsString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "parsing %q: %v\n", treeAsString, err)
		return
	}
	if newroot.Left.Left.Data != "left.left" {
		log.Fatal("test failed")
	}

	fmt.Println("test passed")
}
