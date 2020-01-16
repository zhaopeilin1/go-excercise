package leetcode205

import "testing"

var data = []struct {
	in1 string
	in2 string
	out bool
}{
	{"aabb", "ccdd", true},
	{"aabb", "cccc", false},
	{"aabb", "ccab", false},
}

func TestAll(t *testing.T) {
	for _, d := range data {
		r := isIsomorphic(d.in1, d.in2)
		if r != d.out {
			t.Error(d.out, r)
		}
	}

}
