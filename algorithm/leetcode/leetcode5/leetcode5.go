package leetcode5

import "fmt"

func longestPalindrome(s string) string {
	//最长回文子串,用中心扩展法。
	if len(s) == 0 {
		return s
	} else if len(s) == 1 {
		return s
	}
	str1 := addSep(s)
	str := []rune(str1)
	l := len(str)
	maxR := 0
	maxC := 0
	for i := 1; i < l-1; i++ {
		r := checkPalindrome(str, i)
		//fmt.Println(str, i, r)
		//^b#a#b#a#d$
		//这个判断有点问题
		if str[i-r] == '#' {
			r = r - 1
		}
		if r > maxR {
			maxR = r
			maxC = i
		}
	}
	var result1 []rune
	if maxR == 0 {
		result1 = str[1:2]
	} else {
		result1 = str[maxC-maxR : maxC+maxR+1]
	}

	//fmt.Println(result1, maxC, maxR)
	//去掉分隔符
	result := []rune{}
	for _, ru := range result1 {
		if ru != '#' && ru != '^' && ru != '$' {
			result = append(result, ru)
		}
	}
	return string(result)
}

func checkPalindrome(str []rune, center int) int {
	//检查中心是否回文，回文半径是多大。
	i := 0
	for ; ; i = i + 1 {
		//fmt.Println(str, center, i)
		if str[center-i] == str[center+i] {
		} else {
			i = i - 1
			break
		}
	}
	return i
}

func addSep(s string) string {
	list0 := []rune(s)
	l := len(s)
	l2 := 2*l + 1
	list := make([]rune, l2)
	list[0] = rune('^')
	list[l2-1] = rune('$')
	for i := 1; i < l2-1; i++ {
		if i%2 == 1 {
			list[i] = list0[i/2]
		} else {
			list[i] = rune('#')
		}
	}
	return string(list)
}
