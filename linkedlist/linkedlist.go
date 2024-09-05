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
	l.length += 1
}

func (l *LinkedList[T]) GetHead() (*Node[T], error) {
	if l.head == nil {
		return nil, fmt.Errorf("tried to dereference an empty list")
	}

	return l.head, nil
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
