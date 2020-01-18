package ch2

import "testing"

var data = []struct {
	in int
}{
	{2},
	//{3},
}

func TestAll(t *testing.T) {
	for _, d := range data {
		genStrBy10(d.in)
	}

}
