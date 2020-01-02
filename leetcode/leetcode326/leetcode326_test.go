package leetcode326

import "testing"

var data = []struct{
	in int
	out bool
}{
	{3,true},
	{9,true},
	{12,false},
	{2,false},
}

func TestAll(t *testing.T) {
	for _,d:=range data {
		r:=isPowerOfThree(d.in)
		if r!=d.out {
			t.Error(d.out,r)
		}
	}

}
