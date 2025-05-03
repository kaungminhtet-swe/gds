package binarytree

import (
	"cmp"
	"iter"

	"github.com/kmh-io/gds/tree"
)

func leftRotate[T cmp.Ordered](n *tree.Node[T]) *tree.Node[T] {
	return n
}

type BalancedBinaryTree[T cmp.Ordered] struct {
	root *tree.Node[T]
}

func (b *BalancedBinaryTree[T]) Insert(val T) {
	insert(b.root, val)
}

func insert[T cmp.Ordered](node *tree.Node[T], val T) {
	if node.Value() > val {
		if node.Left() == nil {
			node.SetLeft(val)
		} else {
			insert(node.Left(), val)
		}
	} else {
		if node.Right() == nil {
			node.SetRight(val)
		} else {
			insert(node.Right(), val)
		}
	}
}

func (b *BalancedBinaryTree[T]) InOrder() iter.Seq2[int, T] {
	index := -1
	return func(yield func(index int, value T) bool) {
		for v := range tree.RecursiveInOrder(b.root) {
			index++
			if !yield(index, v) {
				break
			}
		}
	}
}

func New[T cmp.Ordered](value T) *BalancedBinaryTree[T] {
	root := tree.NewNode(value)
	return &BalancedBinaryTree[T]{root: root}
}
