package binary_search_tree

import (
	"fmt"
	"github.com/MaKYaro/algo-and-data-structures/stack"
	"golang.org/x/exp/constraints"
)

type NoSuchKey[T constraints.Ordered] struct {
	elem T
}

func (n NoSuchKey[T]) Error() string {
	return fmt.Sprintf("no such element: %v", n.elem)
}

type EmptyTree[T constraints.Ordered] struct {
}

func (n EmptyTree[T]) Error() string {
	return fmt.Sprintf("tree is empty")
}

type NoSuccessor[T constraints.Ordered] struct {
	elem T
}

func (n NoSuccessor[T]) Error() string {
	return fmt.Sprintf("there is no successor for %v", n.elem)
}

type NoPredecessor[T constraints.Ordered] struct {
	elem T
}

func (n NoPredecessor[T]) Error() string {
	return fmt.Sprintf("there is no predecessor fot %v", n.elem)
}

type Node[T constraints.Ordered] struct {
	key   T
	p     *Node[T]
	left  *Node[T]
	right *Node[T]
}

// Min returns link to Node with minimum key in the binary search subtree
// or nil if subtree is empty
func (n *Node[T]) Min() (*Node[T], error) {
	current := n
	switch current {
	case nil:
		return nil, EmptyTree[T]{}
	default:
		for current.left != nil {
			current = current.left
		}
		return current, nil
	}
}

// Max returns link to Node with maximum key in the binary search subtree
// or nil if subtree is empty
func (n *Node[T]) Max() (*Node[T], error) {
	current := n
	switch current {
	case nil:
		return nil, EmptyTree[T]{}
	default:
		for current.right != nil {
			current = current.right
		}
		return current, nil
	}
}

// Successor returns link to Node with the minimum key among all keys greater than the given one
// or nil if key is biggest in the tree
func (n *Node[T]) Successor() (*Node[T], error) {
	switch n.right {
	case nil:
		current := n
		for current.p != nil && current != current.p.left {
			current = current.p
		}
		if current.p == nil {
			return nil, NoSuccessor[T]{n.key}
		}
		return current.p, nil
	default:
		successor, err := n.right.Min()
		if err != nil {
			fmt.Println(err)
		}
		return successor, nil
	}
}

// Predecessor returns link to Node with the maximum key among all keys smaller than the given one
// or nil if key is smallest in the tree
func (n *Node[T]) Predecessor() (*Node[T], error) {
	switch n.left {
	case nil:
		current := n
		for current.p != nil && current != current.p.right {
			current = current.p
		}
		if current.p == nil {
			return nil, NoPredecessor[T]{n.key}
		}
		return current.p, nil
	default:
		predecessor, err := n.left.Max()
		if err != nil {
			fmt.Println(err)
		}
		return predecessor, nil
	}
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

// Search method returns link to Node with key and nil error
// or nil and noSuchElement if there is no element with key
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
	return nil, NoSuchKey[T]{key}
}

// Min returns link to Node with minimum key in the binary search tree
// or nil if tree is empty
func (t *BinarySearchTree[T]) Min() (*Node[T], error) {
	return t.root.Min()
}

// Max returns link to Node with maximum key in the binary search tree
// or nil if tree is empty
func (t *BinarySearchTree[T]) Max() (*Node[T], error) {
	return t.root.Max()
}
