package leetcode67

//"fmt"

func addBinary(a string, b string) string {
	//二进制加法 101 + 101 = 1010

	var max int

	if len(a) > len(b) {
		max = len(a) + 1
	} else {
		max = len(b) + 1
		a, b = b, a
	}
	l1 := len(a)
	l2 := len(b)
	al := []rune(a)
	bl := []rune(b)
	var jinwei string = "0"
	list := make([]rune, max)
	//fmt.Println(max, l1, l2)
	var r string
	for i := l1 - 1; i >= 0; i-- {
		//{"101", "101", "1010"},
		if i-l1+l2 <= l2-1 && i-l1+l2 >= 0 {
			r, jinwei = add(string(al[i]), string(bl[i-l1+l2]), jinwei)
			//fmt.Println(r, jinwei)
			list[i+1] = []rune(r)[0]
			//fmt.Println(string(list))
		} else {
			r, jinwei = add(string(al[i]), "0", jinwei)
			//fmt.Println(r, jinwei)
			list[i+1] = []rune(r)[0]
			//fmt.Println(string(list))
		}
	}
	if jinwei == "1" {
		list[0] = []rune("1")[0]
	}

	if string(list[0]) == "1" {
		return string(list)
	} else {
		return string(list[1:])
	}
}

func add(a, b, c string) (string, string) {
	//fmt.Println(a, b, c)
	//返回本位，进位
	if a == "1" && b == "1" {
		return c, "1"
	} else if a == "1" && b == "0" {
		if c == "1" {
			return "0", "1"
		} else {
			return "1", "0"
		}
	} else if a == "0" && b == "1" {
		if c == "1" {
			return "0", "1"
		} else {
			return "1", "0"
		}
	} else {
		return c, "0"
	}
}
