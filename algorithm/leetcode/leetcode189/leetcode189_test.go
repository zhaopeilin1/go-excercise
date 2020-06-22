package leetcode189

import "testing"

var data = []struct {
	in1 []int
	in2 int
}{
	//{[]int{1, 2, 3, 4, 5, 6, 7}, 3},
	{[]int{1, 2, 3, 4, 5, 6}, 2},
}

func TestAll(t *testing.T) {
	for _, d := range data {
		rotate(d.in1, d.in2)
		//t.Log(d.in1)
	}

}
