package leetcode350

func intersect(nums1 []int, nums2 []int) []int {
	m1 := make(map[int]int, len(nums1))
	for _, num := range nums1 {
		count, exists := m1[num]
		if exists {
			m1[num] = count + 1
		} else {
			m1[num] = 1
		}
	}
	m2 := make(map[int]int, len(nums2))
	for _, num := range nums2 {
		count, exists := m2[num]
		if exists {
			m2[num] = count + 1
		} else {
			m2[num] = 1
		}
	}
	result := []int{}
	for num1, times1 := range m1 {
		times2, exists := m2[num1]
		if exists {
			for times1 > 0 && times2 > 0 {
				result = append(result, num1)
				times1 = times1 - 1
				times2 = times2 - 1
			}
		}
	}
	return result
}
