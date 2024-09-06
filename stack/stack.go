package stack

import (
	"fmt"

	"github.com/matlagp/go-algorithms-and-data-structures/linkedlist"
)

type Stack[T comparable] struct {
	list *linkedlist.LinkedList[T]
}

func New[T comparable]() *Stack[T] {
	return &Stack[T]{
		list: linkedlist.New[T](),
	}
}
func (s *Stack[T]) Push(val T) {
	s.list.Prepend(val)
}

func (s *Stack[T]) Pop() (T, error) {
	ret, err := s.list.PopHead()
	if err != nil {
		return ret, fmt.Errorf("failed to pop an element from stack: %w", err)
	}

	return ret, nil
}

func (s *Stack[T]) Peek() (T, error) {
	ret, err := s.list.PeekHead()
	if err != nil {
		return ret, fmt.Errorf("failed to peek an element on stack: %w", err)
	}

	return ret, nil
}
