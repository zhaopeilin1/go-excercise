package types

import (
	"testing"
)

func TestRubik(t *testing.T) {
	p1 := Position{1, 1, 1}
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
