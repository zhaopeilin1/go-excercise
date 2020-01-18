package ch2

import (
	"fmt"
)

func genStrBy10(n int) {
	genStrBy10Help(make([]int, n), n)
}

func genStrBy10Help(list []int, n int) {
	if n < 1 {
		fmt.Println(list)
	} else {
		for i := 1; i <= 4; i++ {
			list[n-1] = i
			genStrBy10Help(list, n-1)
		}
	}
}
