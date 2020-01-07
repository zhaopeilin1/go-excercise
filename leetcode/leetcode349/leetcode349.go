package leetcode349

func intersection(nums1 []int, nums2 []int) []int {
	m1 := make(map[int]bool, len(nums1))
	for _, num := range nums1 {
		m1[num] = true
	}
	m2 := make(map[int]bool, len(nums2))
	for _, num := range nums2 {
		m2[num] = true
	}
	result := []int{}
	for k, _ := range m1 {
		_, exists := m2[k]
		if exists {
			result = append(result, k)
		}
	}
	return result
}
