package leetcode198

import (
	"fmt"
)

func rob(nums []int) int {
	var m1 map[int]int = make(map[int]int)
	return rob2(nums, 0, m1)
}

func rob2(nums []int, index int, m1 map[int]int) int {
	l := len(nums)
	if l == 0 {
		return 0
	} else if l == 1 {
		m1[index] = nums[0]
		return nums[0]
	} else if l == 2 {
		r := max(nums[0], nums[1])
		m1[index] = r
		return r
	} else if l == 3 {
		sum1 := nums[0] + nums[2]
		sum2 := nums[1]
		r := max(sum1, sum2)
		m1[index] = r
		return r
	} else {
		cache1, exists1 := m1[2+index]
		if !exists1 {
			cache1 = rob2(nums[2:], 2+index, m1)
			m1[2+index] = cache1
		} else {
			fmt.Println("cache", 2+index)
		}
		cache2, exists2 := m1[3+index]
		if !exists2 {
			cache2 = rob2(nums[3:], 3+index, m1)
			m1[3+index] = cache2
		} else {
			fmt.Println("cache", 3+index)
		}
		sum1 := nums[0] + cache1
		sum2 := nums[1] + cache2
		return max(sum1, sum2)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
