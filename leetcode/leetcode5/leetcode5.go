package leetcode5

func longestPalindrome(s string) string {
	//最长回文子串,用中心扩展法。
	return ""
}

func addSep(s string) string {
	list0 := []rune(s)
	l := len(s)
	l2 := 2*l + 1
	list := make([]rune, l2)
	list[0] = rune('^')
	list[l2-1] = rune('$')
	for i := 1; i < l2-1; i++ {
		if i%2 == 1 {
			list[i] = list0[i/2]
		} else {
			list[i] = rune('#')
		}
	}
	return string(list)
}
