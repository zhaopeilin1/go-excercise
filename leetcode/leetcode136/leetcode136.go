package leetcode136

func singleNumber(nums []int) int {
	//找出只出现一次的元素，其余的出现2次
	num := 0
	for _, v := range nums {
		num = num ^ v
	}
	return num
}
