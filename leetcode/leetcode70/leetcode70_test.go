package leetcode70

import "testing"

var data = []struct {
	in  int
	out int
}{
	{2, 2},
	{3, 3},
	{4, 5},
	{5, 8},
	//{6, 25},
}

func TestAll(t *testing.T) {
	for _, d := range data {
		r := climbStairs3(d.in)
		if r != d.out {
			t.Error(d.out, r)
		}
	}

}
