package leetcode13

import (
	//"fmt"
	"testing"
)

var (
	data = []struct {
		input  string // input
		result int    // expected result
	}{
		{"III", 3},
		{"IV", 4},
		{"IX", 9},
		{"LVIII", 58},
		{"MCMXCIV", 1994},
	}
)

func TestRomanToInt(test *testing.T) {
	for _, d := range data {
		r := romanToInt(d.input)
		if r != d.result {
			test.Error("expectd:", d.result, "but got:", r)
		}

	}

}
