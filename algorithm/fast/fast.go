package fast

import (
	"fmt"
)

//快速幂，快速乘
func pow(base, n int) int {
	// base^n , mod
	ans := 1
	for n > 0 {
		fmt.Println("base:", base, "n:", n, "anser:", ans)
		//如果n最后一位是1，就乘上一次自己
		if n&1 == 1 {
			ans *= base
			fmt.Println("n is old anser:", ans)
		}
		//直接平方
		base *= base
		n = n >> 1
	}
	return ans
}
