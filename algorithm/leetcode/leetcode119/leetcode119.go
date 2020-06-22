package leetcode119

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
