package main

func maxProfit(prices []int) int {
	smallest := prices[0]
	current := prices[0]
	largest := 0
	for i := 0; i < len(prices)-1; i++ {
		current = prices[i+1]
		smallest = min(prices[i], smallest)
		largest = max(largest, current-smallest)
	}

	return largest
}
