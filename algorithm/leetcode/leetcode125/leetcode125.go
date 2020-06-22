package leetcode125

//"fmt"

func isPalindrome(s string) bool {
	//{"azAZ019", false}, //97,122,65,90, 48,49,57
	//TODO
	list := []byte(s)
	for i, j := 0, len(list)-1; i < j; {
		for ; !isCaseOrNumber(list[i]) && i < j; i++ {

		}
		for ; !isCaseOrNumber(list[j]) && i < j; j-- {

		}
		if !isSameNumOrCase(list[i], list[j]) {
			return false
		} else {
			i++
			j--
		}
	}
	return true
}
func isCaseOrNumber(b byte) bool {
	r := (b >= 48 && b <= 57) || (b >= 65 && b <= 90) || (b >= 97 && b <= 122)
	return r
}

func isSameNumOrCase(a, b byte) bool {
	return getBase(a) == getBase(b)
}

func getBase(b byte) byte {
	if b >= 97 {
		return b - 32
	} else {
		return b
	}
}
