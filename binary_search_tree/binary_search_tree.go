package binary_search_tree

import (
	"fmt"
	"github.com/MaKYaro/algo-and-data-structures/stack"
	"golang.org/x/exp/constraints"
)

type noSuchElement[T constraints.Ordered] struct {
	elem T
}

func (n noSuchElement[T]) Error() string {
	return fmt.Sprintf("no such element: %v", n.elem)
}

type Node[T constraints.Ordered] struct {
	key   T
	p     *Node[T]
	left  *Node[T]
	right *Node[T]
}

type BinarySearchTree[T constraints.Ordered] struct {
	root *Node[T]
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
			current = tmp.(*Node[T])
			result = append(result, current.key)
			current = current.right
		}
	}
	return result
}

// Insert is method in a binary search tree is used to add a new node to the tree
// while maintaining the binary search tree property
func (t *BinarySearchTree[T]) Insert(key T) {
	var previous, current *Node[T]
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
		t.root = &Node[T]{key, nil, nil, nil}
	case key > previous.key:
		previous.right = &Node[T]{key, previous, nil, nil}
	default:
		previous.left = &Node[T]{key, previous, nil, nil}
	}
}

func (t *BinarySearchTree[T]) Search(key T) (*Node[T], error) {
	current := t.root
	for current != nil {
		switch {
		case current.key == key:
			return current, nil
		case current.key > key:
			current = current.left
		case current.key < key:
			current = current.right
		}
	}
	return nil, noSuchElement[T]{key}
}
