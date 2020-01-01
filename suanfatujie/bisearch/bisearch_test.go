package bisearch

import "testing"

var data = []struct {
	in1 []int
	in2 int
	out int
}{
	{[]int{1, 3, 5, 7}, 7, 3},
	{[]int{1, 3, 5, 7}, 1, 0},
	{[]int{1, 2, 3, 7, 9}, 3, 2},
	{[]int{1, 2, 3, 7, 9, 13, 17}, 14, -1},
}

func TestAll(t *testing.T) {
	for _, d := range data {
		r := bisearch(d.in1, d.in2)
		if r != d.out {
			t.Error("expect", d.out, "got", r)

		}
	}

}
