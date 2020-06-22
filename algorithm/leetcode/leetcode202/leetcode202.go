package leetcode202

import (
	"fmt"
)

func isHappy(n int) bool {
	m := make(map[int]bool)
	for {
		n = transfer(n)
		if n == 1 {
			return true
		}
		_, exists := m[n]
		if exists {
			return false
		} else {
			m[n] = true
		}
	}
}

func transfer(a int) int {
	//fmt.Println(a)
	sum := 0
	for ; a > 0; a = a / 10 {
		rest := a % 10
		sum = sum + rest*rest
	}
	//fmt.Println(sum)
	return sum
}
