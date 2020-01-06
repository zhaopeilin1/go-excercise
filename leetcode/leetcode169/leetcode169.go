package leetcode169

func majorityElement(nums []int) int {
	l := len(nums)
	m := make(map[int]int)
	for _, num := range nums {
		v, exists := m[num]
		if !exists {
			m[num] = 1
		} else {
			m[num] = v + 1
			if (v + 1) > l/2 {
				return num
			}
		}
	}
	return nums[0]
}
