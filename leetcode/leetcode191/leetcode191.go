package leetcode191

func hammingWeight(num uint32) int {
	//101 -> 2
	if num == 0 {
		return 0
	} else {
		return (1 & int(num)) + hammingWeight(num>>1)
	}
}
