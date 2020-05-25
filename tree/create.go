package tree

import (
	"fmt"
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
// like "(abc (ghi () (jkl)) (def (pork) (beans)))"
// and turns it into a binary tree.
// Stupid name, will change it when I thik of something better.
func CreateFromString(stringrep string) (root *StringNode) {
	runes := []rune(stringrep)
	l := len(runes)
	return treeFromString(runes[1:l-1], 0, l)
}

// Print writes out a tree in the format that
// CreateByParsing can turn into a tree.
// Re-uses interface Node
func Print(node Node) {
	if node.isNil() {
		fmt.Printf("()")
		return
	}
	fmt.Printf("(%s", node) // relies on node fitting fmt.Stringer
	if !node.leftChild().isNil() || !node.rightChild().isNil() {
		Print(node.leftChild())
		Print(node.rightChild())
	}
	fmt.Print(")")
}
