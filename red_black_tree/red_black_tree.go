package red_black_tree

import (
	"fmt"
	bst "github.com/MaKYaro/algo-and-data-structures/binary_search_tree"
	"github.com/MaKYaro/algo-and-data-structures/stack"
	"golang.org/x/exp/constraints"
)

type RBNode[T constraints.Ordered] struct {
	Key   T
	P     *RBNode[T]
	Left  *RBNode[T]
	Right *RBNode[T]
	IsRed bool
}

type RedBlackTree[T constraints.Ordered] struct {
	root *RBNode[T]
	none *RBNode[T]
}

// MakeRedBlackTree function makes new RedBlackTree
// It returns pointer
func MakeRedBlackTree[T constraints.Ordered]() *RedBlackTree[T] {
	none := &RBNode[T]{}
	return &RedBlackTree[T]{none, none}
}

// Height returns the number of black nodes on left path from n to leaf
func (t *RedBlackTree[T]) Height() int {
	height := 0
	curr := t.root
	for curr != t.none {
		if !curr.IsRed {
			height++
		}
		curr = curr.Left
	}

	return height
}

// InOrderedTreeWalk is a method used to traverse a red-black tree in a specific order.
// The traversal starts from the leftmost node of the tree, then visits the node itself, and finally proceeds to the right subtree.
func (t *RedBlackTree[T]) InOrderedTreeWalk() []T {
	current := t.root
	if current == t.none {
		return make([]T, 0)
	}
	result := make([]T, 0)
	s := stack.MakeStack()
	for true {
		if current != t.none {
			err := s.Push(current)
			if err != nil {
				fmt.Println(err)
			}
			current = current.Left
		} else {

			if s.Empty() {
				break
			}
			tmp, err := s.Pop()
			if err != nil {
				fmt.Println(err)
			}
			current = tmp.(*RBNode[T])
			result = append(result, current.Key)
			current = current.Right
		}
	}
	return result
}

// Min returns link to RBNode with minimum key in the red-black subtree
// or nil if subtree is empty
func (t *RedBlackTree[T]) Min(n *RBNode[T]) (*RBNode[T], error) {
	current := n
	switch current {
	case t.none:
		return nil, bst.EmptyTree[T]{}
	default:
		for current.Left != t.none {
			current = current.Left
		}
		return current, nil
	}
}

// Max returns link to RBNode with maximum key in the red-black subtree
// or nil if subtree is empty
func (t *RedBlackTree[T]) Max(n *RBNode[T]) (*RBNode[T], error) {
	current := n
	switch current {
	case t.none:
		return nil, bst.EmptyTree[T]{}
	default:
		for current.Right != nil {
			current = current.Right
		}
		return current, nil
	}
}

// Successor returns link to RBNode with the minimum key among all keys greater than the given one
// or nil if key is biggest in the tree
func (t *RedBlackTree[T]) Successor(n *RBNode[T]) (*RBNode[T], error) {
	switch n.Right {
	case t.none:
		current := n
		for current.P != t.none && current != current.P.Left {
			current = current.P
		}
		if current.P == t.none {
			return nil, bst.NoSuccessor[T]{n.Key}
		}
		return current.P, nil
	default:
		curr := n.Right
		for curr.Left != t.none {
			curr = curr.Left
		}

		return curr, nil
	}
}

// Search method returns link to RBNode with key and nil error
// or nil and NoSuchKey error if there is no element with key
func (t *RedBlackTree[T]) Search(key T) (*RBNode[T], error) {
	current := t.root
	for current != t.none {
		switch {
		case current.Key == key:
			return current, nil
		case current.Key > key:
			current = current.Left
		case current.Key < key:
			current = current.Right
		}
	}
	return nil, bst.NoSuchKey[T]{key}
}

// LeftRotate performs a single counter-clockwise rotation of a node and its right child
// converting a right-leaning red-black tree into a left-leaning red-black tree
func (t *RedBlackTree[T]) leftRotate(n *RBNode[T]) {
	nRight := n.Right
	n.Right = nRight.Left
	if nRight.Left != t.none {
		nRight.Left.P = n
	}
	nRight.P = n.P
	switch {
	case n.P == t.none:
		t.root = nRight
	case n == n.P.Left:
		n.P.Left = nRight
	default:
		n.P.Right = nRight
	}
	nRight.Left = n
	n.P = nRight
}

// RightRotate performs a single counter-clockwise rotation of a node and its left child
// converting a left-leaning red-black tree into a right-leaning red-black tree
func (t *RedBlackTree[T]) rightRotate(n *RBNode[T]) {
	nLeft := n.Left
	n.Left = nLeft.Right
	if nLeft.Right != t.none {
		nLeft.Right.P = n
	}
	nLeft.P = n.P
	switch {
	case n.P == t.none:
		t.root = nLeft
	case n == n.P.Left:
		n.P.Left = nLeft
	default:
		n.P.Right = nLeft
	}
	nLeft.Right = n
	n.P = nLeft
}

