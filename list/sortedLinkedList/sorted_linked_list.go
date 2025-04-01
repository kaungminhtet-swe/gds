package sortedLinkedList

import (
	"errors"
	"iter"
)

type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64 | ~string
}

type Element[T Ordered] struct {
	value T
	next  *Element[T]
}

func (e *Element[T]) SetNext(next *Element[T]) {
	e.next = next
}

func (e *Element[T]) Next() *Element[T] {
	return e.next
}

type SortedLinkedList[T Ordered] struct {
	head, tail *Element[T]
	len        int
}

func (l *SortedLinkedList[T]) Init() {
	l.head, l.tail, l.len = nil, nil, 0
}

func (l *SortedLinkedList[T]) Iterate() iter.Seq2[int, T] {
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

func (l *SortedLinkedList[T]) InsertAll(value ...T) error {
	if len(value) == 0 {
		return errors.New("empty values")
	}

	for _, v := range value {
		l.Insert(v)
	}

	return nil
}

func (l *SortedLinkedList[T]) Insert(value T) {
	if l.IsEmpty() || value <= l.head.value {
		l.insertFront(value)
		return
	}

	if value >= l.tail.value {
		l.insertBack(value)
		return
	}

	head := l.head
	for head != nil {
		if head.Next().value > value {
			e := &Element[T]{value: value}
			e.SetNext(head.Next())
			head.SetNext(e)
			l.incLen()
			return
		}
		head = head.next
	}

}

func (l *SortedLinkedList[T]) Front() (T, error) {
	if l.IsEmpty() {
		var zero T
		return zero, errors.New("empty list")
	}
	return l.head.value, nil
}

func (l *SortedLinkedList[T]) Back() (T, error) {
	if l.IsEmpty() {
		var zero T
		return zero, errors.New("empty list")
	}
	return l.tail.value, nil
}

func (l *SortedLinkedList[T]) RemoveFront() error {
	if l.IsEmpty() {
		return errors.New("empty list")
	}
	l.head = l.head.next
	l.decLen()
	return nil
}

func (l *SortedLinkedList[T]) RemoveBack() error {
	if l.IsEmpty() {
		return errors.New("empty list")
	}

	var prev *Element[T]
	for head := l.head; head.next != nil; head = head.next {
		prev = head
	}

	prev.SetNext(nil)
	l.tail = prev
	l.decLen()
	return nil
}

func (l *SortedLinkedList[T]) Get(index int) (T, error) {
	if index < 0 || index >= l.len {
		var zero T
		return zero, errors.New("index out of bounds")
	}

	for _, v := range l.Iterate() {
		if index == 0 {
			return v, nil
		}
		index--
	}

	var zero T
	return zero, errors.New("something went wrong")
}

func (l *SortedLinkedList[T]) Len() int {
	return l.len
}

func (l *SortedLinkedList[T]) IsEmpty() bool {
	return l.head == nil && l.tail == nil && l.len == 0
}

func (l *SortedLinkedList[T]) insertFront(value T) {
	e := &Element[T]{value: value}

	if l.IsEmpty() {
		l.head = e
		l.tail = e
	} else {
		e.SetNext(l.head)
		l.head = e
	}

	l.incLen()
}

func (l *SortedLinkedList[T]) insertBack(value T) {
	e := &Element[T]{value: value}

	if l.IsEmpty() {
		l.setFront(e)
		l.setTail(e)
	} else {
		l.tail.SetNext(e)
		l.setTail(l.tail.next)
	}

	l.incLen()
}

func (l *SortedLinkedList[T]) setFront(head *Element[T]) {
	l.head = head
}

func (l *SortedLinkedList[T]) setTail(tail *Element[T]) {
	l.tail = tail
}

func (l *SortedLinkedList[T]) incLen() {
	l.len++
}

func (l *SortedLinkedList[T]) decLen() {
	l.len--
}

func New[T Ordered]() *SortedLinkedList[T] {
	l := &SortedLinkedList[T]{}
	l.Init()
	return l
}
