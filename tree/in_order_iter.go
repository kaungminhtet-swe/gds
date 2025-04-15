package tree

import (
	"iter"
)

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
