func maxProfit(prices []int) int {
	maxProfit := 0
	minBuy := prices[0]

	for i := 1; i < len(prices); i++ {
		if prices[i] - minBuy > maxProfit {
			maxProfit = prices[i] - minBuy
		}
		if minBuy > prices[i] {
			minBuy = prices[i]
		}
	}

	return maxProfit
}
