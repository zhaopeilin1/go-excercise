package leetcode38

import (
	"strconv"
)

func countAndSay(n int) string {
	//1 11 21 1211 111221
	if n == 1 {
		return "1"
	} else {
		s := countAndSay(n - 1)
		return trans(s)
	}
}

func trans(s string) string {
	accStr := ""
	dupp := ""
	count := 0
	for _, r := range s {
		if dupp != string(r) {
			if count != 0 {
				accStr = accStr + strconv.Itoa(count) + dupp
			}
			count = 1
			dupp = string(r)
		} else {
			count++
		}
	}
	if count != 0 {
		accStr = accStr + strconv.Itoa(count) + dupp
	}
	return accStr
}
