package leetcode3

//"fmt"

func lengthOfLongestSubstring(s string) int {
	//最长不重复子串，
	//从左开始，遇到不重复的，子串加1，
	//如果遇到和已有子串的，找到和已有的重复的位置，从那个位置+1开始计算不重复子串。
	m := map[rune]int{}
	m2 := map[int]rune{}
	//value index
	uniq := 0
	maxUniq := 0
	for index, value := range s {
		m2[index] = value
		oldIndex, exist := m[value]
		if !exist {
			m[value] = index
			uniq++
		} else {
			//fmt.Println(value, oldIndex)
			maxUniq = max(uniq, maxUniq)
			//uniq -> -n
			m = dropSome(m, oldIndex)
			//fmt.Println(len(m))
			m[value] = index
			uniq = len(m)
			//m drop old, insert new
		}
		//fmt.Println(uniq)
		//fmt.Println(m)
	}
	return max(uniq, maxUniq)
}
func dropSome(m1 map[rune]int, index int) map[rune]int {
	for k, v := range m1 {
		if v <= index {
			delete(m1, k)
		}
	}
	//fmt.Println(m1)
	return m1
}
func max(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}
