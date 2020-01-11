package leetcode190

import "testing"

var data = []struct {
	in  uint32
	out uint32
}{
	{43261596, 964176192},
}

func TestAll(t *testing.T) {
	for _, d := range data {
		r := reverseBits(d.in)
		if r != d.out {
			t.Error(d.out, r)
		}
	}

}
