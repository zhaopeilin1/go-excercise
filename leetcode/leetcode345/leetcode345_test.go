package leetcode345

import "testing"

var data = []struct {
	in  string
	out string
}{
	{"leetcode", "leotcede"},
	{"hello", "holle"},
	{"a.b,.", "a.b,."},
}

func TestAll(t *testing.T) {
	for _, d := range data {
		r := reverseVowels(d.in)
		if r != d.out {
			t.Error(d.out, r)
		}
	}

}
