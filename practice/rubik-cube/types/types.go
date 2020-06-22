package types

import (
	"image/color"

	"gonum.org/v1/gonum/mat"
)

/*
三个方向，
x 正轴是从魔方中心指向我的方向，负轴是相反指向。
y 正轴是指从中心指向右边的方向，负轴是相反方向
z 正轴是从中心向上的方向，负轴是相反的方向

转动矩阵 T(x) 为绕x轴顺时针转动90度，T(x)^2 为绕x轴顺时针转动180度，T(x)^3 为绕x轴顺时针转动270度
转动矩阵 T(y) 为绕y轴顺时针转动90度，T(y)^2 为绕y轴顺时针转动180度，T(y)^3 为绕y轴顺时针转动270度
转动矩阵 T(z) 为绕z轴顺时针转动90度，T(z)^2 为绕y轴顺时针转动180度，T(z)^3 为绕y轴顺时针转动270度
*/
type RubikCube struct {
	Blocks []Block
}

type BlockColors int

const (
	_ BlockColors = iota
	W             //white
	R             //red
	G             //green
	O             //orange
	B             //blue
	Y             //yellow
	X             //not a color

	I int = 2
	J int = 3
	K int = 5
)

var colorMap map[BlockColors]color.Color = map[BlockColors]color.Color{
	W: color.White,
	R: color.RGBA{255, 0, 0, 255},
	G: color.RGBA{0, 255, 0, 255},
	O: color.RGBA{255, 140, 0, 266},
	//OR 88cc ?
	B: color.RGBA{0, 0, 255, 255},
	Y: color.RGBA{255, 255, 0, 255},
	X: color.Black,
}

var data []float64 = []float64{1, 0, 0, 0, 0, 1, 0, -1, 0}

var XR90 *mat.Dense = mat.NewDense(3, 3, data)

//const ()

//cube block
type Block struct {
	P Position
	C ColorDerect
}

//描述小块的位置
//(h2,k2,l2) = T(h1,k1,l1)
type Position struct {
	H int
	K int
	L int
}

//(i2,j2,k2) = T(i1,j1,k1)
type Direct struct {
	I int
	J int
	K int
}

//描述小块顔色取向
type ColorDerect struct {
	I int
	J int
	K int
}
