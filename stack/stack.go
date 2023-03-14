package stack

type Stack []any

// emptyStackError occurs when user try to pop element from empty Stack
type emptyStackError struct {
}

func (e emptyStackError) Error() string {
	return "stack is empty"
}

// MakeStack creates new Stack entity
func MakeStack() Stack {
	s := make([]any, 0)
	return s
}

// Empty return true if Stack is empty and false if it isn't
func (s *Stack) Empty() bool {
	return len(*s) == 0
}

// Pop takes the last element from the Stack and return it
// return emptyStackError if Stack is empty
func (s *Stack) Pop() (any, error) {
	if !s.Empty() {
		last := (*s)[len(*s)-1]
		*s = (*s)[:len(*s)-1]
		return last, nil
	}
	return nil, emptyStackError{}
}

// Push adds new element on the top of the Stack
func (s *Stack) Push(elem any) error {
	*s = append(*s, elem)
	return nil
}

// Top returns the last element from the Stack
// but doesn't delete it
func (s *Stack) Top() (any, error) {
	if !s.Empty() {
		return (*s)[len(*s)-1], nil
	}
	return nil, emptyStackError{}
}
