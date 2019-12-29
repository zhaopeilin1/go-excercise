package leetcode58

//给定一个仅包含大小写字母和空格 ' ' 的字符串，返回其最后一个单词的长度。
//如果不存在最后一个单词，请返回 0 。
func lengthOfLastWord(s string) int {
	l := len(s)
	lenth := 0
	arr := []rune(s)
	for i := l - 1; i >= 0; i-- {
		if string(arr[i]) == " " {
			if lenth > 0 {
				return lenth
			} else {
			}

		} else {
			lenth++
		}
	}
	return lenth
}
