package sorting

import (
	"cmp"

	"github.com/matlagp/go-algorithms-and-data-structures/heap"
)

func swap[T any](a *T, b *T) {
	tmp := *a
	*a = *b
	*b = tmp
}

func InsertionSort[T cmp.Ordered](input []T) {
	i := 1
	for i < len(input) {
		j := i
		for j > 0 {
			if input[j-1] > input[j] {
				swap(&input[j-1], &input[j])
			}
			j--
		}
		i++
	}
}

func MinSearchSort[T cmp.Ordered](input []T) {
	i := 0
	for i < len(input) {
		minIdx, j := i, i+1
		for j < len(input) {
			if input[j] < input[minIdx] {
				minIdx = j
			}
			j++
		}

		if minIdx != i {
			swap(&input[minIdx], &input[i])
		}

		i++
	}
}

func MergeSort[T cmp.Ordered](input []T) {
	mergeSort(input, make([]T, len(input)), 0, len(input))
}

func mergeSort[T cmp.Ordered](input []T, work []T, i, j int) {
	if j - i <= 1 {
		return
	}

	k := (i + j) / 2
	mergeSort(input, work, i, k)
	mergeSort(input, work, k, j)

	merge(input, work, i, j, k)
}

func merge[T cmp.Ordered](input []T, work []T, i, j, k int) {
	m := i
	for m < j {
		work[m] = input[m]
		m++
	}

	x, y, z := i, k, i

	for x < k && y < j {
		if work[x] < work[y] {
			input[z] = work[x]
			x++
		} else {
			input[z] = work[y]
			y++
		}
		z++
	}

	for x < k {
		input[z] = work[x]
		x++
		z++
	}

	for y < j {
		input[z] = work[y]
		y++
		z++
	}
}

func QuickSort[T cmp.Ordered](input []T) {
	quickSort(input, 0, len(input) - 1)
}

func quickSort[T cmp.Ordered](input []T, l, r int) {
	if r - l < 1 {
		return
	}

	m := partition(input, l, r)
	quickSort(input, l, m - 1)
	quickSort(input, m + 1, r)
}

func partition[T cmp.Ordered](input []T, l, r int) int {
	pivot := input[r]
	i, j := l, l

	for j < r {
		if input[j] < pivot {
			swap(&input[j], &input[i])
			i++
		}
		j++
	}

	swap(&input[i], &input[r])
	return i
}

func HeapSort[T cmp.Ordered](input []T) {
	h, err := heap.BuildHeapInPlace(input)
	if err != nil {
		return
	}
	h.DestructiveSort()
}
