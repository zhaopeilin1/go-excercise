package leetcode53

import (
	"fmt"
)

func maxSubArray(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	} else if l == 1 {
		return nums[0]
	} else {
		sigleMax := nums[0]
		groupList := []int{}
		groupItem := nums[0]
		for index, num := range nums[1:l] {

			sigleMax = maxInt(sigleMax, num)
			if groupItem <= 0 {
				if num <= 0 {
					groupItem = groupItem + num
				} else {
					//增加一个分组项
					groupList = append(groupList, groupItem)
					groupItem = num
				}
			} else {
				if num >= 0 {
					groupItem = groupItem + num
				} else {
					//增加一个分组项
					groupList = append(groupList, groupItem)
					groupItem = num
				}
			}
			if index == l-2 {
				groupList = append(groupList, groupItem)
			}
		}
		fmt.Println(groupList)
		return maxInt(sigleMax, maxSubArray2(groupList))
	}
}

func maxSubArray2(nums []int) int {
	//合并之后的，[正，负]交替的数组，求出其单个最大，组合最大
	l := len(nums)
	sigleMax := nums[0]
	groupMax := nums[0]
	for _, num := range nums[1:l] {
		sigleMax = maxInt(sigleMax, num)
		acc := groupMax + num
		if num >= 0 {
			if groupMax > 0 {
				groupMax = acc
				sigleMax = maxInt(groupMax, sigleMax)
			} else {
				sigleMax = maxInt(groupMax, sigleMax)
				groupMax = num
			}
		} else {
			if acc > 0 {
				groupMax = acc
				sigleMax = maxInt(groupMax, sigleMax)
			} else {
				sigleMax = maxInt(groupMax, sigleMax)
				groupMax = num
			}
		}
	}
	sigleMax = maxInt(groupMax, sigleMax)
	return sigleMax
}

func maxInt(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}
