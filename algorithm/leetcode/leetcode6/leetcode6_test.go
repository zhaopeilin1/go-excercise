package leetcode6

import "testing"

var data = []struct {
	in1 string
	in2 int
	out string
}{
	{"LEETCODEISHIRING", 3, "LCIRETOESIIGEDHN"},
	{"LEETCODEISHIRING", 4, "LDREOEIIECIHNTSG"},
	{"A", 1, "A"},
}

func TestAll(t *testing.T) {
	for _, d := range data {
		r := convert(d.in1, d.in2)
		if r != d.out {
			t.Error(d.out, r)
		}

	}

}
