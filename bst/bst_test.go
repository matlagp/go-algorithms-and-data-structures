package bst_test

import (
	"iter"
	"slices"
	"testing"

	"github.com/matlagp/go-algorithms-and-data-structures/bst"
	"pgregory.net/rapid"
)

func TestInsert(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		example := rapid.SliceOf(rapid.Int()).Draw(t, "example")
		tr := bst.NewTree[int]()

		for _, elem := range example {
			tr.Insert(elem)
		}

		if !isValidTree(example, tr) {
			t.Fatal("invalid tree")
		}
	})
}

func isValidTree(arr []int, t *bst.Tree[int]) bool {
	slices.Sort(arr)
	next, stop := iter.Pull(t.Each())
	defer stop()

	for _, arrElem := range arr {
		treeElem, ok := next()
		if !ok || treeElem != arrElem {
			return false
		}
	}

	_, ok := next()
	return !ok
}
