package leetcode345

import (
	"fmt"
)

//const sets []byte = []byte("aeuoi")

func reverseVowels(s string) string {
	m := getMap([]byte("aeiouAEIOU"))
	list := []byte(s)
	for i, j := 0, len(list)-1; i < j; {
		for ; !isVowel(list[i], m) && i < j; i++ {
		}
		for ; !isVowel(list[j], m) && i < j; j-- {
		}
		list[i], list[j] = list[j], list[i]
		i++
		j--
	}
	return string(list)
}

func getMap(list []byte) map[byte]bool {
	m := make(map[byte]bool)
	for _, li := range list {
		m[li] = true
	}
	return m
}

func isVowel(b byte, m map[byte]bool) bool {
	_, isExists := m[b]
	return isExists
}
