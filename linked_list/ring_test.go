package linked_list

import (
	"fmt"
	"testing"
)

func TestRingLL_InsertHead(t *testing.T) {
	var tests = []struct {
		in  []any
		out []any
	}{
		{[]any{}, []any{}},
		{[]any{2}, []any{2}},
		{[]any{1, 2, 3, 4, 5}, []any{5, 4, 3, 2, 1}},
		{[]any{'q', true, 0, 9, "are", 1.1}, []any{1.1, "are", 9, 0, true, 'q'}},
	}
	for idx, test := range tests {
		name := fmt.Sprintf("case %v", idx)
		t.Run(name, func(t *testing.T) {
			list := MakeRingLL()
			for _, elem := range test.in {
				list.InsertHead(elem)
			}
			var got any
			gotList := list.returnList()
			for i, want := range test.out {
				got = gotList[i]
				if got != want {
					t.Errorf("got %v want %v", got, want)
				}
			}
		})
	}
}

func TestRingLL_InsertTail(t *testing.T) {
	var tests = []struct {
		in  []any
		out []any
	}{
		{[]any{}, []any{}},
		{[]any{2}, []any{2}},
		{[]any{1, 2, 3, 4, 5}, []any{1, 2, 3, 4, 5}},
		{[]any{'q', true, 0, 9, "are", 1.1}, []any{'q', true, 0, 9, "are", 1.1}},
	}

	for idx, test := range tests {
		name := fmt.Sprintf("case %v", idx)
		t.Run(name, func(t *testing.T) {
			list := MakeRingLL()
			for _, elem := range test.in {
				list.InsertTail(elem)
			}
			var got any
			gotList := list.returnList()
			for i, want := range test.out {
				got = gotList[i]
				if got != want {
					t.Errorf("got %v want %v", got, want)
				}
			}
		})
	}
}
