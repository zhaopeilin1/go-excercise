package leetcode299

import "strconv"

func getHint(secret string, guess string) string {
	//位置和数字都对，给一个A，位置不对数字对，给一个B
	list1 := []byte(secret)
	list2 := []byte(guess)
	sumA := 0
	sumB := 0
	m1 := make(map[byte]int, len(list1))
	for i := 0; i < len(list1); i++ {
		if list1[i] == list2[i] {
			sumA += 1
		} else {
			s, exists := m1[list1[i]]
			if exists {
				m1[list1[i]] = s + 1
			} else {
				m1[list1[i]] = 1
			}
		}
	}
	for i := 0; i < len(list1); i++ {
		if list1[i] == list2[i] {
		} else {
			s, exists := m1[list2[i]]
			if exists {
				sumB += 1
				if s == 1 {
					delete(m1, list2[i])
				} else {
					m1[list2[i]] = s - 1
				}
			}
		}
	}
	return strconv.Itoa(sumA) + "A" + strconv.Itoa(sumB) + "B"
}
