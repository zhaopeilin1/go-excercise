package leetcode171

import "testing"

var data = []struct {
	in  string
	out int
}{
	{"A", 1},
	//{2, "B"},
	{"Z", 26},
	//{27, "AA"},
	{"AZ", 52},
	//{701, "ZY"},
}

func TestAll(t *testing.T) {
	for _, d := range data {
		//r2 := pow(d.in, 2)
		//t.Log(r2)
		r := titleToNumber(d.in)
		// t.Log(d.out, r)
		if r != d.out {
			t.Error(d.out, r)
		}
	}
}
