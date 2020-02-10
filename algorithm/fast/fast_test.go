package fast

import (
	"reflect"
	"testing"
)

var data = []struct {
	in1 int
	in2 int
	out int
}{
	{3, 3, 27},
	//{3, 15, 14348907},
}

var matrxs = []struct {
	in1 Matrix
	in2 Matrix
	out Matrix
}{
	{Matrix{[][]int{[]int{1, 0}, []int{0, 1}}}, Matrix{[][]int{[]int{1, 0}, []int{0, 1}}}, Matrix{[][]int{[]int{2, 0}, []int{0, 2}}}},
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

func TestMatrixAdd(t *testing.T) {
	for _, d := range matrxs {
		r := d.in1.Add(d.in2)
		//t.Log(r)
		//t.Log(d.in1, d.in2, r)
		d.in1.Print()
		d.in2.Print()
		r.Print()
		if !reflect.DeepEqual(r, d.out) {
			t.Error(d.out, r)
		}
	}
}

func TestMatrixNew(t *testing.T) {
	r := NewZeroMatrix(3, 4)
	r.Print()
}

func TestMatrixMul(t *testing.T) {
	t.Run("矩阵乘法", func(t *testing.T) {
		a := Matrix{[][]int{[]int{1, 2, 1}, []int{2, 4, 1}}}
		b := Matrix{[][]int{[]int{1, 3}, []int{2, 1}, []int{3, 2}}}
		c := Matrix{[][]int{[]int{2, 0}, []int{0, 2}}}
		r := a.Mul(b)
		a.Print()
		b.Print()
		r.Print()
		c.Print()
	})
}

func TestMatrixPow(t *testing.T) {
	t.Run("矩阵幂", func(t *testing.T) {
		a := Matrix{[][]int{[]int{2, 0}, []int{0, 2}}}
		r := a.Pow(12)
		a.Print()
		r.Print()
	})
}

func TestMatrixFib(t *testing.T) {
	t.Run("斐波那契", func(t *testing.T) {
		a := Matrix{[][]int{[]int{0, 1}, []int{1, 1}}}
		r := a.Pow(19)
		//0,1,1,2,3,5。 如果是求第n 项，则先计算其n-1次幂，再乘以头两项。则向量的最后一个分量就是第 n项。
		fib := r.Mul(Matrix{[][]int{[]int{0}, []int{1}}})
		r.Print()
		fib.Print()
	})
}
