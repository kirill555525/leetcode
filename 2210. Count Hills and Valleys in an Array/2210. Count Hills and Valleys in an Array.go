package main

func countHillValley(nums []int) int {
	count := 0
	left := 0
	right := 0

	for i := 1; i < len(nums)-1; i++ {
		right = i + 1
		if nums[i] == nums[left] {
			left = i
			continue
		}

		if nums[i] == nums[right] {
			continue
		}

		if (nums[left] > nums[i] && nums[right] > nums[i]) || (nums[left] < nums[i] && nums[right] < nums[i]) {
			count++
		}

		left = i
	}

	return count
}
