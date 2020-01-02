package leetcode121

func maxProfit(prices []int) int {
	//买卖股票，低买高卖。只能买一次，卖一次。先买再卖。
	//如果是小的在前的升序，就可以买卖一次。如果是降序，不买卖。0
	//取第一个数，再找余下的最大的数，计算最大的减去这个数。
	l := len(prices)

	if l <= 1 {
		return 0
	} else {
		maxPrice := prices[l-1]
		maxProfit := 0
		for i := l - 2; i >= 0; i-- {
			maxProfit = max(maxProfit, maxPrice-prices[i])
			maxPrice = max(maxPrice, prices[i])
		}
		return maxProfit
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
