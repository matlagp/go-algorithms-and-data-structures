package sorting_test

import (
	"testing"

	"github.com/matlagp/go-algorithms-and-data-structures/sorting"
	"pgregory.net/rapid"
)

func sorted(collection []int) bool {
	i := 1
	for i < len(collection) {
		if collection[i-1] > collection[i] {
			return false
		}
		i++
	}
	return true
}

func perform(t *testing.T, algorithm func(input []int)) {
	rapid.Check(t, func(t *rapid.T) {
		example := rapid.SliceOf(rapid.Int()).Draw(t, "example")

		algorithm(example)

		if !sorted(example) {
			t.Fatalf("Not sorted: %v", example)
		}
	})
}

func TestInsertionSort(t *testing.T) {
	perform(t, sorting.InsertionSort)
}

func TestMinSearchSort(t *testing.T) {
	perform(t, sorting.MinSearchSort)
}

func TestMergeSort(t *testing.T) {
	perform(t, sorting.MergeSort)
}

func TestQuickSort(t *testing.T) {
	perform(t, sorting.QuickSort)
}
