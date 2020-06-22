package leetcode125

import "testing"

var data = []struct {
	in  string
	out bool
}{
	{"A man, a plan, a canal: Panama", true},
	{"A man, a plan, a canal: Panama1", false},
	{"azAZ019", false}, //97,122,65,90 48,49,57
}

func TestAll(t *testing.T) {
	for _, d := range data {
		r := isPalindrome(d.in)
		if r != d.out {
			t.Error(d.out, r)
		}
	}
}
