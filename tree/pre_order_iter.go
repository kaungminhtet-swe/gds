package tree

import "iter"

func RecursivePreOrder[T comparable](root *Node[T]) iter.Seq[T] {
	return func(yield func(value T) bool) {
		yield(root.Value())

		if root.Left() != nil {
			for x := range RecursivePreOrder(root.Left()) {
				yield(x)
			}
		}

		if root.Right() != nil {
			for x := range RecursivePreOrder(root.Right()) {
				yield(x)
			}
		}
	}
}
