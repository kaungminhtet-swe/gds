package tree

import (
	"iter"
)

type InOrder[T comparable] struct {
	root *Node[T]
}

func (i *InOrder[T]) Iterate() iter.Seq2[int, T] {
	index := 0
	return func(yield func(int, T) bool) {
		for value := range RecursiveInOrder(i.root) {
			if !yield(index, value) {
				break
			}
			index++
		}
	}
}

func RecursiveInOrder[T comparable](root *Node[T]) iter.Seq[T] {
	return func(yield func(value T) bool) {
		if root.Left() != nil {
			for x := range RecursiveInOrder(root.Left()) {
				yield(x)
			}
		}

		yield(root.Value())

		if root.Right() != nil {
			for x := range RecursiveInOrder(root.Right()) {
				yield(x)
			}
		}
	}
}
