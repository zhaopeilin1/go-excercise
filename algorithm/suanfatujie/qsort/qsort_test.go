package qsort

import (
	"fmt"
	"testing"
)

var data = []struct {
	in  []int
	out []int
}{
	{[]int{2, 3, 1, 4, 7, 6, 5}, []int{1, 2, 3, 4, 5, 6, 7}},
}

func TestAll(t *testing.T) {
	for _, d := range data {
		qsort(d.in)
		fmt.Println(d.in)
	}

}
