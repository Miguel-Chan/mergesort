package mergesort

import (
	"testing"
)

func TestSort(t *testing.T) {
	cases := []struct {
		in, want []int
	}{
		{[]int{1, 2, 3, 4}, []int{1, 2, 3, 4}},
		{[]int{2, 1, 4, 5}, []int{1, 2, 4, 5}},
		{[]int{9, 8, 7, 5, 3, 1}, []int{1, 3, 5, 7, 8, 9}},
	}
	for _, c := range cases {
		got := getSlice(SortList(makeList(c.in)))
		flag := len(got) == len(c.want)
		for index := range got {
			if got[index] != c.want[index] {
				flag = false
			}
		}
		if !flag {
			t.Errorf("SortList(%v) == %v, want %v", c.in, got, c.want)
		}
	}
}
