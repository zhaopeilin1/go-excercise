package leetcode367

func isPerfectSquare(num int) bool {
	return isSquare(0, num, num)
}

func isSquare(low, high, num int) bool {
	if low > high {
		return false
	} else if low == high {
		return low*low == num
	} else {
		mid := (low + high) / 2
		guess := mid * mid
		if guess < num {
			return isSquare(mid+1, high, num)
		} else if guess == num {
			return true
		} else {
			return isSquare(low, mid-1, num)
		}
	}
}
