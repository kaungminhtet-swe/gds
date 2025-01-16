package singlyLinkedList

import (
	"errors"
	"github.com/mg-kaung/gds/list"
	"iter"
)

type Element[T comparable] struct {
	value T
	next  *Element[T]
}

type SinglyLinkedList[T comparable] struct {
	head, tail *Element[T]
	len        int
}

func New[T comparable]() *SinglyLinkedList[T] {
	l := &SinglyLinkedList[T]{}
	l.Init()
	return l
}

func (l *SinglyLinkedList[T]) Init() {
	l.tail = nil
	l.head = nil
	l.len = 0
}

func (l *SinglyLinkedList[T]) Iterate() iter.Seq2[int, T] {
	index := -1
	return func(yield func(index int, value T) bool) {
		for l.head != nil {
			index++
			if !yield(index, l.head.value) {
				break
			}
			l.head = l.head.next
		}
	}
}

func (l *SinglyLinkedList[T]) InsertAll(values ...T) {
	for _, value := range values {
		e := &Element[T]{value: value}

		if l.isEmpty() {
			l.head = e
			l.tail = e
		} else {
			l.tail.next = e
			l.tail = e
		}

		l.len++
	}
}

func (l *SinglyLinkedList[T]) InsertList(lists *list.List[T]) {}

// insert value at front of list, increment len
func (l *SinglyLinkedList[T]) InsertFront(value T) {
	e := &Element[T]{value: value}

	if l.isEmpty() {
		l.head = e
		l.tail = e
	} else {
		e.next = l.head
		l.head = e
	}

	l.incLen()
}

// insert value at back of list, increment len
func (l *SinglyLinkedList[T]) InsertBack(value T) {
	e := &Element[T]{value: value}

	if l.isEmpty() {
		l.head = e
		l.tail = e
	} else {
		l.tail.next = e
		l.tail = e
	}

	l.incLen()
}

func (l *SinglyLinkedList[T]) InsertAtIndex(index int, value T) error {
	if index == 0 {
		l.InsertFront(value)
		return nil
	}

	if l.isEmpty() {
		return errors.New("empty list")
	}

	if index < 0 || index > l.len {
		return errors.New("index out of range")
	}

	var prev *Element[T]
	for head := l.head; index >= 0; index-- {
		if index == 0 {
			e := &Element[T]{value: value}
			e.next = head
			prev.next = e
			l.incLen()
			return nil
		}
		prev = head
		head = head.next
	}
	return nil
}

func (l *SinglyLinkedList[T]) RemoveFront() {
	//TODO implement me
	panic("implement me")
}

func (l *SinglyLinkedList[T]) RemoveBack() {
	//TODO implement me
	panic("implement me")
}

func (l *SinglyLinkedList[T]) RemoveAtIndex(index int) error {
	//TODO implement me
	panic("implement me")
}

// Front returns the first value of list l or default nil value of type T if the list is empty.
func (l *SinglyLinkedList[T]) Front() (T, error) {
	if l.isEmpty() {
		var t T
		return t, errors.New("empty list")
	}
	return l.head.value, nil
}

// Back returns the last value of list l or default nil value of type T if the list is empty.
func (l *SinglyLinkedList[T]) Back() (T, error) {
	if l.isEmpty() {
		var t T
		return t, errors.New("empty list")
	}
	return l.tail.value, nil
}

func (l *SinglyLinkedList[T]) GetIndex(index int) T {
	//TODO implement me
	panic("implement me")
}

func (l *SinglyLinkedList[T]) Len() int {
	return l.len
}

func (l *SinglyLinkedList[T]) isEmpty() bool {
	return l.head == nil && l.tail == nil && l.len == 0
}

func (l *SinglyLinkedList[T]) incLen() {
	l.len++
}

func (l *SinglyLinkedList[T]) decLen() {
	l.len--
}
