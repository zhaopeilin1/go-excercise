package leetcode168

import "testing"

var data = []struct {
	in  int
	out string
}{
	//{1, "A"},
	{2, "B"},
	{27, "AA"},
	{701, "ZY"},
}

func TestAll(t *testing.T) {
	for _, d := range data {
		r1, r2 := convertToTitle2(d.in)
		t.Log(r1, r2)
		//r := convertToTitle(d.in)
		// if r != d.out {
		// 	t.Error(d.out, r)
		// }
	}

}
