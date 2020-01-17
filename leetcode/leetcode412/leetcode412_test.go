package leetcode412

import (
	"reflect"
	"testing"
)

var data = []struct {
	in  int
	out []string
}{{15, []string{"1",
	"2",
	"Fizz",
	"4",
	"Buzz",
	"Fizz",
	"7",
	"8",
	"Fizz",
	"Buzz",
	"11",
	"Fizz",
	"13",
	"14",
	"FizzBuzz"}},
}

func TestAll(t *testing.T) {
	for _, d := range data {
		r := fizzBuzz(d.in)
		if !reflect.DeepEqual(r, d.out) {
			t.Error(d.out, r)
		}
	}

}
