package leetcode371

func getSum(a int, b int) int {
	if a == 0 {
		return b
	} else if a > 0 {
		for a > 0 {
			a--
			b++
		}
		return b

	} else {
		for a < 0 {
			a++
			b--
		}
		return b

	}
}
