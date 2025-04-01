package solution

func maxProfit(prices []int) int {
	minPrice, maxProfit := int(1e4), 0
	for _, p := range prices {
		if p < minPrice {
			// Only can buy
			minPrice = p
		} else if p-minPrice > maxProfit {
			// Only can sell
			maxProfit = p - minPrice
		}
	}
	return maxProfit
}
