package queue

import (
	"errors"

	"github.com/kmh-io/gds/list/singlyLinkedList"
)

type Queue[T comparable] interface {
	Enqueue(value T)
	Dequeue() (T, error)
	Peek() (T, error)
	Size() int
	IsEmpty() bool
}

type queue[T comparable] struct {
	l *singlyLinkedList.SinglyLinkedList[T]
}

func (q *queue[T]) Enqueue(value T) {
	q.l.InsertBack(value)
}

func (q *queue[T]) Dequeue() (T, error) {
	var zero T
	if q.IsEmpty() {
		return zero, errors.New("empty queue")
	}

	value, err := q.Peek()
	if err != nil {
		return zero, err
	}

	err = q.l.RemoveFront()
	if err != nil {
		return zero, err
	}

	return value, nil
}

func (q *queue[T]) Peek() (T, error) {
	if q.IsEmpty() {
		var zero T
		return zero, errors.New("empty queue")
	}
	return q.l.Front()
}

func (q *queue[T]) Size() int {
	return q.l.Len()
}

func (q *queue[T]) IsEmpty() bool {
	return q.l.Len() == 0
}

func New[T comparable]() Queue[T] {
	return &queue[T]{
		l: singlyLinkedList.New[T](),
	}
}
