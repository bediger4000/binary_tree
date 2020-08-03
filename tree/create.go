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

type runeStack struct {
	array []rune
}

func (rs *runeStack) push(r rune) {
	rs.array = append(rs.array, r)
}

func (rs *runeStack) empty() bool {
	return len(rs.array) == 0
}

func (rs *runeStack) pop() rune {
	l := len(rs.array) - 1
	r := rs.array[l]
	rs.array = rs.array[:l]
	return r
}

func findIndex(str []rune, startIdx, endIdx int) int {
	if startIdx > endIdx {
		return -1
	}

	stk := &runeStack{}

	for i := startIdx; i <= endIdx; i++ {
		if str[i] == '(' {
			stk.push(str[i])
		} else if str[i] == ')' {
			if stk.array[len(stk.array)-1] == '(' {
				stk.pop()

				if stk.empty() {
					return i
				}
			}
		}
	}
	return -1
}

func treeFromString(str []rune, startIdx, endIdx int) *StringNode {
	if startIdx > endIdx {
		return nil
	}

	identifier := findIdentifier(str, startIdx, endIdx)
	idLen := len(identifier)
	if idLen == 0 {
		return nil
	}
	node := &StringNode{Data: identifier}
	index := -1

	if startIdx+idLen <= endIdx && str[startIdx+idLen] == '(' {
		index = findIndex(str, startIdx+idLen, endIdx)
	}

	if index != -1 {
		node.Left = treeFromString(str, startIdx+idLen+1, index)
		node.Right = treeFromString(str, index+2, endIdx)
	}

	return node
}

func numericTreeFromString(str []rune, startIdx, endIdx int) *NumericNode {
	if startIdx > endIdx {
		return nil
	}

	identifier := findIdentifier(str, startIdx, endIdx)
	idLen := len(identifier)
	if idLen == 0 {
		return nil
	}
	num, err := strconv.ParseInt(identifier, 10, 64)
	if err != nil {
		return nil
	}
	node := &NumericNode{Data: int(num)}
	index := -1

	if startIdx+idLen <= endIdx && str[startIdx+idLen] == '(' {
		index = findIndex(str, startIdx+idLen, endIdx)
	}

	if index != -1 {
		node.Left = numericTreeFromString(str, startIdx+idLen+1, index)
		node.Right = numericTreeFromString(str, index+2, endIdx)
	}

	return node
}

func findIdentifier(str []rune, startIdx, endIdx int) string {
	var identifier []rune
	for i := startIdx; i < endIdx; i++ {
		if unicode.IsLetter(str[i]) || unicode.IsDigit(str[i]) {
			identifier = append(identifier, str[i])
		} else {
			return string(identifier)
		}
	}
	return string(identifier)
}

// CreateFromString parses a single string
// like "(abc(ghi()(jkl))(def(pork)(beans)))"
// and turns it into a binary tree.
func CreateFromString(stringrep string) (root *StringNode) {
	runes := []rune(stringrep)
	l := len(runes)
	return treeFromString(runes[1:l-1], 0, l)
}

// CreateNumericFromString parses a single string
// like "(2(0()(12))(34(-2)(100)))"
// and turns it into a binary tree of the given shape.
func CreateNumericFromString(stringrep string) (root *NumericNode) {
	runes := []rune(stringrep)
	l := len(runes)
	return numericTreeFromString(runes[1:l-1], 0, l)
}

// Print writes out a tree in the format that
// CreateByParsing can turn into a tree.
// Re-uses interface Node
func Print(node Node) {
	Printf(os.Stdout, node)
}

// Printf writes a tree on "out" in the format that
// CreateByParsing can turn into a tree.
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

func GeneralCreateFromString(stringrep string, nc NodeCreatorFn) Node {
	runes := []rune(stringrep)
	l := len(runes)
	return genericTreeFromString(runes[1:l-1], 0, l, nc)
}

func genericTreeFromString(str []rune, startIdx, endIdx int, nodeCreator NodeCreatorFn) Node {
	if startIdx > endIdx {
		return nil
	}

	identifier := findIdentifier(str, startIdx, endIdx)
	idLen := len(identifier)
	if idLen == 0 {
		return nil
	}
	node := nodeCreator(identifier)
	index := -1

	if startIdx+idLen <= endIdx && str[startIdx+idLen] == '(' {
		index = findIndex(str, startIdx+idLen, endIdx)
	}

	if index != -1 {
		node.SetLeftChild(genericTreeFromString(str, startIdx+idLen+1, index, nodeCreator))
		node.SetRightChild(genericTreeFromString(str, index+2, endIdx, nodeCreator))
	}

	return node
}
