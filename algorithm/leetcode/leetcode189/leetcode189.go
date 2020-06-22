package leetcode189

import (
	"fmt"
)

func rotate(nums []int, k int) {
	l := len(nums)
	tmp := nums[0]
	start := 0
	j := 0
	i := 0
	k = k % l
	for i < l-1 {
		//fmt.Println(j, (j+l-k)%l)
		if ((j + l - k) % l) == start {
			nums[j] = tmp
			j = (j + 1) % l
			tmp = nums[j]
			start = j
		} else {
			nums[j] = nums[(j+l-k)%l]
			j = (j + l - k) % l
		}
		i++
		//fmt.Println(nums)
	}
	nums[j] = tmp
	//fmt.Println(nums)
}
