package leetcode204

import "testing"

var data = []struct {
	in  int
	out int
}{
	{10, 4},
}

func TestAll(t *testing.T) {
	for _, d := range data {
		r := countPrimes(d.in)
		if r != d.out {
			t.Error(d.out, r)
		}
	}

}
