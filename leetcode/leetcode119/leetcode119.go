package leetcode119

func getRow(rowIndex int) []int {
	r := make([]int, rowIndex+1)
	for i := 0; i <= rowIndex; i++ {
		r[i] = cc(rowIndex, i)
	}
	return r
}

func cc(n, m int) int {
	//n>=m
	return (acc(n) / acc(m)) / acc(n-m)
}

func acc(m int) int {
	if m == 0 {
		return 1
	}
	s := 1
	for i := 1; i <= m; i++ {
		s = s * i
	}
	return s
}
