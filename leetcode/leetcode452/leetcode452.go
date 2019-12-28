package leetcode452

import (
	"fmt"
	//	"math"
	"sort"
)

func main() {

	fmt.Println("ok")
	a := []int{10, 16}
	b := []int{2, 8}
	c := []int{1, 6}
	d := []int{7, 12}

	var list [][]int
	list = [][]int{a, b, c, d} //[[10,16], [2,8], [1,6], [7,12]];

	fmt.Println(list)
	// sort1(data(list))
	r := findMinArrowShots(list)

	fmt.Println(r)

}

//路径，轨迹
type data [][]int

func (a data) Len() int {
	return len(a)
}
func (a data) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a data) Less(i, j int) bool {
	return a[i][0] < a[j][0]
}
func findMinArrowShots(points [][]int) int {
	sort.Sort(data(points))
	return findMinArrowShots2(0, points)

}
func findMinArrowShots2(acc int, points [][]int) int {
	if points == nil || len(points) == 0 {
		return acc
	} else {
		rest := trimFromLeft(points)
		return findMinArrowShots2(acc+1, rest)
	}
}

/*
一次尽可能从左边取掉最大的有交集部分。
*/
func trimFromLeft(points [][]int) [][]int {
	if points == nil || len(points) < 2 {
		return nil
	} else {
		share01 := calShare(points[0], points[1])
		if share01 == nil {
			//去掉一个
			return points[1:]
		} else {
			//至少去掉两个了。
			points[1] = share01
			return trimFromLeft(points[1:])
		}
	}
}

/*
计算两个slice的公共部分。
*/
func calShare(a, b []int) []int {
	//[a0,a1], [b0,b1]  [1,3],[2,4] ->2,3, [1,2],[2,4]
	if a[1] < b[0] {
		return nil
	} else {
		start := min(a[1], b[0])
		end := min(a[1], b[1])
		return []int{start, end}
	}
}
func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
