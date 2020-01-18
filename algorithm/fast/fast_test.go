package fast

import "testing"

var data = []struct {
	in1 int
	in2 int
	out int
}{
	{3, 3, 27},
	//{3, 15, 14348907},
}

func TestAll(t *testing.T) {
	for _, d := range data {
		r := pow(d.in1, d.in2)
		//t.Log(r)
		if r != d.out {
			t.Error(d.out, r)
		}
	}

}
