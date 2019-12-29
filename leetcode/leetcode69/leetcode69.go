package leetcode69

import (
	//"fmt"
	"strconv"
)

func mySqrt(x int) int {
	//计算并返回 x 的平方根，其中 x 是非负整数。
	//由于返回类型是整数，结果只保留整数的部分，小数部分将被舍去。
	//求y*y<=x, (y+1)^2 > x
	//折半搜索法,近似搜索法？
	if x == 0 {
		return 0
	}
	str := strconv.Itoa(x)
	l := len(str)
	//1->1 2->1 3->2 4->2 5->3
	guessL := (l + 1) >> 1
	//fmt.Println(guessL)
	low, high := genRange(guessL)
	//fmt.Println(low, high)
	r := biSearch(x, low, high)
	//fmt.Println(r)
	return r
}

func genRange(l int) (int, int) {
	lowStr := "1"
	highStr := "9"
	for i := 0; i < l-1; i++ {
		lowStr = lowStr + "0"
		highStr = highStr + "9"
	}
	high, _ := strconv.Atoi(highStr)
	low, _ := strconv.Atoi(lowStr)
	return low, high
}

func biSearch(sqrt, low, high int) int {
	//fmt.Println(low, high)
	if low == high {
		return low
	}
	guess := (low + high) / 2
	r := guess * guess
	if r == sqrt {
		return guess
	} else if r > sqrt {
		return biSearch(sqrt, low, guess)
	} else {
		if (guess+1)*(guess+1) > sqrt {
			return guess
		} else {
			return biSearch(sqrt, guess, high)
		}
	}
}
