package leetcode172

import "testing"

var data = []struct {
	in  int
	out int
}{
	{3, 0},
	{5, 1},
	{8, 1},
	{10, 2},
	{14, 2},
	{15, 3},
	{20, 4},
	{25, 6},
	{30, 7},
}

func TestAll(t *testing.T) {
	for _, d := range data {
		r := trailingZeroes(d.in)
		if r != d.out {
			t.Error(d.in, d.out, r)
		}
	}
}
