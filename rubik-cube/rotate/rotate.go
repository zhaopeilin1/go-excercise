package rotate

import (
	"sort"

	"github.com/zhaopeilin1/go-excercise/rubik-cube/types"

	"gonum.org/v1/gonum/mat"
)

//mirror
//return 1:a mirror ColorDerect,2:the transform process
//example: 1,-1,1 return 1,1,1  1,-1,1
func Mirror(in types.ColorDerect) (out, mirror types.ColorDerect) {
	out = types.ColorDerect{IntAbs(in.I), IntAbs(in.J), IntAbs(in.K)}
	mirror = types.ColorDerect{out.I / in.I, out.J / in.J, out.K / in.K}
	return
}

//
func ReMirror(in, mirror types.ColorDerect) types.ColorDerect {
	return types.ColorDerect{mirror.I * in.I, mirror.J * in.J, mirror.K * in.K}

}

//do right hand to a Collor  reorder i,j,k
//example {2,3,5} -> {2,3,5} {2,5,3}->{2,3,5}, {-5,-3,-2}->{-2,-3,-5}
//{i,j,k}->{i,j,k} {i,-k,-j}->{i,-j,-k}...
func RightHandApply(in types.ColorDerect) types.ColorDerect {
	AbsI := IntAbs(in.I)
	AbsJ := IntAbs(in.J)
	Absk := IntAbs(in.K)
	var m map[int]int = map[int]int{AbsI: in.I, AbsJ: in.J, Absk: in.K}
	list := []int{AbsI, AbsJ, Absk}
	sort.Ints(list)
	return types.ColorDerect{m[list[0]], m[list[1]], m[list[2]]}
}

//colorDirect rotate 90
// mirror  rotate90 mirror  righthand
func ColorXR90(in types.ColorDerect) types.ColorDerect {
	out, mirror := Mirror(in)

	b := mat.NewDense(3, 1, []float64{
		float64(out.I), float64(out.J), float64(out.K),
	})
	// Take the matrix product of a and b and place the result in c.
	var c mat.Dense
	c.Mul(types.XR90, b)
	remirror := ReMirror(types.ColorDerect{int(c.At(0, 0)), int(c.At(1, 0)), int(c.At(2, 0))}, mirror)
	return RightHandApply(remirror)
}

func PositionXR90(in types.Position) types.Position {
	position := mat.NewDense(3, 1, []float64{
		float64(in.H), float64(in.K), float64(in.L),
	})
	// Take the matrix product of a and b and place the result in c.
	var c mat.Dense
	c.Inverse(types.XR90)

	var d mat.Dense
	d.Mul(&c, position)
	return types.Position{int(d.At(0, 0)), int(d.At(1, 0)), int(d.At(2, 0))}
}

func IntAbs(i int) int {
	if i >= 0 {
		return i
	} else {
		return -i
	}
}
