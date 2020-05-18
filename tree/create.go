package tree

import (
	"fmt"
	"os"
	"strconv"
)

func CreateNumeric(numberRepr []string) (root *Node) {

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
