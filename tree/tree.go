package tree

import "iter"

type Traversal[T comparable] func(*Node[T]) iter.Seq[T]

type Tree[T comparable] struct {
	root        *Node[T]
	traversalFn Traversal[T]
}

func (t *Tree[T]) SetTraverser(fn Traversal[T]) {
	t.traversalFn = fn
}

func (t *Tree[T]) Iterate() iter.Seq2[int, T] {
	index := 0
	return func(yield func(int, T) bool) {
		for value := range t.traversalFn(t.root) {
			if !yield(index, value) {
				break
			}
			index++
		}
	}
}

func New[T comparable](root *Node[T], fn ...Traversal[T]) *Tree[T] {
	t := &Tree[T]{root: root}
	if len(fn) > 0 {
		t.SetTraverser(fn[0])
	} else {
		t.SetTraverser(RecursiveInOrder)
	}

	return t
}
