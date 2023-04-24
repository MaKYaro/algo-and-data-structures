package red_black_tree

import (
	"fmt"
	bst "github.com/MaKYaro/algo-and-data-structures/binary_search_tree"
	"math"
	"testing"
)

func TestRedBlackTree_Insert(t *testing.T) {
	arr := make([]int, 100000)
	for i := 0; i < 100000; i++ {
		arr[i] = i
	}
	var tests = []struct {
		in  []int
		out []int
	}{
		{[]int{}, []int{}},
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{[]int{6, 17, 9, 4, 58, 38, 12, 54}, []int{4, 6, 9, 12, 17, 38, 54, 58}},
		{arr, arr},
	}
	for idx, test := range tests {
		name := fmt.Sprintf("case %v", idx)
		t.Run(name, func(t *testing.T) {
			rbt := MakeRedBlackTree[int]()
			for _, elem := range test.in {
				rbt.Insert(elem)
			}
			gotOut := rbt.InOrderedTreeWalk()
			for i, want := range test.out {
				got := gotOut[i]
				if got != want {
					t.Errorf("got %v want %v", got, want)
				}
			}
			treeHeight := float64(rbt.Height())
			if len(gotOut) > 0 && treeHeight >= 2*math.Log(float64(len(gotOut)+1)) {
				t.Error("Red-black tree is not balanced")
			}
		})
	}
}

func TestRedBlackTree_Search(t *testing.T) {
	var tests = []struct {
		in   []int
		elem int
		out  int
		err  error
	}{
		{[]int{}, 0, -1, bst.NoSuchKey[int]{0}},
		{[]int{1}, 0, -1, bst.NoSuchKey[int]{0}},
		{[]int{1, 2, 3, 4, 5}, 3, 3, nil},
		{[]int{6, 17, 9, 4, 58, 38, 12, 54}, 9, 9, nil},
		{[]int{6, 17, 9, 4, 58, 38, 12, 54}, 10, -1, bst.NoSuchKey[int]{10}},
	}
	for idx, test := range tests {
		name := fmt.Sprintf("case %v", idx)
		t.Run(name, func(t *testing.T) {
			rbt := MakeRedBlackTree[int]()
			for _, elem := range test.in {
				rbt.Insert(elem)
			}
			got, gotErr := rbt.Search(test.elem)
			want, wantErr := test.out, test.err
			if gotErr != wantErr {
				t.Errorf("got error %v want error %v", gotErr, wantErr)
			}
			switch gotErr {
			case nil:
				if got.Key != want {
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

func TestRedBlackTree_Delete(t *testing.T) {
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
			rbt := MakeRedBlackTree[int]()
			for _, elem := range test.in {
				rbt.Insert(elem)
			}
			n, _ := rbt.Search(test.elem)
			rbt.Delete(n)
			gotRbt := rbt.InOrderedTreeWalk()
			for i, want := range test.wantBst {
				got := gotRbt[i]
				if got != want {
					t.Errorf("got %v want %v", got, want)
				}
			}
			treeHeight := float64(rbt.Height())
			if len(gotRbt) > 0 && treeHeight >= 2*math.Log(float64(len(gotRbt)+1)) {
				t.Error("Red-black tree is not balanced")
			}
		})
	}
}
