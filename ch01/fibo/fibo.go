package fibo

//1,1,2,3,5,8...
func Fib(n int) int {
	if n < 3 {
		return 1
	} else {
		return Fib(n-1) + Fib(n-2)
	}

}
