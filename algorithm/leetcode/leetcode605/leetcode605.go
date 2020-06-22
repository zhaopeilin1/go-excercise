package leetcode605

import (
	"fmt"
)

func canPlaceFlowers(flowerbed []int, n int) bool {
	// flowerbed = [1,0,0,0,1], n = 1
	//紧密种植，0,2,4,6这样。空的（长度+1）/2。如果有一棵树，它的左边，右边没有树时，(长度-1)再进行紧密种植。
	//两棵树中间。如果是奇数个0，长度-2再进行紧密种值。偶数个0，减2，再进行紧密种植
	l := 0
	sum := 0
	for _, v := range flowerbed {
		if v == 1 {
			sum = sum + genEmptyFulltree(l-1)
			//mt.Println(sum)
			l = -1
		} else {
			l++
		}
	}
	sum = sum + genEmptyFulltree(l)
	return sum >= n
}
func genEmptyFulltree(n int) int {
	return (n + 1) / 2
}
