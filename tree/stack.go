package tree

// Stack constitutes a push-down stack of binary tree nodes
type Stack struct {
	array []Node
}

// Pop retrieves the last Node pushed on the Stack
func (nq *Stack) Pop() (tail Node) {
	l := len(nq.array) - 1
	if l < 0 {
		return
	}
	tail = nq.array[l]
	nq.array = nq.array[:l]
	return
}

// Dequeue retrieves the first Node pushed on the Stack
func (nq *Stack) Dequeue() (head Node) {
	if len(nq.array) == 0 {
		return nil
	}
	head = nq.array[0]
	nq.array = nq.array[1:]
	return
}

// Empty returns true if the stack contains no Node elements,
// false otherwise.
func (nq *Stack) Empty() bool {
	return len(nq.array) == 0
}

// Push enters a Node at the top of the stack
func (nq *Stack) Push(n Node) {
	if n.IsNil() {
		return
	}
	nq.array = append(nq.array, n)
}
