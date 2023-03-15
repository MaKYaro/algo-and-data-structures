package linked_list

import (
	"fmt"
	"testing"
)

func TestDoublyLL_InsertHead(t *testing.T) {
	var tests = []struct {
		in  []any
		out []any
	}{
		{[]any{}, []any{}},
		{[]any{1, 2, 3, 4}, []any{4, 3, 2, 1}},
		{[]any{'q', 0, true, 9.81, "me"}, []any{"me", 9.81, true, 0, 'q'}},
	}

	for idx, test := range tests {
		name := fmt.Sprintf("case %v", idx)
		t.Run(name, func(t *testing.T) {
			list := MakeSinglyLL()
			for _, elem := range test.in {
				list.InsertHead(elem)
			}
			gotList := list.returnList()
			for i, want := range test.out {
				got := gotList[i]
				if want != got {
					t.Errorf("got %v want %v", got, want)
				}
			}
		})
	}
}

func TestDoublyLL_InsertTail(t *testing.T) {
	var tests = []struct {
		in  []any
		out []any
	}{
		{[]any{}, []any{}},
		{[]any{1, 2, 3, 4, 5, 6}, []any{1, 2, 3, 4, 5, 6}},
		{[]any{"test", true, 'r', 3.14, 0}, []any{"test", true, 'r', 3.14, 0}},
	}

	for idx, test := range tests {
		name := fmt.Sprintf("case %v", idx)
		t.Run(name, func(t *testing.T) {
			list := MakeDoublyLL()
			for _, elem := range test.in {
				list.InsertTail(elem)
			}
			gotList := list.returnList()
			for i, want := range test.out {
				got := gotList[i]
				if got != want {
					t.Errorf("got %v want %v", got, want)
				}
			}
		})
	}
}
