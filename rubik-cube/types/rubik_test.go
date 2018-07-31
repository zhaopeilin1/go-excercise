package types

import (
	"testing"
)

var colors = []struct {
	in     ColorDerect
	out    ColorDerect
	mirror ColorDerect
}{
	{ColorDerect{2, 3, 5}, ColorDerect{2, 3, 5}, ColorDerect{1, 1, 1}},
	{ColorDerect{-2, 3, 5}, ColorDerect{2, 3, 5}, ColorDerect{-1, 1, 1}},
	{ColorDerect{-2, -3, 5}, ColorDerect{2, 3, 5}, ColorDerect{-1, -1, 1}},
	{ColorDerect{-2, -3, -5}, ColorDerect{2, 3, 5}, ColorDerect{-1, -1, -1}},
}

func TestRubik(t *testing.T) {
	p1 := Position{h: 1, k: 1, l: 1}
	t.Log(p1)
	//做一个矩阵乘法
	//再做一个色块
}

func TestColor(t *testing.T) {
	t.Log(colorMap)
}

func TestMatrix(t *testing.T) {
	t.Log(XR90)
	//	XR90.MulElem()
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
	in  ColorDerect
	out ColorDerect
}{
	{ColorDerect{2, 3, 5}, ColorDerect{2, 3, 5}},
	{ColorDerect{-2, 5, 3}, ColorDerect{-2, 3, 5}},
	{ColorDerect{-5, -3, 2}, ColorDerect{2, -3, -5}},
	{ColorDerect{-3, -2, -5}, ColorDerect{-2, -3, -5}},
}

func TestRightHandApply(t *testing.T) {
	for _, righthand := range righthands {
		out := RightHandApply(righthand.in)
		if out != righthand.out {
			t.Errorf("expectd out %v, but got %v", righthand.out, out)
		}
	}

}
