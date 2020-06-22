package leetcode136

import "testing"

var data = []struct {
	in  []int
	out int
}{
	{[]int{2, 1, 2}, 1},
	{[]int{2, 2, 3}, 3},
	{[]int{4, 2, 2}, 4},
}

func TestAll(t *testing.T) {
	for _, d := range data {
		r := singleNumber(d.in)
		if r != d.out {
			t.Error(d.out, r)
		}
	}

}
