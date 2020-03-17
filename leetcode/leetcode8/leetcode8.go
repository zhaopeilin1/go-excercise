package leetcode8

import (
	"errors"
	"fmt"
)

var (
	flag   map[rune]int = map[rune]int{'+': 1, '-': -1}
	numMap map[rune]int = map[rune]int{'0': 0, '1': 1,
		'2': 2,
		'3': 3,
		'4': 4,
		'5': 5,
		'6': 6,
		'7': 7,
		'8': 8,
		'9': 9,
	}

	blank rune = ' '
	dot   rune = '.'

	max int = 1<<31 - 1
	min int = -(1 << 31)
)

func myAtoi(str string) int {
	//找到第一个不是空格的字符，看是否正负号或数字。如果是正负号，截取其后续最大长度的数字。
	//0开头会怎么样？
	strs := []rune(str)

	strs = trimLeft(strs)

	if len(strs) == 0 {
		return 0
	}
	identity, exists := flag[strs[0]]
	if !exists {
		identity = 1
	} else {
		//去掉符号
		strs = strs[1:]
	}
	//fmt.Println(string(strs))
	//处理后面的空格,截断
	strs = cutByBlank(strs)
	//fmt.Println("cut", string(strs))
	if len(strs) == 0 {
		return 0
	} else {
		//第一个字符，+-符号。

		n, err := parse(strs)
		if err != nil {
			return 0
		} else {
			//判断是否超界
			r := n * identity
			if r > max {
				return max
			} else if r < min {
				return min
			} else {
				return r
			}
		}
	}
	return 0
}

func isOutOfRange(a int) bool {
	return a < min || a > max
}

func parse(list []rune) (int, error) {
	acc := 0
	for _, ru := range list {
		//注意0开头的数字
		oneNum, exists := numMap[ru]
		if !exists {
			return 0, errors.New(fmt.Sprintf("not a num %v", ru))
		} else {
			// if i == 1 && acc == 0 {
			// 	return 0, errors.New(fmt.Sprintf("start with 0 %v", list))
			// }
			acc = acc*10 + oneNum
		}
		if isOutOfRange(acc) {
			return acc, nil
		}
	}
	return acc, nil
}

func cutByBlank(str []rune) []rune {
	//用后面的空格分割
	if len(str) == 0 {
		return str
	} else {
		indexBlank := 0
		for index, ru := range str {
			_, exists := numMap[ru]
			if !exists {
				indexBlank = index
				return str[0:indexBlank]
			}

		}
		return str
	}
}

func trimLeft(str []rune) []rune {
	//去掉左边空格
	if len(str) == 0 {
		return str
	} else {
		if str[0] == blank {
			return trimLeft(str[1:])
		} else {
			return str
		}
	}
}
