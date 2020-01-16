package leetcode299

import "testing"

var data = []struct {
	in1 string
	in2 string
	out string
}{
	{"1807", "7810", "1A3B"},
	{"1123", "0111", "1A1B"},
}

func TestAll(t *testing.T) {
	for _, d := range data {
		r := getHint(d.in1, d.in2)
		if r != d.out {
			t.Error(d.out, r)
		}
	}

}
