package leetcode645

import "testing"
import "reflect"

var data = []struct {
	in  []int
	out []int
}{
	{[]int{1, 2, 2, 4}, []int{2, 3}},
}

func TestAll(t *testing.T) {
	for _, d := range data {
		r := findErrorNums(d.in)
		if !reflect.DeepEqual(r, d.out) {
			t.Error(d.out, r)
		}
	}

}
