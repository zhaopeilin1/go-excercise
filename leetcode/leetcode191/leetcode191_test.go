package leetcode191

import "testing"

var data = []struct {
	in  uint32
	out int
}{
	{3, 2},
	{5, 2},
	{7, 3},
}

func TestAll(t *testing.T) {
	for _, d := range data {
		r := hammingWeight(d.in)
		if r != d.out {
			t.Error(d.out, r)
		}
	}
}
