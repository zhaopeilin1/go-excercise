package leetcode383

import "testing"

var data = []struct {
	in1 string
	in2 string
	out bool
}{
	{"fihjjjjei", "hjibagacbhadfaefdjaeaebgi", false},
}

func TestAll(t *testing.T) {
	for _, d := range data {
		r := canConstruct(d.in1, d.in2)
		if r != d.out {
			t.Error(d.out, r)
		}
	}

}
