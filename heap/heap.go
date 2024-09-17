package heap

import (
	"cmp"
	"iter"
	"math"
	"slices"
)

type Index int
type IndexOverflowError struct {
	i Index
}
type EmptyHeapError struct {}

func (e *IndexOverflowError) Error() string { return "heap index overflow" }
func (e *EmptyHeapError) Error() string { return "empty heap" }

func (i Index) Parent() (Index, error) {
	if i <= 0 {
		return Index(-1), &IndexOverflowError{i}
	}
	return (i - 1) / 2, nil
}
func (i Index) Left() (Index, error) {
	if i < 0 || i > (math.MaxInt-1)/2 {
		return Index(-1), &IndexOverflowError{i}
	}
	return i*2 + 1, nil
}
func (i Index) Right() (Index, error) {
	if i < 0 || i > (math.MaxInt-2)/2 {
		return Index(-1), &IndexOverflowError{i}
	}
	return i*2 + 2, nil
}

type Heap[T cmp.Ordered] struct {
	content []T
	last    Index
}

func (h *Heap[T]) Size() int { return int(h.last + 1) }

func NewHeap[T cmp.Ordered]() *Heap[T] {
	return &Heap[T]{
		content: make([]T, 0),
	}
}

func BuildHeap[T cmp.Ordered](init []T) (*Heap[T], error) {
	return BuildHeapInPlace(slices.Clone(init))
}

func BuildHeapInPlace[T cmp.Ordered](init []T) (*Heap[T], error) {
	h := &Heap[T]{
		content: init,
		last:    Index(len(init) - 1),
	}

	if len(init) < 2 {
		return h, nil
	}

	i, err := h.last.Parent()
	if err != nil {
		return nil, err
	}

	for i >= 0 {
		err = h.maxHeapify(i)
		if err != nil {
			return nil, err
		}
		i--
	}

	return h, nil
}

func (h *Heap[T]) maxHeapify(i Index) error {
	l, lErr := i.Left()
	if lErr != nil {
		return lErr
	}
	r, rErr := i.Right()
	if rErr != nil {
		return rErr
	}
	largestIdx := i

	if l <= h.last && h.content[l] > h.content[largestIdx] {
		largestIdx = l
	}
	if r <= h.last && h.content[r] > h.content[largestIdx] {
		largestIdx = r
	}

	if largestIdx != i {
		h.swap(i, largestIdx)
		return h.maxHeapify(largestIdx)
	}

	return nil
}

func (h *Heap[T]) swap(i, j Index) {
	h.content[i], h.content[j] = h.content[j], h.content[i]
}

func (h *Heap[T]) each() iter.Seq2[Index, T] {
	return func(yield func(Index, T) bool) {
		for i := Index(0); i <= h.last; i++ {
			if !yield(i, h.content[i]) {
				return
			}
		}
	}
}

func (h *Heap[T]) DestructiveSort() error {
	for h.last > 0 {
		h.swap(h.last, 0)
		h.last--
		err := h.maxHeapify(0)
		if err != nil {
			return err
		}
	}
	return nil
}

func (h *Heap[T]) Peek() (T, error) {
	var ret T
	if h.last < 0 {
		return ret, &EmptyHeapError{}
	}
	return h.content[0], nil
}

func (h *Heap[T]) Extract() (T, error) {
	ret, err := h.Peek()
	if err != nil {
		return ret, err
	}
	h.content[0] = h.content[h.last]
	h.last--
	h.maxHeapify(0)
	return ret, nil
}

func (h *Heap[T]) Insert(val T) error {
	h.content = append(h.content, val)
	h.last++

	i := h.last
	for i > 0 {
		p, err := i.Parent()
		if err != nil {
			return err
		}
		if h.content[p] < h.content[i] {
			h.swap(i, p)
			i = p
		} else {
			break
		}
	}

	return nil
}
