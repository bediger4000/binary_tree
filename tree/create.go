package tree

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
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
func CreateFromString(stringrep string) (root *StringNode, err error) {
	runes := []rune(stringrep)
	origLength := len(runes)

	consumed, generic, err := GeneralCreateFromString(runes, createStringNode)

	if err != nil {
		return nil, err
	}

	if consumed != origLength {
		return nil, fmt.Errorf("string rep of length %d, consumed only %d runes",
			origLength, consumed)
	}

	var ok bool
	if root, ok = generic.(*StringNode); ok {
		return
	}
	return nil, errors.New("created tree of incorrect node type")
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
func CreateNumericFromString(stringrep string) (root *NumericNode, err error) {
	runes := []rune(stringrep)
	origLength := len(runes)

	consumed, generic, err := GeneralCreateFromString(runes, createNumericNode)

	if consumed != origLength {
		return nil, fmt.Errorf("string rep of length %d, consumed only %d runes",
			origLength, consumed)
	}

	var ok bool
	if root, ok = generic.(*NumericNode); ok {
		return
	}
	return nil, errors.New("created tree of nodes of incorrect type")
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
// do a type assertion to get the correct type. The int return is the
// number of runes read from the front of the runes []rune argument.
func GeneralCreateFromString(runes []rune, nc NodeCreatorFn) (int, Node, error) {

	if runes[0] != '(' {
		return 0, nil, errors.New("first character not opening parenthesis")
	}

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
			c, n, e := GeneralCreateFromString(runes[consumed:], nc)
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
		if setLeft {
			return consumed, nil, errors.New("no data value with child node(s)")
		}
		return consumed, nil, nil
	}

	newNode := nc(string(value))
	newNode.SetLeftChild(left)
	newNode.SetRightChild(right)

	return consumed, newNode, nil
}
