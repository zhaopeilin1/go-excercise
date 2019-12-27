package leetcode26

import (
	"fmt"
	"testing"
)

var data = []struct {
	in  []int
	out int
}{
	{[]int{1, 1, 3, 3, 5, 6}, 4},
}

func TestAll(t *testing.T) {
	for _, d := range data {
		r := removeDuplicates(d.in)
		fmt.Println(d.in)
		if r != d.out {
			t.Error("expected:", d.out, " got:", r)
		}
	}

}
