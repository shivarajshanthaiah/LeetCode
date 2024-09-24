func maxProfit(prices []int) int {
	minPrice := prices[0]
	maxProfit := 0
	for _, num := range prices {
		if num < minPrice {
			minPrice = num
		} else {
			price := num - minPrice
			if price > maxProfit {
				maxProfit = price
			}
		}
	}
	return maxProfit
}