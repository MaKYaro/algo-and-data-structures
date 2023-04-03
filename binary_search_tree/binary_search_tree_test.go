package binary_search_tree

import (
	"fmt"
	"testing"
)

func TestBinarySearchTree_Insert(t *testing.T) {
	var tests = []struct {
		in  []int
		out []int
	}{
		{[]int{}, []int{}},
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{[]int{6, 17, 9, 4, 58, 38, 12, 54}, []int{4, 6, 9, 12, 17, 38, 54, 58}},
	}
	for idx, test := range tests {
		name := fmt.Sprintf("case %v", idx)
		t.Run(name, func(t *testing.T) {
			bst := MakeBinarySearchTree[int]()
			for _, elem := range test.in {
				bst.Insert(elem)
			}
			gotOut := bst.InOrderedTreeWalk()
			for i, want := range test.out {
				got := gotOut[i]
				if got != want {
					t.Errorf("got %v want %v", got, want)
				}
			}
		})
	}
}

func TestBinarySearchTree_Search(t *testing.T) {
	var tests = []struct {
		in   []int
		elem int
		out  int
		err  error
	}{
		{[]int{}, 0, -1, noSuchElement[int]{0}},
		{[]int{1}, 0, -1, noSuchElement[int]{0}},
		{[]int{1, 2, 3, 4, 5}, 3, 3, nil},
		{[]int{6, 17, 9, 4, 58, 38, 12, 54}, 9, 9, nil},
		{[]int{6, 17, 9, 4, 58, 38, 12, 54}, 10, -1, noSuchElement[int]{10}},
	}
	for idx, test := range tests {
		name := fmt.Sprintf("case %v", idx)
		t.Run(name, func(t *testing.T) {
			bst := MakeBinarySearchTree[int]()
			for _, elem := range test.in {
				bst.Insert(elem)
			}
			got, gotErr := bst.Search(test.elem)
			want, wantErr := test.out, test.err
			if gotErr != wantErr {
				t.Errorf("got error %v want error %v", gotErr, wantErr)
			}
			switch gotErr {
			case nil:
				if got.key != want {
					t.Errorf("got %v want %v", got, want)
				}
			default:
				if got != nil {
					t.Errorf("got %v want %v", got, nil)
				}
			}
		})
	}
}
