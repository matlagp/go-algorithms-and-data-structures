package linkedlist

import "iter"

type DoubleNode[T comparable] struct {
	Value T
	next  *DoubleNode[T]
	prev  *DoubleNode[T]
}

type DoublyLinkedList[T comparable] struct {
	head   *DoubleNode[T]
	last   *DoubleNode[T]
	length uint
}

func NewDoublyLinkedList[T comparable]() *DoublyLinkedList[T] {
	return &DoublyLinkedList[T]{}
}

func (l *DoublyLinkedList[T]) GetLength() uint {
	return l.length
}

func (l *DoublyLinkedList[T]) Append(val T) {
	n := &DoubleNode[T]{Value: val}

	if l.last == nil {
		l.head = n
		l.last = n
	} else {
		l.last.next = n
		l.last = n
		n.prev = l.last
	}

	l.length++
}

func (l *DoublyLinkedList[T]) Prepend(val T) {
	n := &DoubleNode[T]{Value: val}

	if l.head == nil {
		l.head = n
		l.last = n
	} else {
		l.head.prev = n
		l.head = n
		n.next = l.head
	}

	l.length++
}


func (l *DoublyLinkedList[T]) Each() iter.Seq[T] {
	return func(yield func(T) bool) {
		for n := l.head; n != nil; n = n.next {
			if !yield(n.Value) {
				return
			}
		}
	}
}

func (l *DoublyLinkedList[T]) Reverse() iter.Seq[T] {
	return func(yield func(T) bool) {
		for n := l.last; n != nil; n = n.prev {
			if !yield(n.Value) {
				return
			}
		}
	}
}

