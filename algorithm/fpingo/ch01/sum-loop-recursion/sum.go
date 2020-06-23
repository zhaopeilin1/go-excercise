package sum

//循环求和
func SumLoop(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

//递归求和
func SumRecursive(nums []int) int {
	if len(nums) == 0 {
		return 0
	} else {
		return nums[0] + SumRecursive(nums[1:])
	}
}

//尾递归来求和
func SumTailCall(nums []int) int {
	return SumTailCall2(0, nums)
}

//尾递归主体函数
func SumTailCall2(n int, nums []int) int {
	if len(nums) == 0 {
		return n
	} else {
		return SumTailCall2(n+nums[0], nums[1:])
	}
}
