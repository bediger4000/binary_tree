package main

import (
	"binary_tree/tree"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

// A binary tree node can be locked or unlocked only if all of its
// descendants or ancestors are not locked.

type LockNode struct {
	Data              int
	Left              *LockNode
	Right             *LockNode
	Parent            *LockNode
	Locked            bool
	LockedDescendants int
}

// If *LockNode doesn't match interface tree.Node,
// this fails.
var _ tree.Node = (*LockNode)(nil)

// IsNil, LeftChild, RightChild, SetLeftChild, SetRightChild
// make *LockNode fit interface tree.Node.
func (node *LockNode) IsNil() bool {
	return node == nil
}
func (node *LockNode) LeftChild() tree.Node {
	return node.Left
}
func (node *LockNode) RightChild() tree.Node {
	return node.Right
}
func (node *LockNode) SetLeftChild(child tree.Node) {
	if c, ok := child.(*LockNode); ok {
		node.Left = c
	}
}
func (node *LockNode) SetRightChild(child tree.Node) {
	if c, ok := child.(*LockNode); ok {
		node.Right = c
	}
}

// String makes a *LockNode into a fmt.Strinter,
// which is what tree.Printf expects.
func (node *LockNode) String() string {
	note := 'U'
	if node.Locked {
		note = 'L'
	}
	return fmt.Sprintf("%d/%c/%d", node.Data, note, node.LockedDescendants)
}

// createLockNode used when parsing a string from the "create"
// tree explorer operation.
func createLockNode(str string) tree.Node {
	n, err := strconv.Atoi(str)
	if err != nil {
		log.Print(err)
		return nil
	}
	return &LockNode{Data: n}
}

// addParents called after tree.GeneralCreateFromString invocation,
// since that function doesn't think it's doing trees with nodes
// with a parent back pointer.
func addParents(node *LockNode) {
	if node.Left != nil {
		node.Left.Parent = node
		addParents(node.Left)
	}
	if node.Right != nil {
		node.Right.Parent = node
		addParents(node.Right)
	}
}

// funcs IsLocked, Lock, Unlock meet the requirements of this problem

// IsLocked returns whether the node is locked
func (node *LockNode) IsLocked() bool {
	return node.Locked
}

// Lock, which attempts to lock the node.  If it cannot be locked,
// then it should return false.
func (node *LockNode) Lock() bool {

	if node.LockedDescendants > 0 {
		// one or more descendants are locked
		fmt.Printf("%d descendants locked\n", node.LockedDescendants)
		return false
	}

	// A binary tree node can be locked only if all of its ancestors are not
	// locked.

	for ancestor := node.Parent; ancestor != nil; ancestor = ancestor.Parent {
		if ancestor.IsLocked() {
			fmt.Printf("ancestor with value %d locked\n", ancestor.Data)
			return false
		}
	}

	node.Locked = true
	for ancestor := node.Parent; ancestor != nil; ancestor = ancestor.Parent {
		ancestor.LockedDescendants++
	}
	return true
}

// descendantsLocked returns true if node or any descendant of node are locked.
// This is the unoptimized, whole-sub-tree-search version that doesn't
// fit the O(h) requirement.
func descendantsLocked(node *LockNode) bool {
	if node == nil {
		return false
	}
	if node.IsLocked() {
		return true
	}
	if descendantsLocked(node.Left) {
		return false
	}
	return descendantsLocked(node.Right)
}

// Unlock unlocks the node.  If it cannot be unlocked, then it
// should return false.  Otherwise, it should unlock it and return
// true.
func (node *LockNode) Unlock() bool {

	if node.LockedDescendants > 0 {
		// one or more descendants are locked
		fmt.Printf("%d descendants locked\n", node.LockedDescendants)
		return false
	}

	// A binary tree node can be unlocked only if all of its ancestors are not
	// locked.
	for ancestor := node.Parent; ancestor != nil; ancestor = ancestor.Parent {
		if ancestor.IsLocked() {
			fmt.Printf("ancestor with value %d locked\n", ancestor.Data)
			return false
		}

	}

	node.Locked = false
	for ancestor := node.Parent; ancestor != nil; ancestor = ancestor.Parent {
		ancestor.LockedDescendants--
	}
	return true
}

// func CheckAll, Find, used in the "tree explorer" part
// of the code. Most of the tree explorer just calls library
// functions, or Lock and Unlock

func CheckAll(node *LockNode) {
	if node == nil {
		return
	}

	lockphrase := ""
	if !node.IsLocked() {
		lockphrase = "not "
	}

	CheckAll(node.Left)
	fmt.Printf("node value %d @ %p (%d), parent %p %slocked\n", node.Data, node, node.LockedDescendants, node.Parent, lockphrase)
	CheckAll(node.Right)
}

func Find(root *LockNode, value int) *LockNode {
	if root == nil {
		return nil
	}
	if root.Data == value {
		return root
	}
	if nd := Find(root.Left, value); nd != nil {
		return nd
	}
	return Find(root.Right, value)
}

// Insert puts nodes into a tree, keeping Binary Search Tree property
func Insert(node *LockNode, value int) *LockNode {
	if node == nil {
		return &LockNode{Data: value}
	}
	if node.Data > value {
		node.Left = Insert(node.Left, value)
		node.Left.Parent = node
	} else {
		node.Right = Insert(node.Right, value)
		node.Right.Parent = node
	}
	return node
}

func main() {
	var root *LockNode
	if len(os.Args) > 1 {
		for _, str := range os.Args[1:] {
			n, err := strconv.Atoi(str)
			if err == nil {
				root = Insert(root, n)
			}
		}
		fmt.Println("Input tree: ")
		tree.Printf(os.Stdout, root)
		fmt.Println()
	}

	fmt.Printf("Locked binary tree explorer\n")
READLOOP:
	for {
		fmt.Printf("> ")
		var valueString string
		var operation string
		nscanned, err := fmt.Scanf("%s %s\n", &operation, &valueString)
		if err == io.EOF {
			fmt.Printf("EOF on read\n")
			return
		}
		if nscanned != 1 && err != nil {
			fmt.Printf("Failed to read: %v\n", err)
			return
		}
		var value int
		switch operation {
		case "find", "check", "lock", "unlock":
			value, err = strconv.Atoi(valueString)
			if err != nil {
				log.Print(err)
			}
		}

		switch operation {
		case "?":
			usage()
		case "quit":
			break READLOOP

		case "create":
			if valueString == "" {
				break
			}
			_, tmp, err := tree.GeneralCreateFromString([]rune(strings.TrimSpace(valueString)), createLockNode)
			if err != nil {
				fmt.Fprintf(os.Stderr, "parsing %q: %v\n", valueString, err)
				break
			}
			root = tmp.(*LockNode)
			addParents(root)

		case "find":
			node := Find(root, value)
			if node != nil {
				fmt.Printf("found node with value %d at %p, locked %v\n", node.Data, node, node.Locked)
			}

		case "checkall":
			if root != nil {
				fmt.Printf("root is node with value %d\n", root.Data)
			}
			CheckAll(root)

		case "check":
			node := Find(root, value)
			if node == nil {
				fmt.Printf("did not find node with value %d\n", value)
				continue
			}
			phrase := "is not"
			if node.IsLocked() {
				phrase = "is"
			}
			fmt.Printf("node with value %d at %p %s locked\n", node.Data, node, phrase)
		case "lock":
			node := Find(root, value)
			if node == nil {
				fmt.Printf("did not find node with value %d\n", value)
				continue
			}
			phrase := "did not lock"
			if node.Lock() {
				phrase = "locked"
			}
			fmt.Printf("%s node with value %d at %p\n", phrase, node.Data, node)

		case "unlock":
			node := Find(root, value)
			if node == nil {
				fmt.Printf("did not find node with value %d\n", value)
				continue
			}
			phrase := "did not unlock"
			if node.Unlock() {
				phrase = "unlocked"
			}
			fmt.Printf("%s node with value %d at %p\n", phrase, node.Data, node)

		case "print":
			tree.Printf(os.Stdout, root)
			fmt.Println()

		default:
			fmt.Printf("Operation: %q\n", operation)
			fmt.Printf("Value: %d\n", value)
		}
	}
}
func usage() {
	fmt.Printf("locking node binary tree explorer\n")
	fmt.Printf("Operations:\n")
	fmt.Printf("print - print lisp-like string rep of tree\n")
	fmt.Printf("checkall - show lock status of all nodes\n")
	fmt.Printf("check N - show lock status of node with value N\n")
	fmt.Printf("lock N - lock node with value N\n")
	fmt.Printf("unlock N - unlock node with value N\n")
	fmt.Printf("find N - print info about node with value N\n")
	fmt.Printf("create (...) - parse lisp-like tree rep, use it thereafter\n")
}
