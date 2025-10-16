package tree

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"unicode"
)

// CreateNumeric converts a []string to a binary search tree
// with integer values of data at nodes. If it has a problem
// converting a string to an integer it prints a message on stderr
// and moves on, ignoring the unparseable string.
func CreateNumeric(numberRepr []string) (root *NumericNode) {

	for _, str := range numberRepr {
		val, err := strconv.Atoi(str)

		if err == nil {
			root = Insert(root, val)
		} else {
			fmt.Fprintf(os.Stderr, "Problem with %q: %s\n", str, err)
		}
	}

	return
}

// CreateFromString parses a single string
// like "(abc(ghi()(jkl))(def(pork)(beans)))"
// and turns it into a binary tree.
func CreateFromString(stringrep string) (root *StringNode) {
	generic := GeneralCreateFromString(stringrep, createStringNode)
	var ok bool
	if root, ok = generic.(*StringNode); ok {
		return
	}
	return nil
}

func createStringNode(stringValue string) Node {
	return &StringNode{Data: stringValue}
}

// createNumericNode fills in a struct NumericNode
// from a string argument, then returns a pointer to that NumericNode,
// except as something that fits interface Node.
func createNumericNode(stringValue string) Node {
	n, err := strconv.Atoi(stringValue)
	if err != nil {
		fmt.Fprintf(os.Stderr, "creating tree.Numeric node from %q: %v\n",
			stringValue, err)
		return nil
	}
	return &NumericNode{Data: n}
}

// CreateNumericFromString parses a single string
// like "(2(0()(12))(34(-2)(100)))"
// and turns it into a binary tree of the given shape.
func CreateNumericFromString(stringrep string) (root *NumericNode) {
	generic := GeneralCreateFromString(stringrep, createNumericNode)
	var ok bool
	if root, ok = generic.(*NumericNode); ok {
		return
	}
	return nil
}

// Print writes out a tree in the format that
// CreateByParsing can turn into a tree.
// Re-uses interface Node
func Print(node Node) {
	Printf(os.Stdout, node)
}

// Printf writes a tree on "out" in the format that CreateFromString or
// CreateNumericFromString can turn into a tree.
func Printf(out io.Writer, node Node) {
	fmt.Fprintf(out, "(%s", node)

	leftNil := node.LeftChild().IsNil()
	rightNil := node.RightChild().IsNil()

	if !leftNil && !rightNil {
		out.Write([]byte(" "))
		Printf(out, node.LeftChild())
		out.Write([]byte(" "))
		Printf(out, node.RightChild())
	} else if !leftNil && rightNil {
		out.Write([]byte(" "))
		Printf(out, node.LeftChild())
	} else if leftNil && !rightNil {
		out.Write([]byte(" () "))
		Printf(out, node.RightChild())
	} // else both child nodes empty
	out.Write([]byte(")"))
}

// GeneralCreateFromString uses a func argument to create a tree
// of type Node. It returns the root Node, on which the caller should
// do a type assertion to get the correct type. This func basically
// sets up to call genericTreeFromString with details that the caller
// shouldn't have to know.
func GeneralCreateFromString(stringrep string, nc NodeCreatorFn) Node {
	_, root, err := genericTreeFromString([]rune(strings.TrimSpace(stringrep)), nc)
	if err != nil {
		fmt.Fprintf(os.Stderr, "creating tree from string: %v\n", err)
	}
	return root
}

func genericTreeFromString(runes []rune, nc NodeCreatorFn) (int, Node, error) {

	var value []rune
	var left, right Node
	setLeft := false
	foundClosing := false

	max := len(runes)
	consumed := 1 // skip opening parentheses

loop:
	for consumed < max {

		switch runes[consumed] {
		case '(':
			c, n, e := genericTreeFromString(runes[consumed:], nc)
			if e != nil {
				return consumed, nil, e
			}
			consumed += c
			if !setLeft {
				left = n
				setLeft = true
			} else {
				right = n
			}
		case ')':
			consumed++
			foundClosing = true
			break loop
		default:
			if unicode.IsSpace(runes[consumed]) {
				consumed++
				continue
			}
			value = append(value, runes[consumed])
			consumed++
		}
	}

	if !foundClosing {
		return consumed, nil, errors.New("failed to find closing paren")
	}

	if len(value) == 0 {
		return consumed, nil, nil
	}

	newNode := nc(string(value))
	newNode.SetLeftChild(left)
	newNode.SetRightChild(right)

	return consumed, newNode, nil
}
