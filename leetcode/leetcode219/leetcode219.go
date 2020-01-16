package leetcode219

func containsNearbyDuplicate(nums []int, k int) bool {
	m1 := make(map[int]int)
	for i, num := range nums {
		pos, exists := m1[num]
		if exists {
			if (i - pos) <= k {
				return true
			} else {
				m1[num] = i
			}
		} else {
			m1[num] = i
		}
	}
	return false
}
