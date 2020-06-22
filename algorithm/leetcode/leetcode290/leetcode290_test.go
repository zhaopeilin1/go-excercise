package leetcode290

import "testing"

var data = []struct {
	in1 string
	in2 string
	out bool
}{
	//{,,true},
	{"abba", "dog cat cat fish", false},
}

func TestAll(t *testing.T) {
	for _, d := range data {
		r := wordPattern(d.in1, d.in2)
		if r != d.out {
			t.Error(d.out, r)
		}
	}

}
