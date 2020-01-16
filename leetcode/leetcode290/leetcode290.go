package leetcode290

import "strings"

func wordPattern(pattern string, str string) bool {
	words := strings.Split(str, " ")
	bytes := []byte(pattern)
	m1 := make(map[byte]int)
	m2 := make(map[string]int)
	l := len(bytes)
	if l != len(words) {
		return false
	}
	for i := 0; i < l; i++ {
		sum1, _ := m1[bytes[i]]
		m1[bytes[i]] = sum1 + i + 1
		sum2, _ := m2[words[i]]
		m2[words[i]] = sum2 + i + 1
		if m1[bytes[i]] != m2[words[i]] {
			return false
		}
	}
	return true
}
