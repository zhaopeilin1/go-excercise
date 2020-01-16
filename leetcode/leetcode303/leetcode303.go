package leetcode303

type NumArray struct {
	sums []int
}

func Constructor(nums []int) NumArray {
	sums := make([]int, len(nums)+1)
	sum := 0
	for i, num := range nums {
		sum += num
		sums[i+1] = sum
	}
	return NumArray{sums}
}

func (this *NumArray) SumRange(i int, j int) int {	
		return this.sums[j+1] - this.sums[i]	
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */
