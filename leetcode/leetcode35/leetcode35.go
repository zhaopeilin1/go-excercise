package leetcode35

func searchInsert(nums []int, target int) int {
	return searchInsert2(nums, target, 0, len(nums))
}

func searchInsert2(nums []int, target int, start int, end int) int {
	//如果小于最小的，就0，大于最大的，就在最大的index+1
	//否则大于值a,小于等于值b,就a的index +1
	//l := len(nums)
	if target <= nums[start] {
		return start
	} else if target > nums[end-1] {
		return end
	} else {
		a := (start + end) / 2
		//0,a)[a,)
		if target <= nums[a] {
			return searchInsert2(nums, target, start, a)
		} else {
			return searchInsert2(nums, target, a, end)
		}
	}
}
