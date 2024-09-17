package heap

import (
	"math"
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
