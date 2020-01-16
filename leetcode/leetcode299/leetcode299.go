package leetcode299

func getHint(secret string, guess string) string {
	//位置和数字都对，给一个A，位置不对数字对，给一个B
	list1 := []byte(secret)
	list2 := []byte(guess)

	for i := 0; i < len(list1); i++ {
		if list1[i]==list2[i]

	}
}
