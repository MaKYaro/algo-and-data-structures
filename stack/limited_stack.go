package stack

import "fmt"

// LimitedStack contains defined number of elements
type LimitedStack struct {
	cap uint
	arr []any
}

// fullStackError occurs when user tries to add cap + 1 element to LimitedStack
type fullStackError struct {
	cap uint
}

func (e fullStackError) Error() string {
	return fmt.Sprintf("stack size is %v", e.cap)
}

// MakeLimitedStack creates new LimitedStack entity
func MakeLimitedStack(cap uint) LimitedStack {
	return LimitedStack{cap, make([]any, cap)}
}

// Empty returns true if LimitedStack is empty
// and false if it isn't
func (s *LimitedStack) Empty() bool {
	return len(s.arr) == 0
}

// Pop takes the last element from the LimitedStack and returns it
// returns emptyStackError if LimitedStack is empty
func (s *LimitedStack) Pop() (any, error) {
	if !s.Empty() {
		last := s.arr[len(s.arr)-1]
		s.arr = s.arr[:len(s.arr)-1]
		return last, nil
	}
	return nil, emptyStackError{}
}

// Push adds new element to the LimitedStack
// if LimitedStack is full returns fullStackError
func (s *LimitedStack) Push(elem any) error {
	if len(s.arr) == int(s.cap) {
		return fullStackError{cap: s.cap}
	}
	s.arr = append(s.arr, elem)
	return nil
}

// Top returns the last element from the LimitedStack
// doesn't take the element
func (s *LimitedStack) Top() (any, error) {
	if !s.Empty() {
		return s.arr[len(s.arr)-1], nil
	}
	return nil, emptyStackError{}
}
