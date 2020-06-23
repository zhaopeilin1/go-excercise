package fibo

import "fmt"

//1,1,2,3,5,8...
type Memo func(int) int

func Memoize(f Memo) Memo {
	cache := make(map[int]int)
	return func(n int) int {
		v, ok := cache[n]
		if ok {
			return v
		} else {
			value := f(n)
			cache[n] = value
			return value
		}
	}
}

var fibMemo Memo

func FibMemo(n int) int {
	fibMemo = Memoize(
		func(m int) int {
			if m < 3 {
				return 1
			} else {
				fmt.Println("cache compute:", m)
				return fibMemo(m-1) + fibMemo(m-2)
			}
		})
	return fibMemo(n)
}
