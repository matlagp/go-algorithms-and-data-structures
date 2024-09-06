package stack_test

import (
	"fmt"
	"testing"

	"github.com/matlagp/go-algorithms-and-data-structures/stack"
)

func TestPopEmptyStackShouldError(t *testing.T) {
	s := stack.New[int]()

	_, err := s.Pop()
	if err == nil {
		t.Fatal("Popping an element from an empty stack should return an error, but none was returned")
	}
}

func TestPushThenPop(t *testing.T) {
	type scenarioType struct {
		push []int
		popNum int
		expected []int
	}

	scenarios := []scenarioType{
		{ push: []int{}, popNum: 0, expected: []int{}},
		{ push: []int{1}, popNum: 1, expected: []int{}},
		{ push: []int{1, 2, 3}, popNum: 1, expected: []int{2, 1}},
		{ push: []int{4, 3, 2, 1}, popNum: 2, expected: []int{3, 4}},
	}

	for _, sc := range scenarios {
		t.Run(fmt.Sprintf("Pop %d times from %v", sc.popNum, sc.push), func(t *testing.T) {
			t.Parallel()

			s := stack.New[int]()

			for _, elem := range sc.push {
				s.Push(elem)
			}

			for range sc.popNum {
				s.Pop()
			}

			for _, elem := range sc.expected {
				actual, err := s.Pop()
				if err != nil {
					t.Fatal(fmt.Errorf("Error while popping an element from a stack: %w", err))
				}

				if actual != elem {
					t.Fatalf("Error while poping an element from a stack, expected %d, but got %d", elem, actual)
				}
			}
		})
	}
}
