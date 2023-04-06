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
		{[]int{}, 0, -1, NoSuchKey[int]{0}},
		{[]int{1}, 0, -1, NoSuchKey[int]{0}},
		{[]int{1, 2, 3, 4, 5}, 3, 3, nil},
		{[]int{6, 17, 9, 4, 58, 38, 12, 54}, 9, 9, nil},
		{[]int{6, 17, 9, 4, 58, 38, 12, 54}, 10, -1, NoSuchKey[int]{10}},
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

func TestBinarySearchTree_Min(t *testing.T) {
	var tests = []struct {
		in  []int
		min int
		err error
	}{
		{[]int{}, 0, EmptyTree[int]{}},
		{[]int{1, 2, 3, 4, 5}, 1, nil},
		{[]int{6, 17, 9, 4, 58, 38, 12, 54}, 4, nil},
	}
	for idx, test := range tests {
		name := fmt.Sprintf("case %v", idx)
		t.Run(name, func(t *testing.T) {
			bst := MakeBinarySearchTree[int]()
			for _, elem := range test.in {
				bst.Insert(elem)
			}
			gotNode, gotErr := bst.Min()
			want, wantErr := test.min, test.err
			if gotErr != wantErr {
				t.Errorf("got error %v want error %v", gotErr, wantErr)
			}
			if gotErr == nil {
				got := gotNode.key
				if got != want {
					t.Errorf("got %v want %v", got, want)
				}
			}
		})
	}
}

func TestBinarySearchTree_Max(t *testing.T) {
	var tests = []struct {
		in  []int
		min int
		err error
	}{
		{[]int{}, 0, EmptyTree[int]{}},
		{[]int{1, 2, 3, 4, 5}, 5, nil},
		{[]int{6, 17, 9, 4, 58, 38, 12, 54}, 58, nil},
	}
	for idx, test := range tests {
		name := fmt.Sprintf("case %v", idx)
		t.Run(name, func(t *testing.T) {
			bst := MakeBinarySearchTree[int]()
			for _, elem := range test.in {
				bst.Insert(elem)
			}
			gotNode, gotErr := bst.Max()
			want, wantErr := test.min, test.err
			if gotErr != wantErr {
				t.Errorf("got error %v want error %v", gotErr, wantErr)
			}
			if gotErr == nil {
				got := gotNode.key
				if got != want {
					t.Errorf("got %v want %v", got, want)
				}
			}
		})
	}
}

func TestNode_Successor(t *testing.T) {
	var tests = []struct {
		in   []int
		elem int
		suc  int
		err  error
	}{
		{[]int{1, 2, 3, 4, 5}, 5, 0, NoSuccessor[int]{5}},
		{[]int{6, 17, 9, 4, 58, 38, 12, 54}, 4, 6, nil},
		{[]int{6, 17, 9, 4, 58, 38, 12, 54}, 6, 9, nil},
		{[]int{6, 17, 9, 4, 58, 38, 12, 54}, 9, 12, nil},
		{[]int{6, 17, 9, 4, 58, 38, 12, 54}, 12, 17, nil},
		{[]int{6, 17, 9, 4, 58, 38, 12, 54}, 17, 38, nil},
	}
	for idx, test := range tests {
		name := fmt.Sprintf("case %v", idx)
		t.Run(name, func(t *testing.T) {
			bst := MakeBinarySearchTree[int]()
			for _, elem := range test.in {
				bst.Insert(elem)
			}
			n, _ := bst.Search(test.elem)
			gotNode, gotErr := n.Successor()
			want, wantErr := test.suc, test.err
			if gotErr != wantErr {
				t.Errorf("got error %v want error %v", gotErr, wantErr)
			}
			if gotErr == nil {
				got := gotNode.key
				if got != want {
					t.Errorf("got %v want %v", got, want)
				}
			}
		})
	}
}

func TestNode_Predecessor(t *testing.T) {
	var tests = []struct {
		in   []int
		elem int
		pre  int
		err  error
	}{
		{[]int{1, 2, 3, 4, 5}, 1, 0, NoSuccessor[int]{1}},
		{[]int{6, 17, 9, 4, 58, 38, 12, 54}, 6, 4, nil},
		{[]int{6, 17, 9, 4, 58, 38, 12, 54}, 9, 6, nil},
		{[]int{6, 17, 9, 4, 58, 38, 12, 54}, 12, 9, nil},
		{[]int{6, 17, 9, 4, 58, 38, 12, 54}, 17, 12, nil},
		{[]int{6, 17, 9, 4, 58, 38, 12, 54}, 38, 17, nil},
	}
	for idx, test := range tests {
		name := fmt.Sprintf("case %v", idx)
		t.Run(name, func(t *testing.T) {
			bst := MakeBinarySearchTree[int]()
			for _, elem := range test.in {
				bst.Insert(elem)
			}
			n, _ := bst.Search(test.elem)
			gotNode, gotErr := n.Predecessor()
			want, wantErr := test.pre, test.err
			if gotErr != wantErr {
				t.Errorf("got error %v want error %v", gotErr, wantErr)
			}
			if gotErr == nil {
				got := gotNode.key
				if got != want {
					t.Errorf("got %v want %v", got, want)
				}
			}
		})
	}
}

func TestBinarySearchTree_Delete(t *testing.T) {
	var tests = []struct {
		in      []int
		elem    int
		wantBst []int
	}{
		{[]int{1}, 1, []int{}},
		{[]int{1, 2, 3, 4}, 4, []int{1, 2, 3}},
		{[]int{6, 17, 9, 4, 58, 38, 12, 54}, 12, []int{4, 6, 9, 17, 38, 54, 58}},
		{[]int{6, 17, 9, 4, 58, 38, 12, 54}, 17, []int{4, 6, 9, 12, 38, 54, 58}},
	}
	for idx, test := range tests {
		name := fmt.Sprintf("case %v", idx)
		t.Run(name, func(t *testing.T) {
			bst := MakeBinarySearchTree[int]()
			for _, elem := range test.in {
				bst.Insert(elem)
			}
			n, _ := bst.Search(test.elem)
			bst.Delete(n)
			gotBst := bst.InOrderedTreeWalk()
			for i, want := range test.wantBst {
				got := gotBst[i]
				if got != want {
					t.Errorf("got %v want %v", got, want)
				}
			}
		})
	}
}
