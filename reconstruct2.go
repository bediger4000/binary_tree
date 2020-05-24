package main

import (
	"binary_tree/tree"
	"fmt"
)

func insert(preorder, inorder []string) *tree.StringNode {
	if len(inorder) == 0 {
		return nil
	}
	node := &tree.StringNode{Data: preorder[0]}

	for idx, str := range inorder {
		if str != node.Data {
			continue
		}
		node.Left = insert(preorder[1:], inorder[:idx])
		node.Right = insert(preorder[idx+1:], inorder[idx+1:])
		break
	}
	return node
}

func main() {

	var root *tree.StringNode

	preorder := []string{"a", "b", "d", "e", "c", "f", "g"}
	inorder := []string{"d", "b", "e", "a", "f", "c", "g"}

	// root := tree.CreaFromString(os.Args[1])

	fmt.Println("/*")
	root = insert(preorder, inorder)
	fmt.Println("*/")

	if root != nil {
		tree.Draw(root)
	}
}
