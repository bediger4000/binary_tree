package main

import (
	"binary_tree/tree"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	// Read value to search for
	targetNodeValue, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Node of interest has value %d\n", targetNodeValue)

	// construct BST from all the remaining command line values
	root := tree.CreateNumeric(os.Args[2:])

	n := tree.Find(root, targetNodeValue)
	if n == nil {
		fmt.Printf("Did not find %d in tree\n", targetNodeValue)
		return
	}
	fmt.Printf("Found node %v\n", n)
}
