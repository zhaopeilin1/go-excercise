package leetcode202

import "testing"

var data = []struct {
	in  int
	out bool
}{
	{19, true},
}

func TestAll(t *testing.T) {
	for _, d := range data {
		r := isHappy(d.in)
		if r != d.out {
			t.Error(d.out, r)
		}
	}

}
