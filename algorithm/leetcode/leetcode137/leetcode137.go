package leetcode137

import (
	"fmt"
)

func singleNumber(nums []int) int {
	sum := 0
	xoSum := 0
	//a,a,a,b a^b
	for _, num := range nums {
		sum = sum + num
		xoSum = xoSum ^ num
	}
	fmt.Println(sum, xoSum)

	fmt.Println(sum ^ xoSum)
	return (sum - ((sum - xoSum) / 2 * 3))
}
