package main

func removeElement(nums []int, val int) int {
	start := 0
	end := len(nums) - 1
	counter := 0

	for start <= end {
		if nums[start] != val {
			start++
			counter++
		} else if nums[end] == val {
			end--
		} else {
			nums[start], nums[end] = nums[end], nums[start]
			start++
			counter++
		}

	}

	return counter
}
