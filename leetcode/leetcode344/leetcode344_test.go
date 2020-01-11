package leetcode344

import "testing"

var data = []struct {
	in  []byte
	out []byte
}{
	{[]byte("hello"), []byte("aa")},
}

func TestAll(t *testing.T) {
	for _, d := range data {
		reverseString(d.in)
		t.Log(string(d.in))

	}

}
