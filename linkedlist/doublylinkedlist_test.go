package linkedlist_test

import (
	"testing"

	"github.com/matlagp/go-algorithms-and-data-structures/linkedlist"
	"pgregory.net/rapid"
)

func TestAppend(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		example := rapid.SliceOf(rapid.Int()).Draw(t, "example")
		l := linkedlist.NewDoublyLinkedList[int]()
		for _, e := range example {
			l.Append(e)
		}

		if int(l.GetLength())!= len(example) {
			t.Fatalf("Expected length of %d but got %d", len(example), l.GetLength())
		}

		i := 0
		for e := range l.Each() {
			if example[i] != e {
				t.Fatalf("Expected %d at position %d, but got %d", example[i], i, e)
			}
			i++
		}
	})
}

func TestPrepend(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		example := rapid.SliceOf(rapid.Int()).Draw(t, "example")
		l := linkedlist.NewDoublyLinkedList[int]()
		for _, e := range example {
			l.Prepend(e)
		}

		if int(l.GetLength())!= len(example) {
			t.Fatalf("Expected length of %d but got %d", len(example), l.GetLength())
		}

		i := 0
		for e := range l.Reverse() {
			if example[i] != e {
				t.Fatalf("Expected %d at position %d, but got %d", example[i], i, e)
			}
			i++
		}
	})
}

