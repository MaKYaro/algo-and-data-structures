package queue

import (
	"fmt"
	"testing"
)

func TestQueue_Add(t *testing.T) {
	var tests = []struct {
		in  []any
		out []any
	}{
		{[]any{1, 2, 3}, []any{1, 2, 3}},
		{[]any{}, []any{}},
		{[]any{1}, []any{1}},
		{[]any{'1', 0, true, false, "qwer"}, []any{'1', 0, true, false, "qwer"}},
	}

	for idx, test := range tests {
		name := fmt.Sprintf("case %v", idx)
		t.Run(name, func(t *testing.T) {
			queue := MakeQueue()
			for _, elem := range test.in {
				queue.Add(elem)
			}
			got := queue.returnQueue()
			for i, want := range test.out {
				if want != got[i] {
					t.Errorf("got %v want %v", got[i], want)
				}
			}
		})
	}
}

func TestQueue_Take(t *testing.T) {
	var tests = []struct {
		in      []any
		elemOut int
		err     []error
		out     []any
	}{
		{[]any{}, 0, []error{}, []any{}},
		{[]any{}, 1, []error{emptyQueueError{}}, []any{nil}},
		{
			[]any{},
			4,
			[]error{emptyQueueError{}, emptyQueueError{}, emptyQueueError{}, emptyQueueError{}},
			[]any{nil, nil, nil, nil},
		},
		{[]any{9}, 1, []error{nil}, []any{9}},
		{[]any{true}, 3, []error{nil, emptyQueueError{}, emptyQueueError{}}, []any{true, nil, nil}},
		{
			[]any{1, 2, 'q', true, "pop"},
			4,
			[]error{nil, nil, nil, nil},
			[]any{1, 2, 'q', true},
		},
	}
	for idx, test := range tests {
		name := fmt.Sprintf("case %v", idx)
		t.Run(name, func(t *testing.T) {
			queue := MakeQueue()
			for _, elem := range test.in {
				queue.Add(elem)
			}
			var got, want any
			var gotErr, wantErr error
			for i := 0; i < test.elemOut; i++ {
				got, gotErr = queue.Take()
				want = test.out[i]
				wantErr = test.err[i]
				if got != want {
					t.Errorf("got %v want %v", got, want)
				}
				if gotErr != wantErr {
					t.Errorf("got error %v want error %v", gotErr, wantErr)
				}
			}
		})
	}
}
