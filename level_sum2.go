package main

import (
	"binary_tree/tree"
	"fmt"
	"os"
)

type levelData struct {
	data  int
	level int
}

func main() {
	root, err := tree.CreateNumericFromString(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "parsing %q: %v\n", os.Args[1], err)
		return
	}

	ch := make(chan *levelData, 0)

	go traverse(root, ch)

	maxLevel := -1
	var levelSums []int

	for ld := range ch {
		if ld.level > maxLevel {
			levelSums = append(levelSums, ld.data)
			maxLevel = ld.level
		} else {
			levelSums[ld.level] += ld.data
		}
	}

	minLevel := 0
	minSum := root.Data

	for level, sum := range levelSums {
		if sum < minSum {
			minSum = sum
			minLevel = level
		}
	}

	fmt.Printf("Level %d has minimum sum of %d\n", minLevel, minSum)
}

func traverse(root *tree.NumericNode, ch chan *levelData) {
	findLevels(root, ch, 0)
	close(ch)
}

func findLevels(node *tree.NumericNode, ch chan *levelData, level int) {
	if node == nil {
		return
	}
	ch <- &levelData{
		data:  node.Data,
		level: level,
	}
	findLevels(node.Left, ch, level+1)
	findLevels(node.Right, ch, level+1)
}
