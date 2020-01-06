package leetcode168

import (
	"fmt"
)

var m1 map[int]string = map[int]string{
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
	str := ""
	weishu, zhi := convertToTitle2(n)
	str = str + m1[zhi]
	for weishu > 1 {
		n = n - pow(26, weishu-1)
		weishu, zhi = convertToTitle2(n)
		str = str + m1[zhi]
	}
	str = str + m1[n]
	return str
}

func pow(m, n int) int {
	sum := 1
	for i := 0; i < n; i++ {
		sum = sum * m
	}
	return sum
}

func convertToTitle2(n int) (int, int) {
	m := n / 26
	weishu := 1
	for m > 26 {
		m = m / 26
		weishu = weishu + 1
	}
	fmt.Println(n, weishu, m)
	return weishu, m
}
