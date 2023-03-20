package stack

import (
	"fmt"
	"testing"
)

func TestLimitedStack_Empty(t *testing.T) {
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
		name := fmt.Sprintf("case %v", idx)
		t.Run(name, func(t *testing.T) {
			s := MakeLimitedStack(4)
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

func TestLimitedStack_Pop(t *testing.T) {
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
			s := MakeLimitedStack(4)
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

func TestLimitedStack_Push(t *testing.T) {
	var tests = []struct {
		cap int
		in  []any
		err []error
	}{
		{0, []any{}, nil},
		{0, []any{1}, []error{fullStackError{0}}},
		{0, []any{1, 2, 3}, []error{fullStackError{0}, fullStackError{0}, fullStackError{0}}},
		{1, []any{true, 1, 3}, []error{nil, fullStackError{1}, fullStackError{1}}},
		{3, []any{1, 3.14, true, 6, 'q'}, []error{nil, nil, nil, fullStackError{3}, fullStackError{3}}},
	}
	for idx, test := range tests {
		name := fmt.Sprintf("case %v", idx)
		t.Run(name, func(t *testing.T) {
			s := MakeLimitedStack(test.cap)
			for i, elem := range test.in {
				gotErr := s.Push(elem)
				wantErr := test.err[i]
				if gotErr != wantErr {
					t.Errorf("got error %e want %e", gotErr, wantErr)
				}
			}
		})
	}
}

func TestLimitedStack_Top(t *testing.T) {
	var tests = []struct {
		cap int
		in  []any
		out int
		top any
		err error
	}{
		{0, []any{}, 0, nil, emptyStackError{}},
		{0, []any{1}, 0, nil, emptyStackError{}},
		{0, []any{1, 'q', true}, 1, nil, emptyStackError{}},
		{3, []any{'t', 0}, 0, 0, nil},
		{3, []any{'t', 0}, 1, 't', nil},
	}
	for idx, test := range tests {
		name := fmt.Sprintf("case %v", idx)
		t.Run(name, func(t *testing.T) {
			s := MakeLimitedStack(test.cap)
			for _, elem := range test.in {
				s.Push(elem)
			}
			for i := 0; i < test.out; i++ {
				s.Pop()
			}
			got, gotErr := s.Top()
			want := test.top
			if got != want {
				t.Errorf("got %v want %v", got, want)
			}
			wantErr := test.err
			if gotErr != wantErr {
				t.Errorf("got error %e want error %e", gotErr, wantErr)
			}
		})
	}
}
