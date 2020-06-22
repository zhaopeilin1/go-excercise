package leetcode88

import "testing"

var data = []struct {
	in1 []int
	in2 int
	in3 []int

	out []int
}{
	{[]int{1, 5, 6, 0, 0, 0}, 3, []int{2, 3, 7}, []int{1, 2, 3, 5, 6, 7}},
	{[]int{0}, 0, []int{1}, []int{1}},

	{[]int{1}, 1, []int{}, []int{1}},
}

func TestAll(t *testing.T) {
	for _, d := range data {
		merge(d.in1, d.in2, d.in3, len(d.in3))
		t.Log(d.in1, d.out)
	}
}
