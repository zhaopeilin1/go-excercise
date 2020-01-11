package leetcode168

import "testing"

var data = []struct {
	in  int
	out string
}{
	{1, "A"},
	//{2, "B"},
	{26, "Z"},
	//{27, "AA"},
	{52, "AZ"},
	//{701, "ZY"},
}

func TestAll(t *testing.T) {
	for _, d := range data {
		//r2 := pow(d.in, 2)
		//t.Log(r2)
		r := convertToTitle(d.in)
		// t.Log(d.out, r)
		if r != d.out {
			t.Error(d.out, r)
		}
	}
}
