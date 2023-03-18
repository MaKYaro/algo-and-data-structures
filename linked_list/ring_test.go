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

func TestRingLL_Delete(t *testing.T) {
	var tests = []struct {
		in  []any
		key any
		err error
		out []any
	}{
		{[]any{}, 1, noSuchElement{1}, []any{}},
		{[]any{'t'}, 't', nil, []any{}},
		{[]any{'t'}, 0, noSuchElement{0}, []any{'t'}},
		{[]any{1, 2, 3, 4}, 1, nil, []any{2, 3, 4}},
		{[]any{'q', true, 1, false, "qwe"}, true, nil, []any{'q', 1, false, "qwe"}},
		{[]any{"qf", 1, 4, 'p', true}, true, nil, []any{"qf", 1, 4, 'p'}},
		{[]any{2, 4, 0, "q"}, 3, noSuchElement{3}, []any{2, 4, 0, "q"}},
	}
	for idx, test := range tests {
		name := fmt.Sprintf("case %v", idx)
		t.Run(name, func(t *testing.T) {
			list := MakeRingLL()
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

func TestRingLL_Search(t *testing.T) {
	var tests = []struct {
		in  []any
		key any
		err error
		idx int
	}{
		{[]any{}, 1, noSuchElement{1}, -1},
		{[]any{'e'}, 'e', nil, 0},
		{[]any{'e'}, 9, noSuchElement{9}, -1},
		{[]any{1, 2, 3, 'q'}, 4, noSuchElement{4}, -1},
		{[]any{'w', 33, true, 7, 0}, 'w', nil, 0},
		{[]any{1, 0, 1, 0, true, 1}, true, nil, 4},
		{[]any{1, 0, 1, 0, 1, 0}, 0, nil, 1},
		{[]any{'u', 1, 9, "io"}, "io", nil, 3},
	}
	for idx, test := range tests {
		name := fmt.Sprintf("case %v", idx)
		t.Run(name, func(t *testing.T) {
			list := MakeRingLL()
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
