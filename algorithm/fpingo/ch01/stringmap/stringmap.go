package main

import (
	"fmt"
)

func main() {
	var a string = "abcd"
	var b string = "ad11c"
	r := numJewelsInStones(a, b)
	fmt.Println(r)

}

func numJewelsInStones(J string, S string) int {
	m := make(map[rune]bool)
	strs := []rune(J)
	for _, i := range strs {
		m[i] = true
	}

	r := 0
	owns := []rune(S)
	for _, item := range owns {
		_, ok := m[item]
		if ok {
			r++
		}

	}
	return r

}
