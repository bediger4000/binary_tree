package main

import (
	"binary_tree/tree"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	graphViz := flag.Bool("g", false, "GraphViz dot output on stdout")
	n := flag.Int("n", 5, "number of nodes in random tree")
	nonrandvals := flag.Bool("r", false, "use sequential node values")
	flag.Parse()

	rand.Seed(time.Now().UnixNano() + int64(os.Getpid()))

	root := generateRandom(*n, *nonrandvals)

	if *graphViz {
		tree.Draw(root)
		return
	}

	tree.Print(root)
	fmt.Println()
}

func generateRandom(n int, nonrandvals bool) tree.Node {
	ary := make([]int, n)

	if nonrandvals {
		for i := 0; i < n; i++ {
			ary[i] = i
		}
	} else {
		for i := 0; i < n; i++ {
			ary[i] = int(rand.Intn(n))
		}
	}

	return something(ary)
}

func something(ary []int) *tree.NumericNode {
	if len(ary) == 0 {
		return nil
	}

	if len(ary) == 1 {
		return &tree.NumericNode{Data: ary[0]}
	}

	div := rand.Intn(len(ary))
	if div >= len(ary) {
		div--
	}

	return &tree.NumericNode{
		Data:  ary[div],
		Left:  something(ary[:div]),
		Right: something(ary[div+1:]),
	}
}
