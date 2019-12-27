package leetcode28

func strStr(haystack string, needle string) int {
	//顺序查询，复杂度问题
	if needle == "" {
		return 0
	}
	a := []rune(haystack)
	b := []rune(needle)
	l := len(b)

	for i := 0; i <= len(a)-l; i++ {
		if equals(a[i:i+l], b) {
			return i
		}
	}
	return -1
}

func equals(a, b []rune) bool {
	if len(a) == len(b) {
		for index, r := range a {
			if b[index] != r {
				return false
			}
		}
		return true

	} else {
		return false
	}

}
