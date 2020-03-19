package leetcode12

var m map[int]string = map[int]string{
	1:    "I",
	5:    "V",
	10:   "X",
	50:   "L",
	100:  "C",
	500:  "D",
	1000: "M",
}

func intToRoman(num int) string {
	// 1 到 3999
	r := ""
	for position := 1000; position > 0; position = position / 10 {
		r = r + getStringByNum(num/position, m[position], m[position*5], m[position*10])
		num = num % position
	}
	return r
}

func getStringByNum(num int, one, five, ten string) string {
	if num == 9 {
		return one + ten
	} else if num < 9 && num > 4 {
		r := five
		for j := 5; j < num; j++ {
			r = r + one
		}
		return r

	} else if num == 4 {
		return one + five
	} else if num < 4 && num > 0 {
		r := one
		for j := 1; j < num; j++ {
			r = r + one
		}
		return r
	} else {
		return ""
	}
}
