package leetcode8

import "testing"

var data = []struct {
	in  string
	out int
}{
	{"-134 sadf", -134},
	{"1", 1},
	{"-1", -1},
	{"0", 0},
	{"134", 134},
	{"   -42 a", -42},
	{"words and 987", 0},
	{"-91283472332", -2147483648}, //[−2^31,  2^31 − 1]
	{"  0000000000012345678", 12345678},
	{"3.14159", 3},
	{"  -0012a42", -12},
	{"9223372036854775808", 2147483647},
	{"-5-", -5},
}

func TestAll(t *testing.T) {
	for _, d := range data {
		r := myAtoi(d.in)
		if r != d.out {
			t.Error(d.in, d.out, r)
		}
	}

}
