package rotate

import (
	"testing"

	"github.com/zhaopeilin1/go-excercise/rubik-cube/types"
)

const (
	I = types.I
	J = types.J
	K = types.K
)

var colors = []struct {
	in     types.ColorDerect
	out    types.ColorDerect
	mirror types.ColorDerect
}{
	{types.ColorDerect{I, J, K}, types.ColorDerect{I, J, K}, types.ColorDerect{1, 1, 1}},
	{types.ColorDerect{-I, J, K}, types.ColorDerect{I, J, K}, types.ColorDerect{-1, 1, 1}},
	{types.ColorDerect{-I, -J, K}, types.ColorDerect{I, J, K}, types.ColorDerect{-1, -1, 1}},
	{types.ColorDerect{-I, -J, -K}, types.ColorDerect{I, J, K}, types.ColorDerect{-1, -1, -1}},
}

func TestMirror(t *testing.T) {
	for _, collor := range colors {
		out, mirror := Mirror(collor.in)
		if out != collor.out {
			t.Errorf("expectd out %v, but got %v", collor.out, out)
		}
		if mirror != collor.mirror {
			t.Errorf("expectd mirror %v, but got %v", collor.mirror, mirror)
		}
	}
}

var righthands = []struct {
	in  types.ColorDerect
	out types.ColorDerect
}{
	{types.ColorDerect{I, J, K}, types.ColorDerect{I, J, K}},
	{types.ColorDerect{-I, K, J}, types.ColorDerect{-I, J, K}},
	{types.ColorDerect{-K, -J, I}, types.ColorDerect{I, -J, -K}},
	{types.ColorDerect{-J, -I, -K}, types.ColorDerect{-I, -J, -K}},
}

func TestRightHandApply(t *testing.T) {
	for _, righthand := range righthands {
		out := RightHandApply(righthand.in)
		if out != righthand.out {
			t.Errorf("expectd out %v, but got %v", righthand.out, out)
		}
	}
}

var colorR90s = []struct {
	in  types.ColorDerect
	out types.ColorDerect
}{
	{types.ColorDerect{I, J, K}, types.ColorDerect{I, -J, K}},
	{types.ColorDerect{I, -J, K}, types.ColorDerect{I, -J, -K}},
	{types.ColorDerect{I, -J, -K}, types.ColorDerect{I, J, -K}},
	{types.ColorDerect{I, J, -K}, types.ColorDerect{I, J, K}},
}

func TestColorR90(t *testing.T) {
	for _, colorR90 := range colorR90s {
		out := ColorR90(colorR90.in)
		if out != colorR90.out {
			t.Errorf("expectd out %v, but got %v", colorR90.out, out)
		}
	}
}
