package main

func candy(ratings []int) int {

	n := len(ratings)
	candies := make([]int, n)

	candies[0] = 1

	index := 0

	for i := 1; i < n; i++ {

		if ratings[i] < ratings[i-1] {
			continue
		}

		for index != i {
			candies[index] = max(candies[index], i-index)
			index++
		}

		if ratings[i] > ratings[i-1] {
			candies[i] = candies[i-1] + 1
			index = i
		}

		if ratings[i] == ratings[i-1] {
			candies[i] = 1
			index = i
		}
	}

	for index != n {
		candies[index] = max(candies[index], n-index)
		index++
	}

	total := 0

	for _, elem := range candies {
		total += elem
	}

	return total

}
