package stack

import "github.com/mg-kaung/gds/list/singlyLinkedList"

type Stack[T comparable] interface {
	Push(value T)
	Pop() (value T, err error)
	Seek() (value T, err error)
	Size() int
	IsEmpty() bool
}

type stack[T comparable] struct {
	l *singlyLinkedList.SinglyLinkedList[T]
}

func New[T comparable]() Stack[T] {
	return &stack[T]{
		l: singlyLinkedList.New[T](),
	}
}

func (s stack[T]) Push(value T) {
	s.l.InsertFront(value)
}

func (s stack[T]) Pop() (T, error) {
	var zero T
	value, err := s.l.Front()
	if err != nil {
		return zero, err
	}

	err = s.l.RemoveFront()
	if err != nil {
		return zero, err
	}

	return value, nil
}

func (s stack[T]) Seek() (T, error) {
	return s.l.Front()
}

func (s stack[T]) Size() int {
	return s.l.Len()
}

func (s stack[T]) IsEmpty() bool {
	return s.l.Len() == 0
}
