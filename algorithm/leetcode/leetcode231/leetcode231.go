package leetcode231

func isPowerOfTwo(n int) bool {
	acc := 1
	for i := 0; i < 32; i++ {
		if (acc ^ n) == 0 {
			return true
		} else if acc > n {
			return false
		}
		acc = acc << 1
	}
	return false
}
