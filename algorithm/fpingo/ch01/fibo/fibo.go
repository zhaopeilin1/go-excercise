package fibo

import (
	"fmt"
)

//1,1,2,3,5,8...
func Fib(n int) int {
	if n < 3 {
		return 1
	} else {
		fmt.Println("compute fib :", n)

		return Fib(n-1) + Fib(n-2)
	}

}
