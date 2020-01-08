package leetcode383

func canConstruct(ransomNote string, magazine string) bool {
	m := make(map[rune]int)
	for _, r := range magazine {
		count, exists := m[r]
		if exists {
			m[r] = count + 1
		} else {
			m[r] = 1
		}
	}
	for _, r := range ransomNote {
		count, exists := m[r]
		if exists {
			if count == 1 {
				delete(m, r)
			} else {
				m[r] = count - 1
			}
		} else {
			return false
		}
	}
	return true
}
