package leetcode122

import (
	"fmt"
)

func maxProfit(prices []int) int {
	//
	if len(prices) < 2 {
		return 0
	}
	buyIn := prices[0]
	sellOut := prices[0]
	profit := 0
	for _, price := range prices[1:] {
		//fmt.Println("in:", buyIn, "out:", sellOut)
		if price > sellOut {
			sellOut = price
		} else if price < sellOut {
			profit = profit + (sellOut - buyIn)
			buyIn = price
			sellOut = price
		}
	}
	profit = profit + (sellOut - buyIn)
	return profit
}
