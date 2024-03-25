package main

// Illustrate Wikipedia's "succint" encoding and decoding
// of binary trees

import (
	"binary_tree/tree"
	"bytes"
	"flag"
	"fmt"
	"os"
)

func main() {
	encode := flag.Bool("e", false, "encode the following lisp-like tree")
	decode := flag.Bool("d", false, "decode the following bit string")
	flag.Parse()

	if *encode && *decode {
		fmt.Fprintf(os.Stderr, "Only one of -d, -e allowed\n")
		flag.Usage()
		return
	}

	if !*encode && !*decode {
		fmt.Fprintf(os.Stderr, "need one of -d or -e, along with a tree string rep, or bit string respectively\n")
		flag.Usage()
		return
	}

	if flag.NArg() < 1 {
		if *encode {
			fmt.Fprintf(os.Stderr, "Need tree lisp-like string representation\n")
		} else {
			fmt.Fprintf(os.Stderr, "Need bit string representation of tree\n")
		}
		flag.Usage()
		return
	}

	if *encode {
		stringRep := flag.Arg(0)
		root := tree.CreateFromString(stringRep)
		bitString := EncodeSuccint(root)
		buf := &bytes.Buffer{}
		tree.Printf(buf, root)
		treeAsString := buf.String()
		fmt.Printf("%s\t%q\n", treeAsString, bitString)
		return
	}
	bitString := []rune(flag.Arg(0))

	root, n := DecodeSuccint(bitString)
	if n != len(bitString) {
		fmt.Fprintf(os.Stderr, "bit string has %d elements, used %d\n", len(bitString), n)
	}
	buf := &bytes.Buffer{}
	tree.Printf(buf, root)
	treeAsString := buf.String()
	fmt.Printf("%q\t%s\n", string(bitString), treeAsString)
}

func EncodeSuccint(node tree.Node) string {
	if node.IsNil() {
		return "0" // nil means "leaf node"
	}
	note := "1"
	note += EncodeSuccint(node.LeftChild())
	note += EncodeSuccint(node.RightChild())

	return note
}

func DecodeSuccint(bitString []rune) (*tree.StringNode, int) {

	switch bitString[0] {
	case '0':
		return nil, 1
	case '1':
		left, l := DecodeSuccint(bitString[1:])
		right, r := DecodeSuccint(bitString[l+1:])
		node := &tree.StringNode{
			Data:  "X",
			Left:  left,
			Right: right,
		}
		return node, 1 + l + r
	}
	// error case, bitString[0] not '0' or '1'
	return nil, 1
}
