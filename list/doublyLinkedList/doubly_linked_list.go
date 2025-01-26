package doublylinkedlist

import "iter"

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
		e := &Element[T]{value: v}

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

func (l *DoublyLinkedList[T]) isEmpty() bool {
	return l.head == nil && l.tail == nil && l.len == 0
}

func (l *DoublyLinkedList[T]) Len() int {
	return l.len
}
