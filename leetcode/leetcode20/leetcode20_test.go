package leetcode20

import (
	"testing"
)

var data = []struct {
	in  string
	out bool
}{
	{"()", true},
	{"()[]{}", true},
	{"(]", false},
	{"([)]", false},
}

func TestAll(t *testing.T) {

	for _, d := range data {
		r := isValid(d.in)
		if r != d.out {
			t.Error("expect :", d.out, " got:", r)
		}
	}

}
