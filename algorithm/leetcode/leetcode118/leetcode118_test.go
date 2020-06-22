package leetcode118

import "testing"

var data = []struct {
	in  int
	out [][]int
}{
	{5, [][]int{}},
}

func TestAll(t *testing.T) {
	for _, d := range data {
		r := generate(d.in)
		t.Log(r)
	}

}
