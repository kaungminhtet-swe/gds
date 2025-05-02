package list

import (
	"github.com/kmh-io/gds/iterator"
)

type List[T comparable] interface {
	iterator.Iterator[T]

	InsertAll(value ...T)

	InsertFront(value T)
	InsertBack(value T)
	InsertAt(index int, value T) error

	RemoveFront() error
	RemoveBack() error
	RemoveAt(index int) error

	Front() (T, error)
	Back() (T, error)
	Get(index int) (T, error)

	Len() int
	Init()
}
