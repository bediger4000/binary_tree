package main

import (
	"binary_tree/tree"
	"fmt"
	"io"
	"os"
	"strconv"
)

// A binary tree node can be locked or unlocked only if all of its
// descendants or ancestors are not locked.

type LockNode struct {
	Data   int
	Left   *LockNode
	Right  *LockNode
	Parent *LockNode
	Locked bool
}

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

func (node *LockNode) String() string {
	note := 'U'
	if node.Locked {
		note = 'L'
	}
	return fmt.Sprintf("%d/%c", node.Data, note)
}

// IsLocked returns whether the node is locked
func (node *LockNode) IsLocked() bool {
	return node.Locked
}

// lock, which attempts to lock the node.  If it cannot be locked,
// then it should return false.
func (node *LockNode) Lock() bool {
	// A binary tree node can be locked only if all of its ancestors are not
	// locked.
	ancestor := node.Parent
	for ancestor != nil {
		if ancestor.IsLocked() {
			return false
		}
		ancestor = ancestor.Parent
	}
	// no ancestors are locked

	// this call will not allow the tree to be locked if this node is already locked
	if descendantsLocked(node) {
		return false
	}

	node.Locked = true
	return true
}

// descendantsLocked returns true is any descendant of node
// is locked.
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
	// A binary tree node can be unlocked only if all of its ancestors are not
	// locked.
	ancestor := node.Parent
	for ancestor != nil {
		if ancestor.IsLocked() {
			fmt.Printf("ancestor of node value %d,  value %d, locked\n", node.Data, ancestor.Data)
			return false
		}
		ancestor = ancestor.Parent
	}
	// no ancestors are locked
	fmt.Printf("ancestors of node value %d unlocked\n", node.Data)

	if descendantsLocked(node.Left) {
		fmt.Printf("left descendants of node value %d locked\n", node.Data)
		return false
	}
	fmt.Printf("left descendants of node value %d unlocked\n", node.Data)
	if descendantsLocked(node.Right) {
		fmt.Printf("right descendants of node value %d locked\n", node.Data)
		return false
	}
	fmt.Printf("right descendants of node value %d unlocked\n", node.Data)

	node.Locked = false
	return true
}

func CheckAll(node *LockNode) {
	if node == nil {
		return
	}

	lockphrase := ""
	if !node.IsLocked() {
		lockphrase = "not "
	}

	CheckAll(node.Left)
	fmt.Printf("node value %d %slocked\n", node.Data, lockphrase)
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
	for _, str := range os.Args[1:] {
		n, err := strconv.Atoi(str)
		if err == nil {
			root = Insert(root, n)
		}
	}
	tree.Printf(os.Stdout, root)
	fmt.Println()

READLOOP:
	for {
		fmt.Printf("%d > ", root.Data)
		var value int
		var operation string
		nscanned, err := fmt.Scanf("%s %d\n", &operation, &value)
		if err == io.EOF {
			fmt.Printf("EOF on read\n")
			return
		}
		if nscanned != 1 && err != nil {
			fmt.Printf("Failed to read: %v\n", err)
			return
		}

		switch operation {
		case "quit":
			break READLOOP

		case "find":
			node := Find(root, value)
			if node != nil {
				fmt.Printf("found node with value %d at %p, locked %v\n", node.Data, node, node.Locked)
			}

		case "checkall":
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
