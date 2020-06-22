package leetcode168

import (
	"fmt"
)

var m1 map[int]string = map[int]string{
	//0:  "Z",
	1:  "A",
	2:  "B",
	3:  "C",
	4:  "D",
	5:  "E",
	6:  "F",
	7:  "G",
	8:  "H",
	9:  "I",
	10: "J",
	11: "K",
	12: "L",
	13: "M",
	14: "N",
	15: "O",
	16: "P",
	17: "Q",
	18: "R",
	19: "S",
	20: "T",
	21: "U",
	22: "V",
	23: "W",
	24: "X",
	25: "Y",
	26: "Z",
}

func convertToTitle(n int) string {
	//从倒数第一位算起，算出余数，如果整体送去余数还有剩，就继续下一步。
	str := ""
	last := 0
	rest := n
	i := 0
	for ; rest > 0 && i < 10; last, rest = calLast(rest) {
		str = m1[last] + str
		fmt.Println(last, rest)
		i++
	}
	str = m1[last] + str
	return str
}

func calLast(n int) (int, int) {
	//计算最后一位，返回计算倒数一位的值。
	if n > 26 {
		if n%26 == 0 {
			return 26, n/26 - 1
		} else {
			return n % 26, n / 26
		}
	} else {
		return n, 0
	}
}
