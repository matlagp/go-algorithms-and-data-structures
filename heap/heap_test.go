package heap

import (
	"math"
	"slices"
	"testing"

	"pgregory.net/rapid"
)

func TestIndex(t *testing.T) {
	_, err := Index(-1).Parent()
	if err == nil {
		t.Fatal("calling Parent on negative index should rise error")
	}

	rapid.Check(t, func(t *rapid.T) {
		example := Index(rapid.IntRange(1, (math.MaxInt-2)/2).Draw(t, "example"))
		p, err := example.Parent()

		if err != nil {
			t.Fatalf("calling Parent on %d should not rise error", example)
		}
		if p >= example {
			t.Fatal("parent should be smaller than self")
		}

		pl, plErr := p.Left()
		pr, prErr := p.Right()

		if plErr != nil || prErr != nil {
			t.Fatalf("calling Left and Right on Parent(%d) should be successful but got errors: %v, %v", example, plErr, prErr)
		}
		if pl != example && pr != example {
			t.Fatalf("Either Left(Parent(%d)) or Right(Parent(%d)) should equal %d", example, example, example)
		}
	})
}

func TestBuildHeap(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		example := rapid.SliceOf(rapid.Int()).Draw(t, "example")
		h, err := BuildHeap(example)
		if err != nil {
			t.Fatalf("error while building heap: %v", err)
		}
		if !h.isValidHeap() {
			t.Fatalf("invalid heap: %v", h)
		}
	})
}

func (h *Heap[int]) isValidHeap() bool {
	for i, elem := range h.each() {
		l, lErr := i.Left()
		r, rErr := i.Right()
		if lErr != nil || rErr != nil {
			return false
		}
		if l <= h.last && elem < h.content[l] {
			return false
		}
		if r <= h.last && elem < h.content[r] {
			return false
		}
	}

	return true
}

func TestInsert(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		initHeap := rapid.SliceOf(rapid.Int()).Draw(t, "initHeap")
		append := rapid.SliceOf(rapid.Int()).Draw(t, "append")

		h, err := BuildHeapInPlace(initHeap)
		if err != nil {
			t.Fatalf("error while building heap: %v", err)
		}
		if !h.isValidHeap() {
			t.Fatalf("invalid heap: %v", h)
		}

		for _, elem := range append {
			err = h.Insert(elem)
			if err != nil  {
				t.Fatalf("error while inserting: %v", err)
			}
			if !h.isValidHeap() {
				t.Fatalf("invalid heap: %v", h)
			}
		}
	})
}

func TestExtract(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		initHeap := rapid.SliceOf(rapid.Int()).Draw(t, "initHeap")
		expected := slices.Clone(initHeap)
		actual := make([]int, len(expected))

		h, err := BuildHeapInPlace(initHeap)
		if err != nil {
			t.Fatalf("error while building heap: %v", err)
		}
		if !h.isValidHeap() {
			t.Fatalf("invalid heap: %v", h)
		}

		for i := range initHeap {
			elem, err := h.Extract()
			actual[i] = elem
			if err != nil  {
				t.Fatalf("error while extracting: %v", err)
			}
		}

		_, err = h.Extract()
		if err == nil {
			t.Fatal("expected error when extracting from empty heap, but got none")
		}

		slices.Sort(expected)
		slices.Sort(actual)
		if !slices.Equal(actual, expected) {
			t.Fatal("extracting all elements from heap should return same elements as those inserted")
		}
	})
}
