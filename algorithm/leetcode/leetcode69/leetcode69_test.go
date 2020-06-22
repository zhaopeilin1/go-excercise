package leetcode69

import "testing"

var data = []struct {
	in  int
	out int
}{
	{1, 1},
	{3, 1},
	{4, 2},
	{8, 2},
	{9, 3},
	{10, 3},
	{16, 4},
	{25, 5},
	{100, 10},
	{120, 10},
	{29792276, 5458},
}

func TestAll(t *testing.T) {
	for _, d := range data {
		r := mySqrt(d.in)
		if r != d.out {
			t.Error(d.out, r)
		}
	}

}
