package leetcode27

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	} else {
		found := 0
		for index, value := range nums {
			if value == val {
				found++
			} else {
				if found > 0 {
					nums[index-found] = value
				}
			}
		}
		return len(nums) - found
	}

}
func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	} else {
		found := 0
		for index, value := range nums {
			if value == val {
				found++
			} else {
				if found > 0 {
					nums[index-found] = value
				}
			}
		}
		return len(nums) - found
	}

}
