package leetcode20

import (
	"fmt"
)

var m map[string]string = map[string]string{
	"(": ")",
	"{": "}",
	"[": "]",
}

type stack struct {
	list []string
}

func (st *stack) insert(s string) {
	st.list = append(st.list, s)
}

func (st *stack) drop() string {
	if len(st.list) > 0 {
		r := st.list[len(st.list)-1]
		st.list = st.list[:(len(st.list) - 1)]
		return r
	} else {
		return ""
	}
}

func isValid(s string) bool {
	st := &stack{[]string{}}
	for _, str := range s {
		str1 := string(str)
		if str1 == "[" || str1 == "{" || str1 == "(" {
			st.insert(str1)
		}
		if str1 == "]" || str1 == "}" || str1 == ")" {
			drop := st.drop()
			if drop == "" || str1 != m[drop] {
				return false
			}
		}
	}
	if len(st.list) == 0 {
		return true
	} else {
		return false
	}
}
