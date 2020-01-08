package leetcode605

import "testing"

var data = []struct {
	in1 []int
	in2 int
	out bool
}{
	{[]int{1, 0, 0, 0, 1}, 1, true},
	{[]int{1, 0, 0, 0, 1}, 2, false},
	{[]int{0, 1, 0, 0, 0, 1}, 2, false},
	{[]int{1, 0, 0, 0, 1, 0}, 2, false},
	{[]int{0, 1, 0, 0, 0, 1, 0}, 2, false},
	{[]int{0, 1, 0, 0, 0, 1, 0, 0, 1}, 2, false},
	{[]int{1, 0, 0, 0, 1, 0, 0}, 2, true},
}

func TestAll(t *testing.T) {
	for _, d := range data {
		r := canPlaceFlowers(d.in1, d.in2)
		if r != d.out {
			t.Error(d.in1, d.out, r)
		}
	}

}
