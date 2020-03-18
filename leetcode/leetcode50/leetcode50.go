package leetcode50

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	} else if n > 0 {
		return calPosity(x, n)
	} else {
		return 1.0 / calPosity(x, -n)
	}
}
func calPosity(x float64, n int) float64 {
	acc := 1.0
	for n > 0 {
		if n&1 == 1 {
			acc *= x
		}
		x *= x
		n = n >> 1
	}
	return acc
}

// for ; position < n; position = position << 2 {
// 	acc = acc * acc
// }
// return acc
