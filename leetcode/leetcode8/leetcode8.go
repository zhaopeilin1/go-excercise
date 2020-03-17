package leetcode8

const (
	flag map[rune]bool = map[rune]bool{
		{'+', true},
		{'-', true},
	}
	num map[rune]bool = map[rune]bool{
		{'0', true},
		{'1', true},
		{'2', true},
		{'3', true},
		{'4', true},
		{'5', true},
		{'6', true},
		{'7', true},
		{'8', true},
		{'9', true},
	}

	blank rune = ' '
)

func myAtoi(str string) int {
	//找到第一个不是空格的字符，看是否正负号或数字。如果是正负号，截取其后续最大长度的数字。
	//0开头会怎么样？
	str = trimLeft([]rune(str))

	str = cutByBlank(str)

	isNegate := good[0] == '-'

	preIsBlank := true
	for i, ru := range str {
		if i == ' ' {

		} else {
			preIsBlank = false
		}

	}
	return 0
}

func cutByBlank(str []rune) []rune {
	//用后面的空格分割
	if len(str) == 0 {
		return str
	} else {
		indexBlank = 0
		for index, ru := range str {
			if ru == blank {
				indexBlank = index
				break
			}
		}
		return str[0:indexBlank]
	}

}

func trimLeft(str []rune) []rune {
	//去掉左边空格
	if len(str) == 0 {
		return str
	} else {
		if str[0] == blank {
			return trim(str[1:])
		} else {
			return str
		}
	}
}
