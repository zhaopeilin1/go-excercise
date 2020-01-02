package leetcode118

func generate(numRows int) [][]int {
	//给定一个非负整数 numRows，生成杨辉三角的前 numRows 行。
	r := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		row := getRow(i)
		r[i] = row
	}
	return r
}
func getRow(rowIndex int) []int {
	r := make([]int, rowIndex+1)
	for i := 0; i <= rowIndex; i++ {
		r[i] = cc(rowIndex, i)
	}
	return r
}

func cc(n, m int) int {
	if n == m {
		return 1
	} else {
		m2 := n - m
		var min int
		if m2 < m {
			min = m2
		} else {
			min = m
		}

		acc1 := 1
		acc2 := 1

		for i := 1; i <= min; i++ {
			acc1 = acc1 * (n + 1 - i)
			acc2 = acc2 * i
			if acc2 > 1 && acc1%acc2 == 0 {
				acc1 = acc1 / acc2
				acc2 = 1
			}
		}
		return acc1 / acc2
	}
}
