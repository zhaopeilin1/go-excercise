package leetcode70

func climbStairs(n int) int {
	//每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
	return climbStairs2(n)
}

func climbStairs2(n int) int {
	if n <= 1 {
		return 1
	} else {
		return climbStairs2(n-1) + climbStairs2(n-2)
	}
}

func climbStairs3(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		pre1 := 1
		pre2 := 1
		sum := 0
		for i := 2; i <= n; i++ {
			sum = pre1 + pre2
			pre2 = pre1
			pre1 = sum

		}

		return sum
	}

}
