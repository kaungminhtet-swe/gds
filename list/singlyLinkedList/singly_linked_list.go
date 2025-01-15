package singlyLinkedList

import (
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

func (l *SinglyLinkedList[T]) Iterate() iter.Seq2[int, T] {
	//TODO implement me
	panic("implement me")
}

func (l *SinglyLinkedList[T]) InsertAll(value T) {
	//TODO implement me
	panic("implement me")
}

func (l *SinglyLinkedList[T]) InsertList(list *list.List[T]) {
	//TODO implement me
	panic("implement me")
}

func (l *SinglyLinkedList[T]) InsertFront(value T) {
	//TODO implement me
	panic("implement me")
}

func (l *SinglyLinkedList[T]) InsertBack(value T) {
	//TODO implement me
	panic("implement me")
}

func (l *SinglyLinkedList[T]) InsertIndex(index int, value T) {
	//TODO implement me
	panic("implement me")
}

func (l *SinglyLinkedList[T]) RemoveFront() {
	//TODO implement me
	panic("implement me")
}

func (l *SinglyLinkedList[T]) RemoveBack() {
	//TODO implement me
	panic("implement me")
}

func (l *SinglyLinkedList[T]) RemoveIndex(index int) {
	//TODO implement me
	panic("implement me")
}

func (l *SinglyLinkedList[T]) Front() T {
	//TODO implement me
	panic("implement me")
}

func (l *SinglyLinkedList[T]) Back() T {
	//TODO implement me
	panic("implement me")
}

func (l *SinglyLinkedList[T]) GetIndex(index int) T {
	//TODO implement me
	panic("implement me")
}

func (l *SinglyLinkedList[T]) Len() int {
	//TODO implement me
	panic("implement me")
}

func (l *SinglyLinkedList[T]) Init() {
	l.tail = nil
	l.head = nil
	l.len = 0
}
