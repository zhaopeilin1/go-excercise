package leetcode28

import "testing"

var data = []struct {
	in1 string
	in2 string
	out int
}{
	{"hello", "ll", 2},
	{"aaaaa", "bba", -1},
}

func TestAll(t *testing.T) {
	for _, d := range data {
		r := strStr(d.in1, d.in2)
		if r != d.out {
			t.Error("expected:", d.out, " got:", r)
		}

	}

}
