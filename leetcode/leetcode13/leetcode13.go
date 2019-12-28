package leetcode13

//"fmt"

/*
{"III", 3},
{"IV", 4},
{"IX", 9},
{"LVIII", 58},
{"MCMXCIV", 1994},
I 可以放在 V (5) 和 X (10) 的左边，来表示 4 和 9。
X 可以放在 L (50) 和 C (100) 的左边，来表示 40 和 90。
C 可以放在 D (500) 和 M (1000) 的左边，来表示 400 和 900

*/

/*
罗马数字包含以下七种字符: I， V， X， L，C，D 和 M。
                       1, 5, 10,50,100,500,1000
*/
func romanToInt(s string) int {
	list := make([]int, len(s))
	for index, str := range s {
		list[index] = map1[string(str)]
	}
	sum := sum(list)
	return sum
}
func sum(list []int) int {
	sum := 0
	pre := 0
	for _, value := range list {
		sum = sum + value
		if pre > 0 && value > pre {
			sum = sum - (2 * pre)
		}
		pre = value

	}
	return sum
}

var map1 map[string]int = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}
