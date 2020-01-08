package leetcode389

func findTheDifference(s string, t string) byte {
	//s<t
	var acc byte = 0
	for _, b := range []byte(s) {
		acc = acc ^ b
	}

	for _, b := range []byte(t) {
		acc = acc ^ b
	}
	return acc
}
