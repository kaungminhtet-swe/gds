package tree

import "iter"

func RecursivePostOrder[T comparable](root *Node[T]) iter.Seq[T] {
	return func(yield func(value T) bool) {
		if root.Left() != nil {
			for x := range RecursivePostOrder(root.Left()) {
				yield(x)
			}
		}

		if root.Right() != nil {
			for x := range RecursivePostOrder(root.Right()) {
				yield(x)
			}
		}

		yield(root.Value())
	}
}
