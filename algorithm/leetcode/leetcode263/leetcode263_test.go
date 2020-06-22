package leetcode263

import "testing"

var data = []struct {
	in  int
	out bool
}{
	{10, true},
	{11, false},
	{14, false},
}

func TestAll(t *testing.T) {
	for _, d := range data {
		r := isUgly(d.in)
		if r != d.out {
			t.Error(d.out, r)
		}
	}

}
