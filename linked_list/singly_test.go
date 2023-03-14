package linked_list

import (
	"fmt"
	"testing"
)

func TestSinglyLL_InsertHead(t *testing.T) {
	var tests = []struct {
		in  []any
		out []any
	}{
		{[]any{1, 2, 3}, []any{3, 2, 1}},
		{[]any{}, []any{}},
		{[]any{'q', 4, struct{}{}, 0, true}, []any{true, 0, struct{}{}, 4, 'q'}},
	}

	for idx, test := range tests {
		name := fmt.Sprintf("case %v", idx)
		t.Run(name, func(t *testing.T) {
			list := MakeSinglyLL()
			for _, elem := range test.in {
				list.InsertHead(elem)
			}
			got := list.returnList()
			for i, want := range test.out {
				if want != got[i] {
					t.Errorf("got %v want %v", got[i], want)
				}
			}
		})
	}
}

func TestSinglyLL_InsertTail(t *testing.T) {
	var tests = []struct {
		in  []any
		out []any
	}{
		{[]any{1, 2, 3}, []any{1, 2, 3}},
		{[]any{}, []any{}},
		{[]any{'1', 0, true, false, "qwer"}, []any{'1', 0, true, false, "qwer"}},
	}

	for idx, test := range tests {
		name := fmt.Sprintf("case %v", idx)
		t.Run(name, func(t *testing.T) {
			list := MakeSinglyLL()
			for _, elem := range test.in {
				list.InsertTail(elem)
			}
			got := list.returnList()
			for i, want := range test.out {
				if want != got[i] {
					t.Errorf("got %v want %v", got[i], want)
				}
			}
		})
	}
}

func TestSinglyLL_Delete(t *testing.T) {
	var tests = []struct {
		in  []any
		key any
		err error
		out []any
	}{
		{[]any{}, 1, noSuchElement{1}, []any{}},
		{[]any{1, 2, 3, 4}, 1, nil, []any{2, 3, 4}},
		{[]any{'q', true, 1, false, "qwe"}, true, nil, []any{'q', 1, false, "qwe"}},
		{[]any{"qf", 1, 4, 'p', true}, true, nil, []any{"qf", 1, 4, 'p'}},
		{[]any{2, 4, 0, "q"}, 3, noSuchElement{3}, []any{2, 4, 0, "q"}},
	}
	for idx, test := range tests {
		name := fmt.Sprintf("case %v", idx)
		t.Run(name, func(t *testing.T) {
			list := MakeSinglyLL()
			for _, elem := range test.in {
				list.InsertTail(elem)
			}
			err := list.Delete(test.key)
			if err != test.err {
				t.Errorf("should be another error %e", err)
			}
			resultList := list.returnList()
			for i, want := range test.out {
				got := resultList[i]
				if got != want {
					t.Errorf("got %v want %v", got, want)
				}
			}
		})
	}
}

func TestSinglyLL_Search(t *testing.T) {
	var tests = []struct {
		in  []any
		key any
		err error
		idx int
	}{
		{[]any{}, 1, noSuchElement{1}, -1},
		{[]any{1, 2, 3, 'q'}, 4, noSuchElement{4}, -1},
		{[]any{'w', 33, true, 7, 0}, 'w', nil, 0},
		{[]any{1, 0, 1, 0, true, 1}, true, nil, 4},
		{[]any{1, 0, 1, 0, 1, 0}, 0, nil, 1},
		{[]any{'u', 1, 9, "io"}, "io", nil, 3},
	}
	for idx, test := range tests {
		name := fmt.Sprintf("case %v", idx)
		t.Run(name, func(t *testing.T) {
			list := MakeSinglyLL()
			for _, elem := range test.in {
				list.InsertTail(elem)
			}
			got, gotErr := list.Search(test.key)
			want, wantErr := test.idx, test.err
			if got != want {
				t.Errorf("got %v want %v", got, want)
			}
			if gotErr != wantErr {
				t.Errorf("should be another error %e", wantErr)
			}
		})
	}
}
