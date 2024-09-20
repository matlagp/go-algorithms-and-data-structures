package bst

import (
	"cmp"
	"iter"
)

type EmptyTreeError struct{}

func (e EmptyTreeError) Error() string {
	return "empty tree"
}

type node[T cmp.Ordered] struct {
	val         T
	parent, left, right *node[T]
}

func (n *node[T]) walk(yield func(T) bool) {
	if n.left != nil {
		n.left.walk(yield)
	}
	if !yield(n.val) {
		return
	}
	if n.right != nil {
		n.right.walk(yield)
	}
}

func (n *node[T]) search(val T) bool {
	if val == n.val {
		return true
	} else if val < n.val && n.left != nil {
		return n.left.search(val)
	} else if val > n.val && n.right != nil {
		return n.right.search(val)
	}

	return false
}

func (n *node[T]) min() T {
	if n.left != nil {
		return n.left.min()
	}

	return n.val
}

func (n *node[T]) max() T {
	if n.right != nil {
		return n.right.max()
	}

	return n.val
}

type Tree[T cmp.Ordered] struct {
	root *node[T]
}

func NewTree[T cmp.Ordered]() *Tree[T] {
	return &Tree[T]{}
}

func (t *Tree[T]) Has(val T) bool {
	if t.root == nil {
		return false
	}
	return t.root.search(val)
}
func (t *Tree[T]) Min() (T, error) {
	var ret T
	if t.root == nil {
		return ret, EmptyTreeError{}
	}
	return t.root.min(), nil
}
func (t *Tree[T]) Max() (T, error) {
	var ret T
	if t.root == nil {
		return ret, EmptyTreeError{}
	}
	return t.root.max(), nil
}

func (t *Tree[T]) Each() iter.Seq[T] {
	return func(yield func(T) bool) {
		if t.root != nil {
			t.root.walk(yield)
		}
	}
}

func (t *Tree[T]) Insert(val T) {
	n := node[T]{ val: val }
	if t.root == nil {
		t.root = &n
	} else {
		var prev, curr *node[T]
		curr = t.root
		for curr != nil {
			prev = curr
			if val < curr.val {
				curr = curr.left
			} else {
				curr = curr.right
			}
		}
		n.parent = prev
		if val < prev.val {
			prev.left = &n
		} else {
			prev.right = &n
		}
	}
}
