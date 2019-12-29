package leetcode66

func plusOne(digits []int) []int {
	return plusOne2(digits, len(digits)-1)
}

func plusOne2(digits []int, index int) []int {
	//最后一位加1，如果最后一位是9，那么需要进位。进位后继续以上步骤
	if index == -1 {
		result := make([]int, len(digits)+1)
		result[0] = 1
		for i, v := range digits {
			result[i+1] = v
		}
		return result
	} else {
		sum := digits[index] + 1
		if sum == 10 {
			digits[index] = 0
			return plusOne2(digits, index-1)
		} else {
			digits[index] = sum
			return digits
		}
	}
}
