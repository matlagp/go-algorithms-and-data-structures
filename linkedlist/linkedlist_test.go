package linkedlist_test

import (
	"fmt"
	"testing"

	"github.com/matlagp/go-algorithms-and-data-structures/linkedlist"
)

func TestNewListHasZeroLength(t *testing.T) {
	var expected uint = 0
	l := linkedlist.New[int]()

	actual := l.GetLength()
	if actual != expected {
		t.Fatalf("New list should have a length of %d, got %d instead", expected, actual)
	}
}

func TestNewListIsEmpty(t *testing.T) {
	expected := true
	l := linkedlist.New[int]()

	actual := l.Empty()
	if actual != expected {
		t.Fatal("New list should be empty")
	}
}

func TestAppendIncreasesLength(t *testing.T) {
	var expected uint = 1
	l := linkedlist.New[int]()

	l.Append(1234)

	actual := l.GetLength()
	if actual != expected {
		t.Fatalf("List after appending should have a length of %d, got %d instead", expected, actual)
	}
}

func TestAppendTwiceInsertsElementsInOrder(t *testing.T) {
	var expectedLength uint = 2
	firstVal, secondVal := 1234, 4321
	l := linkedlist.New[int]()

	l.Append(firstVal)
	l.Append(secondVal)

	actualLength := l.GetLength()
	if expectedLength != actualLength {
		t.Fatalf("List after appending should have a length of %d, got %d instead", expectedLength, actualLength)
	}

	n, err := l.PopHead()
	if err != nil {
		t.Fatal(fmt.Errorf("Failed to dereference list the first time: %w", err))
	}
	if n != firstVal {
		t.Fatalf("Expected first value to be %d, got %d instead", firstVal, n)
	}

	n, err = l.PopHead()
	if err != nil {
		t.Fatal(fmt.Errorf("Failed to dereference list the second time: %w", err))
	}
	if n != secondVal {
		t.Fatalf("Expected second value to be %d, got %d instead", secondVal, n)
	}
}

func TestPrependTwiceInsertsElementsInReverse(t *testing.T) {
	var expectedLength uint = 2
	firstVal, secondVal := 1234, 4321
	l := linkedlist.New[int]()

	l.Prepend(firstVal)
	l.Prepend(secondVal)

	actualLength := l.GetLength()
	if expectedLength != actualLength {
		t.Fatalf("List after prepending should have a length of %d, got %d instead", expectedLength, actualLength)
	}

	n, err := l.PopHead()
	if err != nil {
		t.Fatal(fmt.Errorf("Failed to dereference list the first time: %w", err))
	}
	if n != secondVal {
		t.Fatalf("Expected first value to be %d, got %d instead", secondVal, n)
	}

	n, err = l.PopHead()
	if err != nil {
		t.Fatal(fmt.Errorf("Failed to dereference list the second time: %w", err))
	}
	if n != firstVal {
		t.Fatalf("Expected second value to be %d, got %d instead", firstVal, n)
	}
}

func TestListHasElementAfterAppend(t *testing.T) {
	element := 1234
	l := linkedlist.New[int]()

	if l.Has(element) {
		t.Fatalf("New list should be empty, but it contains %d", element)
	}

	l.Append(element)

	if !l.Has(element) {
		t.Fatalf("List should have an element %d after appending, but it's missing", element)
	}
}

func TestRemove(t *testing.T) {
	scenarios := [][][]int{
		{{}, {1}, {}},
		{{1}, {1}, {}},
		{{1, 2, 3}, {1}, {2, 3}},
		{{1, 2, 1, 3}, {1}, {2, 3}},
		{{1, 2, 1, 3}, {1, 3}, {2}},
	}

	for _, s := range scenarios {
		t.Run(fmt.Sprintf("Delete %v from %v", s[1], s[0]), func(t *testing.T) {
			t.Parallel()

			l := linkedlist.New[int]()

			for _, elem := range s[0] {
				l.Append(elem)
			}

			for _, elem := range s[1] {
				l.Remove(elem)
			}

			if int(l.GetLength()) != len(s[2]) {
				t.Fatalf("Expected list to have a length of %d, but it has %d", len(s[2]), l.GetLength())
			}

			for _, elem := range s[1] {
				if l.Has(elem) {
					t.Fatalf("Expected %d to be removed from the list, but it's present", elem)
				}
			}
			for _, elem := range s[2] {
				if !l.Has(elem) {
					t.Fatalf("Expected %d to be in the list, but it's missing", elem)
				}
			}

		})
	}
}
