package leetcode392

func isSubsequence(s string, t string) bool {
	//判断s是否t的子字符串
	//hash算法，高效hash算法，假设是0-9的数字
	//先找第一个，找到后再找后一个，直到找到的字符串长度等于s
	list1 := []byte(s)
	list2 := []byte(t)
	l := len(list1)
	if l == 0 {
		return true
	}
	pos := 0
	for _, b := range list2 {
		if pos > l-1 {
			return true
		}
		if b == list1[pos] {
			pos++
		}
	}
	return pos > l-1
}
