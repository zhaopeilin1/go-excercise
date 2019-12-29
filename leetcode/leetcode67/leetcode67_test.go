package leetcode67

import "testing"

var data = []struct {
	in1 string
	in2 string
	out string
}{
	{"101", "101", "1010"},
	{"1", "10", "11"},
	{"11", "11", "110"},
	{"11", "1", "100"},
}

func TestAll(t *testing.T) {
	for _, d := range data {
		r := addBinary(d.in1, d.in2)
		if r != d.out {
			t.Error("error", d.out, r)
		}
	}

}
