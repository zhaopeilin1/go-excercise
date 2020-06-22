package leetcode50

import "testing"

var data = []struct {
	in1 float64
	in2 int
	out float64
}{
	{2.00000, 0, 1.00000},
	{2.00000, 1, 2.00000},
	{2.00000, 4, 16.00000},
	{2.00000, 10, 1024.00000},
	{2.10000, 3, 9.26100},
	{2.00000, -2, 0.25000},
}

func TestAll(t *testing.T) {
	for _, d := range data {
		r := myPow(d.in1, d.in2)
		if r != d.out {
			t.Error(d.in1, d.in2, d.out, r)
		}
	}

}
