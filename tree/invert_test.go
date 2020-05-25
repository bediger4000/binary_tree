package tree

import "testing"

func TestNumericNode_Invert(t *testing.T) {
	var tree1 *NumericNode = CreateNumeric([]string{"11", "2", "6", "1", "3", "0", "-1", "10", "100"})
	var tree2 *NumericNode = CreateNumeric([]string{"11", "2", "6", "1", "3", "0", "-1", "10", "100"})

	tree1.Invert()

	if Equals(tree1, tree2) {
		t.Errorf("inverted and original equal")
	}

	tree1.Invert()
	if !Equals(tree1, tree2) {
		t.Errorf("double inverted and original no equal")
	}

	tree1 = CreateNumeric([]string{"1", "0", "2"})
	tree2 = &NumericNode{Data: 1, Left: &NumericNode{Data: 2}, Right: &NumericNode{Data: 0}}

	tree1.Invert()
	if !Equals(tree1, tree2) {
		t.Errorf("inverted tree %v, not right %v", tree1, tree2)
	}
}
