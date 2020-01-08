package leetcodelcp1

func game(guess []int, answer []int) int {
	acc := 0
	for index, r := range guess {
		if answer[index] == r {
			acc++
		}
	}
	return acc
}
