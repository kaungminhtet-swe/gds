package list

import (
	"github.com/mg-kaung/gds/iterator"
)

type List[T comparable] interface {
	iterator.Iterator[T]

	InsertAll(value ...T)
	InsertList(l *List[T])

	InsertFront(value T)
	InsertBack(value T)
	InsertIndex(index int, value T)

	RemoveFront()
	RemoveBack()
	RemoveIndex(index int)

	Front() (T, error)
	Back() (T, error)
	GetIndex(index int) T

	Len() int
	Init()
}
