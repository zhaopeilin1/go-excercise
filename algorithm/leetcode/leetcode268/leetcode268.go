package leetcode268

func missingNumber(nums []int) int {
	sum := 0
	max := 0
	var has0 bool
	for _, num := range nums {
		if num == 0 {
			has0 = true
		}
		sum = sum + num
		max = getmax(max, num)
	}
	sum2 := max * (max + 1) / 2
	if has0 {
		if sum == sum2 {
			return max + 1
		} else {
			return sum2 - sum
		}
	} else {
		return 0
	}
}
func getmax(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
