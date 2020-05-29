package leetcode15

func threeSum(nums []int) [][]int {
	//nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有满足条件且不重复的三元组。
	//-1, 0, 1, 2, -1, -4

	//分治法？ 分成三种，负，0，正
	//做成map[int]int. 数字：出现的次数
	//n平方？
	m := make(map[int]int)
	for index, value := range nums {
		times, exists := m[value]
		if exists {
			m[value] = times + 1
		} else {
			m[value] = 1
		}
	}
}
