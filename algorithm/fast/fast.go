package fast

import (
	"fmt"
)

//"fmt"
//"reflect"

type Matrix struct {
	val [][]int
}

func NewZeroMatrix(row, col int) Matrix {
	//零矩阵
	val := make([][]int, row)
	for i := 0; i < row; i++ {
		val[i] = make([]int, col)
	}
	return Matrix{val}
}

func NewOneMatrix(n int) Matrix {
	//单位矩阵
	m := NewZeroMatrix(n, n)
	for i := 0; i < n; i++ {
		m.val[i][i] = 1
	}
	return m
}

func (m Matrix) Mul(a Matrix) Matrix {
	//矩阵相乘
	//return Matrix{}
	r := NewZeroMatrix(len(m.val), len(a.val[0]))
	//reflect.Copy(r, m.val)
	for row, list := range r.val {
		for col, _ := range list {
			// r: row,col = sum(m[row] * a[col])
			sum := 0
			for i := 0; i < len(m.val[0]); i++ {
				sum += m.val[row][i] * a.val[i][col]
			}
			list[col] = sum
		}
	}
	return r
}

func (m Matrix) Add(a Matrix) Matrix {
	//矩阵相加
	r := [][]int{}
	//reflect.Copy(r, m.val)
	for row, list := range m.val {
		r = append(r, []int{})
		for col, v := range list {
			r[row] = append(r[row], v+a.val[row][col])
			//list[col] = v + a[row][col]
		}
	}
	return Matrix{r}
}

func (base Matrix) Pow(n int) Matrix {
	//矩阵快速幂
	ans := NewOneMatrix(len(base.val))
	for n > 0 {
		if n&1 == 1 {
			ans = ans.Mul(base)
		}
		base = base.Mul(base)
		n = n >> 1
	}
	return ans
}
func (m Matrix) Print() {
	fmt.Println(len(m.val), "X", len(m.val[0]), "matrix:")
	for _, list := range m.val {
		fmt.Println(list)
	}
}

//快速幂，快速乘
func pow(base, n int) int {
	// base^n , mod
	ans := 1
	for n > 0 {
		//fmt.Println("base:", base, "n:", n, "anser:", ans)
		//如果n最后一位是1，就乘上一次自己
		if n&1 == 1 {
			ans *= base
			//fmt.Println("n is old anser:", ans)
		}
		//直接平方
		base *= base
		n = n >> 1
	}
	return ans
}
