package tree

// Invert changes the "sense" of a BST tree:
// from least-to-most value, inverted as most-to-least
func (p *Node) Invert() {
	if p == nil {
		return
	}
	p.Left.Invert()
	p.Right.Invert()
	p.Left, p.Right = p.Right, p.Left
}
