package leetcode7

func reverse(x int) int {
	// [−231,  231 − 1]
	r := reverse2(0, x)
	max := 1<<31 - 1
	min := -(1 << 31)
	//fmt.Println(max, min)
	if r >= min && r <= max {
		return r
	} else {
		return 0
	}
}

func reverse2(sum, x int) int {
	if x != 0 {
		return reverse2(sum*10+x%10, x/10)
	} else {
		return sum
	}
}
