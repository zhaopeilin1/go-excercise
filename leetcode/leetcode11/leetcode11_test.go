package leetcode11

import "testing"

var data = []struct {
	in  []int
	out int
}{
	{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
}

func TestAll(t *testing.T) {
	for _, d := range data {
		r := maxArea(d.in)
		if r != d.out {
			t.Error(d.out, r)
		}
	}

}
