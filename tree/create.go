package tree

import (
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
		fmt.Fprintf(out, "()")
		return
	}
	fmt.Fprintf(out, "(%s", node) // relies on node fitting fmt.Stringer
	if !node.LeftChild().IsNil() || !node.RightChild().IsNil() {
		Printf(out, node.LeftChild())
		Printf(out, node.RightChild())
	}
	fmt.Fprintf(out, ")")
}

// GeneralCreateFromString uses a func argument to create a tree
// of type Node. It returns the root Node, on which the caller should
// do a type assertion to get the correct type. This func basically
// sets up to call genericTreeFromString with details that the caller
// shouldn't have to know.
func GeneralCreateFromString(stringrep string, nc NodeCreatorFn) Node {
	runes := []rune(stringrep)
	root, _ := genericTreeFromString(runes, 0, len(runes), nc)
	return root
}

func genericTreeFromString(runes []rune, offset, end int, nodeCreator NodeCreatorFn) (Node, int) {
	xoffset := eatWhiteSpace(runes, offset, end)
	if runes[xoffset] != '(' {
		fmt.Fprintf(os.Stderr, "no leading ( on %q at offset %d\n", string(runes), xoffset)
	}
	xend := findRightParen(runes, xoffset)
	xoffset++ // skip '('

	// eat whitespace after '(' but before string or number or whatever
	xoffset = eatWhiteSpace(runes, xoffset, xend)

	str, xoffset := readStuff(runes, xoffset, xend)

	var node Node
	if len(str) > 0 {
		node = nodeCreator(str)
	} else {
		xoffset = eatWhiteSpace(runes, xoffset, xend)
		xoffset++ // eat ')'
		return nil, xoffset
	}

	// eat whitespace after string or number or whatever, before '(' or ')'
	xoffset = eatWhiteSpace(runes, xoffset, xend)

	if runes[xoffset] == ')' {
		xoffset++
		return node, xoffset
	}

	leftChild, xoffset := genericTreeFromString(runes, xoffset, xend, nodeCreator)
	node.SetLeftChild(leftChild)

	if runes[xoffset] == ')' {
		xoffset++
		return node, xoffset
	}

	rightChild, xoffset := genericTreeFromString(runes, xoffset, xend, nodeCreator)
	node.SetRightChild(rightChild)

	xoffset = eatWhiteSpace(runes, xoffset, xend)
	xoffset++ // eat trailing ')'

	return node, xoffset
}

// readStuff returns a non-whitespace, non-parentheses string
// from runes []rune, and an index of where that string ends.
func readStuff(runes []rune, offset int, end int) (string, int) {
	var valueRunes []rune
	for {
		if runes[offset] == '(' {
			break
		}
		if runes[offset] == ')' {
			break
		}
		if offset == end {
			break
		}
		if unicode.IsSpace(runes[offset]) {
			break
		}
		valueRunes = append(valueRunes, runes[offset])
		offset++
	}

	return string(valueRunes), offset
}

// eatWhiteSpace starts at index offset in runes,
// and returns the index of the next non-whitespace rune.
func eatWhiteSpace(runes []rune, offset int, end int) int {
	for unicode.IsSpace(runes[offset]) {
		offset++
		if offset == end {
			break
		}
	}
	return offset
}

// findRightParen takes an array of runes, where '(' is at
// index 0, and a matching ')' is at some greater index.
// Returns that greater index
func findRightParen(r []rune, offset int) int {
	stack := make([]rune, 1)
	stack[0] = r[offset]
	end := offset

	for idx := offset + 1; len(stack) > 0; idx++ {
		switch r[idx] {
		case '(':
			stack = append(stack, '(')
		case ')':
			if stack[len(stack)-1] == '(' {
				stack = stack[0 : len(stack)-1]
			}
		}
		end++
	}
	return end + 1
}
