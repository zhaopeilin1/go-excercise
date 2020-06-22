package leetcode38

import "testing"

var data = []struct {
	in  int
	out string
}{
	{1, "1"},
	{2, "11"},
	{3, "21"},
	{4, "1211"},
	{5, "111221"},
}

func TestAll(t *testing.T) {
	for _, d := range data {
		r := countAndSay(d.in)
		if r != d.out {
			t.Error("expect:", d.out, " got:", r)

		}
	}

}
