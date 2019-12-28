package leetcode14

import (
	"testing"
)

var (
	data = []struct {
		input  []string // input
		result string   // expected result
	}{
		{[]string{"flower", "flow", "flight"}, "fl"},

		{[]string{"dog", "d", "doa"}, "d"},
		{[]string{"aabb", "f", "flight"}, ""},
		{[]string{"aa", "aaa", "ad"}, "a"},
	}
)

func TestAll(t *testing.T) {
	for _, d := range data {
		r := longestCommonPrefix(d.input)
		if r != d.result {
			t.Error("expect :", d.result, " got:", r)

		}

	}

}
