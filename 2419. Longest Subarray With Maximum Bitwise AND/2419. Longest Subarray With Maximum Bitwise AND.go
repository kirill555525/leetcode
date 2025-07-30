package main

func longestSubarray(nums []int) int {
	m := 0
	for _, num := range nums {
		if num > m {
			m = num
		}
	}

	maxCount := 0
	currCount := 0

	for _, num := range nums {
		if num == m {
			currCount++
		} else {
			if maxCount < currCount {
				maxCount = currCount
			}
			currCount = 0
		}

	}

	if maxCount < currCount {
		maxCount = currCount
	}

	return maxCount
}