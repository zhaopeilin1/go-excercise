package leetcode412

import "strconv"

func fizzBuzz(n int) []string {
	list := make([]string, n)
	for i := 1; i <= n; i++ {
		s := ""
		if i%3 == 0 {
			s += "Fizz"
		}
		if i%5 == 0 {
			s += "Buzz"
		}
		if s == "" {
			list[i-1] = strconv.Itoa(i)
		} else {
			list[i-1] = s
		}
	}
	return list
}
