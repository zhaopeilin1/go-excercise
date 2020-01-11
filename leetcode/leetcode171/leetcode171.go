package leetcode171

var m1 map[string]int = map[string]int{
	//0:  "Z",
	"A": 1,
	"B": 2,
	"C": 3,
	"D": 4,
	"E": 5,
	"F": 6,
	"G": 7,
	"H": 8,
	"I": 9,
	"J": 10,
	"K": 11,
	"L": 12,
	"M": 13,
	"N": 14,
	"O": 15,
	"P": 16,
	"Q": 17,
	"R": 18,
	"S": 19,
	"T": 20,
	"U": 21,
	"V": 22,
	"W": 23,
	"X": 24,
	"Y": 25,
	"Z": 26,
}

func titleToNumber(s string) int {
	l := len(s)
	sum := 0
	for i, r := range s {
		sum = sum + m1[string(r)]*pow(26, l-1-i)
	}
	return sum
}

func pow(n, x int) int {
	acc := 1
	for i := 0; i < x; i++ {
		acc = acc * n
	}
	return acc
}
