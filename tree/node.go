package tree

type Node[T comparable] struct {
	value T
	left  *Node[T]
	right *Node[T]
}

func NewNode[T comparable](value T) *Node[T] {
	return &Node[T]{
		value: value,
	}
}

func (n *Node[T]) Value() T {
	return n.value
}

func (n *Node[T]) Left() *Node[T] {
	return n.left
}

func (n *Node[T]) Right() *Node[T] {
	return n.right
}

func (n *Node[T]) SetValue(value T) {
	n.value = value
}

func (n *Node[T]) SetLeft(value T) *Node[T] {
	left := NewNode(value)
	n.left = left
	return left
}

func (n *Node[T]) SetRight(value T) *Node[T] {
	right := NewNode(value)
	n.right = right
	return right
}
