package types

import (
	"image/color"

	//	"math"
	"sort"

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

)

var colorMap map[BlockColors]color.Color = map[BlockColors]color.Color{
	W: color.White,
	R: color.RGBA{255, 0, 0, 1},
	G: color.RGBA{0, 255, 0, 1},
	O: color.RGBA{255, 140, 0, 1},
	//OR 88cc ?
	B: color.RGBA{0, 0, 255, 1},
	Y: color.RGBA{255, 255, 0, 1},
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
	h int
	k int
	l int
}

//(i2,j2,k2) = T(i1,j1,k1)
type Direct struct {
	i int
	j int
	k int
}

//描述小块顔色取向
type ColorDerect struct {
	i int
	j int
	k int
}

//mirror
//return 1:a mirror ColorDerect,2:the transform process
//example: 1,-1,1 return 1,1,1  1,-1,1
func Mirror(in ColorDerect) (out, mirror ColorDerect) {
	out = ColorDerect{IntAbs(in.i), IntAbs(in.j), IntAbs(in.k)}
	mirror = ColorDerect{out.i / in.i, out.j / in.j, out.k / in.k}
	return
}

//do right hand to a Collor  reorder i,j,k
//example {2,3,5} -> {2,3,5} {2,5,3}->{2,3,5}, {-5,-3,-2}->{-2,-3,-5}
//{i,j,k}->{i,j,k} {i,-k,-j}->{i,-j,-k}...
func RightHandApply(in ColorDerect) ColorDerect {
	AbsI := IntAbs(in.i)
	AbsJ := IntAbs(in.j)
	Absk := IntAbs(in.k)
	var m map[int]int = map[int]int{AbsI: in.i, AbsJ: in.j, Absk: in.k}
	list := []int{AbsI, AbsJ, Absk}
	sort.Ints(list)
	return ColorDerect{m[list[0]], m[list[1]], m[list[2]]}
}

func IntAbs(i int) int {
	if i >= 0 {
		return i
	} else {
		return -i
	}
}
