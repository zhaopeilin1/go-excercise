package leetcode172

func trailingZeroes(n int) int {
	if n < 5 {
		return 0
	} else {
		return n/5 + trailingZeroes(n/5)
	}
}
