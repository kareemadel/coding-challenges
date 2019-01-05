func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	minPrice, maxProfit := prices[0], 0
	for _, v := range prices {
		minPrice = min(v, minPrice)
		maxProfit = max(maxProfit, v-minPrice)
	}
	return maxProfit
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}