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

var colorXR90s = []struct {
	in  types.ColorDerect
	out types.ColorDerect
}{
	{types.ColorDerect{I, J, K}, types.ColorDerect{I, -J, K}},
	{types.ColorDerect{I, -J, K}, types.ColorDerect{I, -J, -K}},
	{types.ColorDerect{I, -J, -K}, types.ColorDerect{I, J, -K}},
	{types.ColorDerect{I, J, -K}, types.ColorDerect{I, J, K}},
}

func TestColorXR90(t *testing.T) {
	for _, colorXR90 := range colorXR90s {
		out := ColorXR90(colorXR90.in)
		if out != colorXR90.out {
			t.Errorf("expectd out %v, but got %v", colorXR90.out, out)
		}
	}
}

var positionXR90s = []struct {
	in  types.Position
	out types.Position
}{
	{types.Position{1, 1, 1}, types.Position{1, -1, 1}},
	{types.Position{1, -1, 1}, types.Position{1, -1, -1}},
	{types.Position{1, -1, -1}, types.Position{1, 1, -1}},
	{types.Position{1, 1, -1}, types.Position{1, 1, 1}},
}

func TestPositionXR90(t *testing.T) {
	for _, positionXR90 := range positionXR90s {
		out := PositionXR90(positionXR90.in)
		if out != positionXR90.out {
			t.Errorf("PositionXR90 expect %v but got %v", positionXR90.out, out)
		}

	}

}
