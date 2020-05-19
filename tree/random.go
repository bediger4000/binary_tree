package tree

import (
	"math/rand"
	"os"
	"time"
)

// RandomValueTree constructs a binary search tree with
// random-number-valued nodes. This is just a support function
// for seeing operations on whole trees (invert, max depth, etc) work.
func RandomValueTree(max, nodeCount int, setSeed bool) (root *Node) {

	if nodeCount > max {
		return nil
	}

	if setSeed {
		rand.Seed(time.Now().UnixNano() + int64(os.Getpid()))
	}

	found := make(map[int]bool)

	for i := 0; i < nodeCount; i++ {

		x := rand.Intn(max)
		for found[x] {
			x = rand.Intn(max)
		}

		found[x] = true
		root = Insert(root, x)
	}

	return root
}
