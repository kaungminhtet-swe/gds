package binarytree

import (
	"cmp"
	"iter"
)

type BinaryTree[T cmp.Ordered] interface {
	Insert(value T)
	Delete(value T)
	Search(value T) bool
	FindMin() T
	FindMax() T
	InOrder() iter.Seq2[int, T]
	PreOrder() iter.Seq2[int, T]
	PostOrder() iter.Seq2[int, T]
	Height() T
	Size() T
	Clear()
}
