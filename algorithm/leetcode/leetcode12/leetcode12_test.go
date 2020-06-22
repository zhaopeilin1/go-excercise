package leetcode12

import "testing"

var data = []struct {
	in  int
	out string
}{
	{3, "III"},
	{4, "IV"},
	{9, "IX"},
	{58, "LVIII"},
	{1994, "MCMXCIV"},
}

func TestAll(t *testing.T) {
	for _, d := range data {
		r := intToRoman(d.in)
		if r != d.out {
			t.Error(d.in, d.out, r)
		}
	}

}
