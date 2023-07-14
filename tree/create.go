package tree

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
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
	if node.IsNil() {
		return
	}
	leaf := node.LeftChild().IsNil() && node.RightChild().IsNil()
	if !leaf {
		fmt.Fprint(out, "(")
	}
	fmt.Fprintf(out, "%s", node)
	if !node.LeftChild().IsNil() {
		fmt.Fprint(out, " ")
	} else if !node.RightChild().IsNil() {
		fmt.Fprint(out, " ()")
	}
	Printf(out, node.LeftChild())
	if !node.RightChild().IsNil() {
		fmt.Fprint(out, " ")
	}
	Printf(out, node.RightChild())
	if !leaf {
		fmt.Fprint(out, ")")
	}
}

// GeneralCreateFromString uses a func argument to create a tree
// of type Node. It returns the root Node, on which the caller should
// do a type assertion to get the correct type. This func basically
// sets up to call genericTreeFromString with details that the caller
// shouldn't have to know.
func GeneralCreateFromString(stringrep string, nc NodeCreatorFn) Node {
	tokens := tokenize(stringrep)
	root, _ := genericTreeFromString(tokens, 0, nc)
	return root
}

func genericTreeFromString(tokens []string, index int, nc NodeCreatorFn) (Node, int) {
	l := len(tokens)
	if index >= l || l == 0 {
		return nil, index
	}

	var n Node
	looping := true
	setLeft := false

	for looping && index < l {
		switch tokens[index] {
		case "(":
			index++
			child, nextindex := genericTreeFromString(tokens, index, nc)
			index = nextindex
			if n == nil {
				n = child
			} else if setLeft {
				n.SetRightChild(child)
			} else {
				n.SetLeftChild(child)
				setLeft = true
			}
		case ")":
			looping = false
			index++
		default:
			if n == nil {
				n = nc(tokens[index])
			} else {
				// else, naked string child
				child := nc(tokens[index])
				if setLeft {
					n.SetRightChild(child)
				} else {
					setLeft = true
					n.SetLeftChild(child)
				}
			}
			index++
		}
	}
	return n, index
}

func tokenize(str string) []string {
	var tokens []string
	for _, token := range strings.Split(strings.ReplaceAll(strings.ReplaceAll(str, "(", " ( "), ")", " ) "), " ") {
		if token != "" {
			tokens = append(tokens, token)
		}
	}
	return tokens
}
