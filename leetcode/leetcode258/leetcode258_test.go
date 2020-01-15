package leetcode258

import "testing"

var data = []struct {
	in  int
	out int
}{
	{38, 2},
	{999, 9},
	{9998, 8},
}

func TestAll(t *testing.T) {
	for _, d := range data {
		r := addDigits(d.in)
		if r != d.out {
			t.Error(d.out, r)
		}
	}

}
