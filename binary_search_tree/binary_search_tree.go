package binary_search_tree

import (
	"fmt"
	"github.com/MaKYaro/algo-and-data-structures/stack"
	"golang.org/x/exp/constraints"
)

type node[T constraints.Ordered] struct {
	key   T
	p     *node[T]
	left  *node[T]
	right *node[T]
}

type BinarySearchTree[T constraints.Ordered] struct {
	root *node[T]
}

// MakeBinarySearchTree function makes new binary search tree
// It returns pointer
func MakeBinarySearchTree[T constraints.Ordered]() *BinarySearchTree[T] {
	return &BinarySearchTree[T]{nil}
}

// InOrderedTreeWalk is a method used to traverse a binary search tree in a specific order.
// The traversal starts from the leftmost node of the tree, then visits the node itself, and finally proceeds to the right subtree.
func (t *BinarySearchTree[T]) InOrderedTreeWalk() []T {
	current := t.root
	if current == nil {
		return make([]T, 0)
	}
	result := make([]T, 0)
	s := stack.MakeStack()
	for true {
		if current != nil {
			err := s.Push(current)
			if err != nil {
				fmt.Println(err)
			}
			current = current.left
		} else {

			if s.Empty() {
				break
			}
			tmp, err := s.Pop()
			if err != nil {
				fmt.Println(err)
			}
			current = tmp.(*node[T])
			result = append(result, current.key)
			current = current.right
		}
	}
	return result
}

// Insert is method in a binary search tree is used to add a new node to the tree
// while maintaining the binary search tree property
func (t *BinarySearchTree[T]) Insert(key T) {
	var previous, current *node[T]
	previous = nil
	current = t.root
	for current != nil {
		previous = current
		if key < current.key {
			current = current.left
		} else {
			current = current.right
		}
	}
	switch {
	case previous == nil:
		t.root = &node[T]{key, nil, nil, nil}
	case key > previous.key:
		previous.right = &node[T]{key, previous, nil, nil}
	default:
		previous.left = &node[T]{key, previous, nil, nil}
	}
}
