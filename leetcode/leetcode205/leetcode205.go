package leetcode205

func isIsomorphic(s string, t string) bool {
	//aaa bbb
	list1 := []byte(s)
	list2 := []byte(t)
	l := len(list1)
	m1 := make(map[byte]int)
	m2 := make(map[byte]int)
	for i := 0; i < l; i++ {
		b1 := list1[i]
		b2 := list2[i]
		count1, exists1 := m1[b1]
		count2, exists2 := m2[b2]
		if exists2 != exists1 || (exists1 == true && count1 != count2) {
			return false
		}
		if exists1 {
			m1[b1] = count1 + i
		} else {
			m1[b1] = i
		}

		if exists2 {
			m2[b2] = count2 + i
		} else {
			m2[b2] = i
		}

	}
	return true
}
