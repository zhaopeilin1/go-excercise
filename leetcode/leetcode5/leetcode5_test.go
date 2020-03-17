package leetcode5

import "testing"

var data1 = []struct {
	in  string
	out string
}{
	{"a", "^a$"},
	{"ab", "^a#b$"},
}

var data2 = []struct {
	in  string
	out string
}{
	{"a", "a"},
	{"aba", "aba"},
	{"abac", "aba"},
	{"babad", "bab"},
	{"cbbd", "bb"},
	{"ac", "a"},
}

func TestAll(t *testing.T) {
	for _, d := range data1 {
		r := addSep(d.in)
		if r != d.out {
			t.Error(d.out, r)
		}
	}
}

func TestAll2(t *testing.T) {
	for _, d := range data2 {
		r := longestPalindrome(d.in)
		if r != d.out {
			t.Error(d.out, r)
		}
	}
}
