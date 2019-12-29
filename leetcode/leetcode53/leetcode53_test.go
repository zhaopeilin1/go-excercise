package leetcode53

import "testing"

var data = []struct {
	in  []int
	out int
}{
	{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
	{[]int{-2}, -2},
	{[]int{-2, 0}, 0},
	{[]int{-2, -4}, -2},
	{[]int{2, 0}, 2},
	{[]int{2, 20}, 22},
	{[]int{0, 0}, 0},
}

func TestAll(t *testing.T) {
	for _, d := range data {
		r := maxSubArray(d.in)
		if r != d.out {
			t.Error(d.out, r)
		}
	}
}
