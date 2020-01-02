package leetcode119

import "testing"

var data = []struct {
	in  int
	out []int
}{
	{0, []int{1}},
	{1, []int{1, 1}},
	{2, []int{1, 2, 1}},
	{3, []int{1, 3, 3, 1}},
}

func TestAll(t *testing.T) {
	for _, d := range data {
		r := getRow(d.in)
		if len(r) != len(d.out) {
			t.Error("len", d.out, r)
		} /* else {
			for i, v := range r {
				if i != d.out[i] {
					t.Error(d.out[i], v)
				}
			}

		}*/
		t.Log(r)
		t.Log(d.out)

	}

}
