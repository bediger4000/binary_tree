package main

import (
	"binary_tree/tree"
	"fmt"
	"os"
	"sort"
	"strconv"
)

/*
The horizontal distance of a binary tree node describes how far left or right
the node will be when the tree is printed out.

More rigorously, we can define it as follows:

* The horizontal distance of the root is 0.
* The horizontal distance of a left child is hd(parent) - 1.
* The horizontal distance of a right child is hd(parent) + 1.

For example, for the following tree, hd(1) = -2, and hd(6) = 0.

               5
             /   \
           /       \
          3         7
        /  \      /   \
      1     4    6     9
     /                /
    0                8

The bottom view of a tree, then,
consists of the lowest node at each horizontal distance.
If there are two nodes at the same depth and horizontal distance,
either is acceptable.

For this tree, for example, the bottom view could be [0, 1, 3, 6, 8, 9].

Given the root to a binary tree, return its bottom view.
*/
// The tree package's string representation of the example tree is:
// (5 (3 (1 (0)) (4)) (7 (6) (9 (8))))

/* Uses tree.GeneralCreateFromString() to make a tree.
 * That means conforming to interface Node, and having a
 * function to create a ViewNode from a string.
 * struct ViewNode is almost a tree.NumericNode, but it
 * has Depth and Distance elements. That allows this program
 * to do the decorate-sort-undecorate pattern of behavior.
 */
type ViewNode struct {
	Data     int
	Left     *ViewNode
	Right    *ViewNode
	Distance int
	Depth    int
}

// Compiler checks if *ViewNode matches interface tree.Node
var _ tree.Node = (*ViewNode)(nil)

func (node *ViewNode) IsNil() bool {
	return node == nil
}
func (node *ViewNode) LeftChild() tree.Node {
	return node.Left
}
func (node *ViewNode) RightChild() tree.Node {
	return node.Right
}
func (node *ViewNode) SetLeftChild(child tree.Node) {
	if c, ok := child.(*ViewNode); ok {
		node.Left = c
	}
}
func (node *ViewNode) SetRightChild(child tree.Node) {
	if c, ok := child.(*ViewNode); ok {
		node.Right = c
	}
}

func CreateViewNode(stringValue string) tree.Node {
	n, err := strconv.Atoi(stringValue)
	if err != nil {
		return nil
	}
	return &ViewNode{Data: n}
}

func (n *ViewNode) String() string {
	return fmt.Sprintf("%d/%d/%d", n.Data, n.Depth, n.Distance)
}

func main() {

	// create a tree of *ViewNode
	root := tree.GeneralCreateFromString(os.Args[1], CreateViewNode)

	// Decorate the tree with depth in tree and "horizontal distance"
	decorate(root.(*ViewNode), 0, 0)
	m := make(map[int]*ViewNode)

	// Fill in a map with keys of nodes' horizontal distance,
	// and values of deepest *ViewNode at any given distance.
	traverse(root.(*ViewNode), m)

	// m now contains "lowest node at given horizontal difference"
	// but unordered.

	var nodeArray []*ViewNode
	for _, node := range m {
		nodeArray = append(nodeArray, node)
	}
	sort.Sort(NodeArray(nodeArray))

	for _, node := range nodeArray {
		fmt.Printf("%d ", node.Data)
	}
	fmt.Println()
}

// traverse walks a tree of ViewNodes, filling in
// the map m with the deepest ViewNode at any given
// horizontal distance.
// Keys of map m constitute horizontal distance.
func traverse(node *ViewNode, m map[int]*ViewNode) {
	if node == nil {
		return
	}

	/*
		The bottom view of a tree, then,
		consists of the lowest node at each horizontal distance.
		If there are two nodes at the same depth and horizontal distance,
		either is acceptable.
	*/
	dist := node.Distance
	if prevNode, ok := m[dist]; ok {
		// Could use '>' here, but example wants '>='
		if node.Depth >= prevNode.Depth {
			m[dist] = node
		}
	} else {
		m[dist] = node
	}

	traverse(node.Left, m)
	traverse(node.Right, m)
}

// decorate traverses a binary tree, setting depth in tree
// and horizonal distance to every node it visits.
// This could be done during tree construction, but I wanted
// to use the generic tree-construction function.
func decorate(node *ViewNode, depth, distance int) {
	if node == nil {
		return
	}
	node.Depth = depth
	node.Distance = distance
	decorate(node.Left, node.Depth+1, node.Distance-1)
	decorate(node.Right, node.Depth+1, node.Distance+1)
}

type NodeArray []*ViewNode

func (na NodeArray) Len() int           { return len(na) }
func (na NodeArray) Swap(i, j int)      { na[i], na[j] = na[j], na[i] }
func (na NodeArray) Less(i, j int) bool { return na[i].Distance < na[j].Distance }
