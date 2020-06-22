package leetcode219

import "testing"

var data = []struct {
	in1 []int
	in2 int
	out bool
}{
	{[]int{1, 2, 3, 1}, 3, true},
	{[]int{1, 0, 1, 1}, 1, true},
	{[]int{1, 2, 3, 1, 2, 3}, 2, false},
}

func TestAll(t *testing.T) {
	for _, d := range data {
		r := containsNearbyDuplicate(d.in1, d.in2)
		if r != d.out {
			t.Error(d.out, r)
		}
	}

}
