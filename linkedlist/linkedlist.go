package linkedlist

import (
	"fmt"
	"iter"
)

type Node[T comparable] struct {
	Value T
	next  *Node[T]
}

func (n *Node[T]) HasNext() bool {
	return n.next != nil
}

func (n *Node[T]) GetNext() (*Node[T], error) {
	if n.next == nil {
		return nil, fmt.Errorf("there is no next node")
	}

	return n.next, nil
}

type LinkedList[T comparable] struct {
	head   *Node[T]
	length uint
}

func New[T comparable]() *LinkedList[T] {
	return &LinkedList[T]{}
}

func (l *LinkedList[T]) GetLength() uint {
	return l.length
}

func (l *LinkedList[T]) Empty() bool {
	return l.length == 0
}

func (l *LinkedList[T]) Append(val T) {
	if l.head == nil {
		l.head = &Node[T]{Value: val}
	} else {
		curr := l.head
		for curr.HasNext() {
			curr = curr.next
		}

		curr.next = &Node[T]{Value: val}
	}
	l.length++
}

func (l *LinkedList[T]) Prepend(val T) {
	l.head = &Node[T]{
		Value: val,
		next:  l.head,
	}
	l.length++
}

func (l *LinkedList[T]) PeekHead() (T, error) {
	var ret T

	if l.head == nil {
		return ret, fmt.Errorf("tried to dereference an empty list")
	}

	return l.head.Value, nil
}

func (l *LinkedList[T]) PopHead() (T, error) {
	ret, err := l.PeekHead()
	if err != nil {
		return ret, err
	}

	l.head = l.head.next
	l.length--
	return ret, nil

}

func (l *LinkedList[T]) Each() iter.Seq[T] {
	return func(yield func(T) bool) {
		curr := l.head

		for curr != nil {
			if !yield(curr.Value) {
				return
			}

			curr = curr.next
		}
	}
}

func (l *LinkedList[T]) Has(val T) bool {
	for currVal := range l.Each() {
		if currVal == val {
			return true
		}
	}

	return false
}

func (l *LinkedList[T]) Remove(val T) {
	if l.head == nil {
		return
	}

	prev, curr := l.head, l.head
	for curr != nil {
		if curr.Value == val {
			if l.head == curr {
				l.head = curr.next
			}

			prev.next = curr.next
			l.length--
		}

		prev = curr
		curr = curr.next
	}
}
