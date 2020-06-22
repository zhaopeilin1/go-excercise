package bisearch

import (
	"fmt"
)

func bisearch(list []int, item int) int {
	return bisearch2(list, item, 0, len(list)-1)
}

func bisearch2(list []int, item int, start, end int) int {
	if start == end && list[start] != item {
		return -1
	}
	mid := (start + end) / 2
	if list[mid] == item {
		return mid
	} else if list[mid] < item {
		return bisearch2(list, item, mid+1, end)
	} else {
		return bisearch2(list, item, start, mid-1)
	}
}

func bisearch3(list []int, item int) int {
	low := 0
	high := len(list) - 1

	for low <= high {
		fmt.Println(low, high)
		fmt.Println(list)
		mid := (low + high) / 2
		guess := list[mid]
		if guess == item {
			return mid
		} else if guess < item {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}
