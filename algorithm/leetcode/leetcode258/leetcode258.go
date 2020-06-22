package leetcode258

//	"fmt"

func addDigits(num int) int {
	acc := 0
	for num > 0 {
		acc = (acc+(num%10))%10 + (acc+(num%10))/10
		num = num / 10
	}
	return acc
}
