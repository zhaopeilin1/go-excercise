package leetcode1122

import "testing"
import "reflect"

var data = []struct {
	in1 []int
	in2 []int
	out []int
}{
	{[]int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19}, []int{2, 1, 4, 3, 9, 6}, []int{2, 2, 2, 1, 4, 3, 3, 9, 6, 7, 19}},
}

func TestAll(t *testing.T) {
	for _, d := range data {
		r := relativeSortArray(d.in1, d.in2)
		if !reflect.DeepEqual(r, d.out) {
			t.Error(d.out, r)
		}
	}
}
