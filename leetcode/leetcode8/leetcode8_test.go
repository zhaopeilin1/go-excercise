package leetcode8

import "testing"

var data = []struct {
	in  string
	out int
}{
	{"-134 sadf", -134},
	{"134", 134},
	{"   -42", -42},
	{"words and 987", 0},
	{"-91283472332", 0}, //[−2^31,  2^31 − 1]
}

func TestAll(t *testing.T) {
	for _, d := range data {
		r := myAtoi(d.in)
		if r != d.out {
			t.Error(d.out, r)
		}
	}

}
