package leetcode26

/*
 */
func removeDuplicates(nums []int) int {
	//每发现一次重复，数目减一，然后后续的位置减一。
	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return 1
	} else {
		sum := 1
		pre := nums[0]
		repeat := 0
		for index, value := range nums[1:] {
			sum++
			if value == pre {
				//发生了重复，未来的所有index都可以减1
				sum--
				repeat++
			} else {
				pre = value
				nums[index+1-repeat] = value
			}
		}
		return sum
	}

}
