package leetcode342

import "testing"

var data = []struct {
	in  int
	out bool
}{
	{0, false},
	{2, false},
	{4, true},
	{16, true},
	{20, false},
	{64, true},
	{100, false},
}

func TestAll(t *testing.T) {
	for _, d := range data {
		r := isPowerOfFour(d.in)
		if r != d.out {
			t.Error(d.out, r)
		}
	}

}
