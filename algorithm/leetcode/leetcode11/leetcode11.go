package leetcode11

func maxArea(height []int) int {
	l := len(height)
	if l < 2 {
		return 0
	} else if l == 2 {
		a, _ := calArea(height[0], height[1], 1)
		return a
	}
	maxArea := 0
	i, j := 0, l-1
	for i < j {
		area, iLessThanJ := calArea(height[i], height[j], j-i)
		maxArea = calMax(maxArea, area)
		if iLessThanJ {
			i++
		} else {
			j--
		}
	}
	return maxArea
}

func calMax(i1, i2 int) int {
	if i1 >= i2 {
		return i1
	} else {
		return i2
	}
}

func calArea(i1, i2, lenth int) (int, bool) {
	if i1 >= i2 {
		return i2 * lenth, false
	} else {
		return i1 * lenth, true
	}
}
