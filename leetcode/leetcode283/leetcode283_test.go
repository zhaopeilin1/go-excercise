package leetcode283

import "testing"

var data = []struct {
	in []int
}{
	{[]int{0, 1, 3, 2, 0, 4}},
	{[]int{0, 1, 0}},
}

func TestAll(t *testing.T) {
	for _, d := range data {
		moveZeroes(d.in)
		t.Log(d.in)

	}

}
