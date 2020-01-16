package leetcode283

func moveZeroes(nums []int) {
	l := len(nums)
	count := 0
	for index, num := range nums {
		if num == 0 {
			count++
		} else {
			if count > 0 {
				nums[index-count] = nums[index]
				nums[index] = 0
			}
			if index < l-count {
			} else {
				nums[index] = 0
			}
		}
	}
}