// InsertFixUp restores red-black properties
func (t *RedBlackTree[T]) insertFixUp(n *RBNode[T]) {
	for n.P.IsRed {
		if n.P == n.P.P.Left {
			uncle := n.P.P.Right
			if uncle.IsRed {
				n.P.IsRed = false
				uncle.IsRed = false
				n.P.P.IsRed = true
				n = n.P.P
			} else if n == n.P.Right {
				n = n.P
				t.leftRotate(n)

			} else {
				n.P.IsRed = false
				n.P.P.IsRed = true
				t.rightRotate(n.P.P)
			}
		} else {
			uncle := n.P.P.Left
			if uncle.IsRed {
				n.P.IsRed = false
				uncle.IsRed = false
				n.P.P.IsRed = true
				n = n.P.P
			} else if n == n.P.Left {
				n = n.P
				t.rightRotate(n)
			} else {
				n.P.IsRed = false
				n.P.P.IsRed = true
				t.leftRotate(n.P.P)
			}
		}
	}
	t.root.IsRed = false
}

// Insert is method in a red-black tree is used to add a new node to the tree
// while maintaining the red-black properties
func (t *RedBlackTree[T]) Insert(key T) {
	n := &RBNode[T]{key, t.none, t.none, t.none, true}
	prev := t.none
	curr := t.root
	for curr != t.none {
		prev = curr
		if n.Key < curr.Key {
			curr = curr.Left
		} else {
			curr = curr.Right
		}
	}
	n.P = prev
	if prev == t.none {
		t.root = n
	} else if n.Key < prev.Key {
		prev.Left = n
	} else {
		prev.Right = n
	}
	n.Left = t.none
	n.Right = t.none
	t.insertFixUp(n)
}

// Transplant replaces u subtree with v subtree in red-black tree
func (t *RedBlackTree[T]) transplant(u, v *RBNode[T]) {
	if u.P == t.none {
		t.root = v
	} else if u == u.P.Left {
		u.P.Left = v
	} else {
		u.P.Right = v
	}
	v.P = u.P
}

// DeleteFixUp restores red-black properties
func (t *RedBlackTree[T]) deleteFixUp(x *RBNode[T]) {
	for x != t.root && x.IsRed == false {
		var w *RBNode[T]
		if x == x.P.Left {
			w = x.P.Right
			if w.IsRed == true {
				w.IsRed = false
				x.P.IsRed = true
				t.leftRotate(x.P)
				w = x.P.Right
			}
			if w.Left.IsRed == false && w.Right.IsRed == false {
				w.IsRed = true
				x = x.P
			} else if w.Right.IsRed == false {
				w.Left.IsRed = false
				w.IsRed = true
				t.rightRotate(w)
				w = x.P.Right
			} else {
				w.IsRed = x.P.IsRed
				x.P.IsRed = false
				w.Right.IsRed = false
				t.leftRotate(x.P)
				x = t.root
			}
		} else {
			w = x.P.Left
			if w.IsRed {
				w.IsRed = false
				x.P.IsRed = true
				t.rightRotate(x.P)
				w = x.P.Left
			}
			if !w.Right.IsRed && !w.Left.IsRed {
				w.IsRed = true
				x = x.P
			} else if !w.Left.IsRed {
				w.Right.IsRed = false
				w.IsRed = true
				t.leftRotate(w)
				w = x.P.Right
			} else {
				w.IsRed = x.P.IsRed
				x.P.IsRed = false
				w.Left.IsRed = false
				t.rightRotate(x.P)
				x = t.root
			}
		}
	}
	x.IsRed = false
}

// Delete is method in a red-black tree is used to delete node from the tree
// while maintaining the red-black properties
func (t *RedBlackTree[T]) Delete(z *RBNode[T]) {
	var x *RBNode[T]
	y := z
	yOriginalColor := y.IsRed
	if z.Left == t.none {
		x = z.Right
		t.transplant(z, z.Right)
	} else if z.Right == t.none {
		x = z.Left
		t.transplant(z, z.Left)
	} else {
		var err error
		y, err = t.Min(t.root)
		if err != nil {
			fmt.Println(err)
		}
		yOriginalColor = y.IsRed
		x = y.Right
		if y.P == z {
			x.P = y
		} else {
			t.transplant(y, y.Right)
			y.Right = z.Right
			y.Right.P = y
		}
		t.transplant(z, y)
		y.Left = z.Left
		y.Left.P = y
		y.IsRed = z.IsRed
	}
	if yOriginalColor == false {
		t.deleteFixUp(x)
	}
}
