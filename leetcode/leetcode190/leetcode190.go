package leetcode190

import (
	"fmt"
)

func reverseBits(num uint32) uint32 {
	var sum uint32 = 0
	for i := 0; i < 32; i++ {
		sum = num&1 | (sum << 1)
		num = num >> 1
	}
	return sum
}
