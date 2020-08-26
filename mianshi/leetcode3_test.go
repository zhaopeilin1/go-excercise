package leetcode3

import "testing"

var data = []struct {
	in  string
	out int
}{
	{"abcabcbb", 3},
	{"bbbbb", 1},
	{" ", 1},
	{"b", 1},
	{"aab", 2},
	{"pwwkew", 3},
	{"bpfbhmipx", 7},
}

func TestAll(t *testing.T) {
	for _, d := range data {
		r := lengthOfLongestSubstring(d.in)
		if r != d.out {
			t.Error("expected:", d.out, " got:", r)
		}
	}

}
