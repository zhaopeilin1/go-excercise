package leetcode5

import "testing"

var data1 = []struct {
	in  string
	out string
}{
	{"a", "^a$"},
	{"ab", "^a#b$"},
}

func TestAll(t *testing.T) {
	for _, d := range data1 {
		r := addSep(d.in)
		if r != d.out {
			t.Error(d.out, r)
		}
	}

}
