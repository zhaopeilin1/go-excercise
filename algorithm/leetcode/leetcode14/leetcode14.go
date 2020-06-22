package leetcode14

//	"strings"

//最长公共前缀
func longestCommonPrefix(strs []string) string {
	if strs == nil || len(strs) == 0 {
		return ""
	} else if len(strs) == 1 {
		return strs[0]
	} else {
		str := strs[0]
		rest := strs[1:]
		result := make([]byte, 0)
		for index, b := range []byte(str) {
			if findByte(index, b, rest) {
				result = append(result, b)
			} else {
				if len(result) == 0 {
					return ""
				} else {
					return string(result)
				}
			}
		}
		if len(result) == 0 {
			return ""
		} else {
			return string(result)
		}
	}
}

func findByte(index int, b byte, strs []string) bool {
	for _, str := range strs {
		if !(len(str) > index && []byte(str)[index] == b) {
			return false
		}
	}
	return true
}
