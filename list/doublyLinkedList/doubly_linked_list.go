package doublylinkedlist

import (
	"errors"
	"iter"
)

type Element[T comparable] struct {
	value T
	prev  *Element[T]
	next  *Element[T]
}

type DoublyLinkedList[T comparable] struct {
	head *Element[T]
	tail *Element[T]
	len  int
}

func New[T comparable]() *DoublyLinkedList[T] {
	l := &DoublyLinkedList[T]{}
	l.Init()
	return l
}

func (l *DoublyLinkedList[T]) Init() {
	l.head = nil
	l.tail = nil
	l.len = 0
}

func (l *DoublyLinkedList[T]) Iterate() iter.Seq2[int, T] {
	index := -1
	return func(yield func(index int, value T) bool) {
		head := l.head
		for head != nil {
			index++
			if !yield(index, head.value) {
				break
			}
			head = head.next
		}
	}
}

func (l *DoublyLinkedList[T]) InsertAll(values ...T) {
	for _, v := range values {
		l.InsertBack(v)
	}
}

func (l *DoublyLinkedList[T]) InsertFront(value T) {
	e := &Element[T]{value: value}

	if l.isEmpty() {
		l.head = e
		l.tail = e
	} else {
		e.next = l.head
		l.head.prev = e
		l.head = e
	}

	l.len++
}

func (l *DoublyLinkedList[T]) InsertAtIndex(index int, value T) error {
	if index == 0 {
		l.InsertFront(value)
		return nil
	}

	if l.isEmpty() {
		return errors.New("empty list")
	}

	if index < 0 || index >= l.Len() {
		return errors.New("index out of range")
	}

	current := l.head
	for ; index > 0; index-- {
		current = current.next
	}

	e := &Element[T]{value: value}
	current.prev.next = e
	current.prev = e
	e.next = current

	l.len++

	return nil
}

func (l *DoublyLinkedList[T]) InsertBack(value T) {
	e := &Element[T]{value: value}

	if l.isEmpty() {
		l.head = e
		l.tail = e
	} else {
		e.prev = l.tail
		l.tail.next = e
		l.tail = e
	}

	l.len++
}

func (l *DoublyLinkedList[T]) RemoveFront() error {
	if l.isEmpty() {
		return errors.New("empty list")
	}

	l.head = l.head.next
	l.head.prev = nil
	l.len--

	return nil
}

func (l *DoublyLinkedList[T]) RemoveBack() error {
	if l.isEmpty() {
		return errors.New("empty list")
	}

	l.tail = l.tail.prev
	l.tail.next = nil
	l.len--

	return nil
}

func (l *DoublyLinkedList[T]) RemoveAtIndex(index int) error {
	if l.isEmpty() {
		return errors.New("empty list")
	}

	if index < 0 || index >= l.len {
		return errors.New("index out of bounds")
	}

	if index == 0 {
		return l.RemoveFront()
	}

	if index == l.len-1 {
		return l.RemoveBack()
	}

	current := l.head
	for ; index > 0; index-- {
		current = current.next
	}

	current.prev.next = current.next
	current.next.prev = current.prev
	l.len--

	return nil
}

func (l *DoublyLinkedList[T]) isEmpty() bool {
	return l.head == nil && l.tail == nil && l.len == 0
}

func (l *DoublyLinkedList[T]) Len() int {
	return l.len
}
