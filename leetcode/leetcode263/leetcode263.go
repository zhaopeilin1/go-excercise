package leetcode263

func isUgly(num int) bool {
	//2,3,5
	if num <= 0 {
		return false
	}
	r1 := devideBy(num, 2)
	r2 := devideBy(r1, 3)
	r3 := devideBy(r2, 5)
	return r3 == 1
}

func devideBy(n, m int) int {
	if n == 1 {
		return 1
	} else {
		if n%m == 0 {
			return devideBy(n/m, m)
		} else {
			return n
		}
	}
}
