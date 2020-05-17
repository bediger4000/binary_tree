package tree

// Invert changes the "sense" of a BST tree:
// from least-to-most value, inverted as most-to-least
func (p *Node) Invert() {
	if p == nil {
		return
	}
	p.left.Invert()
	p.right.Invert()
	p.left, p.right = p.right, p.left
}
