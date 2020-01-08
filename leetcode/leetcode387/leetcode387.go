package leetcode387

func firstUniqChar(s string) int {
	m1 := make(map[rune][]int)
	for index, ru := range s {
		indexRange, exists := m1[ru]
		if exists {
			m1[ru] = append(indexRange, index)
		} else {
			m1[ru] = []int{index}
		}
	}
	min := -1
	for _, v := range m1 {
		if len(v) == 1 {
			if min == -1 {
				min = v[0]
			} else {
				if min > v[0] {
					min = v[0]
				}
			}
		}
	}
	return min
}
