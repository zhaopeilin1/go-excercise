package leetcode371

import "testing"

var data = []struct {
	in1 int
	in2 int
	out int
}{
	{1, 2, 3},
	{111, 222, 333},
	{-21, 2, -19},
	{-21, 22, 1},
}

func TestAll(t *testing.T) {
	for _, d := range data {
		r := getSum(d.in1, d.in2)
		if r != d.out {
			t.Error(d.out, r)
		}
	}

}
