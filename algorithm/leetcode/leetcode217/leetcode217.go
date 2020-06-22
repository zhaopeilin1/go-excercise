package leetcode217

func containsDuplicate(nums []int) bool {
	m1 := make(map[int]bool)
	for _, num := range nums {
		_, exists := m1[num]
		if exists {
			return true
		} else {
			m1[num] = true
		}
	}
	return false
}
