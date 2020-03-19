package leetcode15

import "testing"
import "reflect"

var data = []struct {
	in  []int
	out [][]int
}{
	{[]int{-1, 0, 1, 2, -1, -4}, [][]int{[]int{-1, 0, 1}, []int{-1, -1, 2}}},
}

func TestAll(t *testing.T) {
	for _, d := range data {
		r := threeSum(d.in)
		if !reflect.DeepEqual(r, d.out) {
			t.Error(d.out, r)
		}

	}

}
