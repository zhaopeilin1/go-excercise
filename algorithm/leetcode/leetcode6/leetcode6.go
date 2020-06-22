package leetcode6

func convert(s string, numRows int) string {
	//按z字形排列字母，以n*2-2个字母进行循环。
	list := make([][]rune, numRows)
	for i := 0; i < numRows; i++ {
		list[i] = []rune{}
	}
	m := (numRows - 1) * 2
	for index, ru := range s {
		rownum := genOrder(index, numRows, m)
		list[rownum] = append(list[rownum], ru)
	}
	acc := ""
	for i := 0; i < numRows; i++ {
		acc += string(list[i])
	}
	return acc
}

func genOrder(index, n, m int) int {
	//根据n给出下标位置。返回小于n的行数位置
	if n == 1 {
		return 0
	}
	//周期
	//根据不同的余数，安排到不同的行。然后再拼接所有行。
	r := index % m
	//如果 m==0时
	if r >= n {
		return m - r
	} else {
		return r
	}
}
