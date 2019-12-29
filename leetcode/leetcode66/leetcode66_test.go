package leetcode66

import (
	"fmt"
	"testing"
)

var data = []struct {
	in  []int
	out []int
}{
	{[]int{1, 2, 3}, []int{1, 2, 4}},
	{[]int{9, 9}, []int{1, 0, 0}},
	{[]int{1, 9, 9}, []int{2, 0, 0}},
}

func TestAll(t *testing.T) {
	for _, d := range data {
		r := plusOne(d.in)

		fmt.Println(d.out, r)

	}

}
