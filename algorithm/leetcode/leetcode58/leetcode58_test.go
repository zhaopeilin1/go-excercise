package leetcode58

import "testing"

var data = []struct {
	in  string
	out int
}{
	{"HelloWorld", 10},
	{"Hello World", 5},
	{"Hello World ", 5},
	{"a ", 1},
}

func TestAll(t *testing.T) {
	for _, d := range data {
		r := lengthOfLastWord(d.in)
		if r != d.out {
			t.Error(d.in, d.out, r)
		}

	}

}
