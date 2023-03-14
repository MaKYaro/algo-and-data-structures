package stack

import (
	"fmt"
	"testing"
)

func TestStack_Empty(t *testing.T) {
	var tests = []struct {
		in    []any
		out   int
		empty bool
	}{
		{[]any{}, 0, true},
		{[]any{}, 4, true},
		{[]any{1, 't', "qwerty", true}, 5, true},
		{[]any{'w', 23432, 0, 0, struct{}{}}, 2, false},
		{[]any{'0', '1', 999, "qw", false}, 0, false},
	}
	for idx, test := range tests {
		name := fmt.Sprint(idx)
		t.Run(name, func(t *testing.T) {
			s := MakeStack()
			for _, elem := range test.in {
				s.Push(elem)
			}
			for i := 0; i < test.out; i++ {
				s.Pop()
			}
			got := s.Empty()
			want := test.empty
			if got != want {
				t.Errorf("got %v want %v", got, want)
			}
		})
	}
}

func TestStack_Pop(t *testing.T) {
	var tests = []struct {
		in      []any
		elemOut int
		err     []error
		out     []any
	}{
		{
			[]any{1, 2, 3},
			3,
			[]error{nil, nil},
			[]any{3, 2, 1},
		},
		{
			[]any{'a', 'b', 'c'},
			3,
			[]error{nil, nil},
			[]any{'c', 'b', 'a'},
		},
		{
			[]any{},
			0,
			nil,
			[]any{},
		},
		{
			[]any{},
			1,
			[]error{emptyStackError{}},
			[]any{nil},
		},
		{
			[]any{1, "wer", 45.466, 9},
			5,
			[]error{nil, nil, nil, nil, emptyStackError{}},
			[]any{9, 45.466, "wer", 1, nil},
		},
	}
	for idx, test := range tests {
		name := fmt.Sprint(idx)
		t.Run(name, func(t *testing.T) {
			s := MakeStack()
			for _, elem := range test.in {
				_ = s.Push(elem)
			}
			for i := 0; i < test.elemOut; i++ {
				got, err := s.Pop()
				want := test.out[i]
				if got != want && err != test.err[i] {
					t.Errorf("got %v, want %v", got, want)
				}
			}
		})
	}
}

func TestStack_Top(t *testing.T) {
	var tests = []struct {
		in  []any
		out int
		top any
		err error
	}{
		{[]any{1, 6, 24}, 2, 1, nil},
		{[]any{"qw", "agag", "qtt4", "3"}, 4, nil, emptyStackError{}},
		{[]any{}, 0, nil, emptyStackError{}},
		{[]any{2.34, 4.0, "afrf"}, 0, "afrf", nil},
	}
	for idx, test := range tests {
		name := fmt.Sprint(idx)
		t.Run(name, func(t *testing.T) {
			s := MakeStack()
			for _, elem := range test.in {
				s.Push(elem)
			}
			for i := 0; i < test.out; i++ {
				s.Pop()
			}
			want := test.top
			got, err := s.Top()
			if want != got && test.err != err {
				t.Errorf("got  %v, want %v", got, want)
			}
		})
	}
}
