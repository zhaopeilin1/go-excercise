package leetcode122

import "testing"

var data = []struct {
	in  []int
	out int
}{
	{[]int{7, 1, 5, 3, 6, 4}, 7},
}

func TestAll(t *testing.T) {
	for _, d := range data {
		r := maxProfit(d.in)
		if r != d.out {
			t.Error(d.out, r)
		}
	}
}
