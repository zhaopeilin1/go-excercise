package leetcode645

func findErrorNums(nums []int) []int {
	//返回重复的数字，然后是缺失的数字
	// [2, 10000] 1到 n 的整数
	m := make(map[int]bool)
	l := len(nums)
	var d int
	acc := 0
	sum := l * (l + 1) / 2
	for _, num := range nums {
		acc += num
		_, exists := m[num]
		if exists {
			d = num
		} else {
			m[num] = true
		}
	}
	return []int{d, sum + d - acc}
}
