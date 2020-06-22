package leetcode367

import "testing"

var data = []struct {
	in  int
	out bool
}{
	{1, true},
	{2, false},
	{3, false},
	{4, true},
	{5, false},
}

func TestAll(t *testing.T) {
	for _, d := range data {
		r := isPerfectSquare(d.in)
		if r != d.out {
			t.Error(d.out, r)
		}
	}

}
