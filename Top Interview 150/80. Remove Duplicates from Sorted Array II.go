package main

func removeDuplicates2(nums []int) int {
	index, occurance := 1, 1

	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			occurance++
		} else {
			occurance = 1
		}

		if occurance <= 2 {
			nums[index] = nums[i]
			index++
		}
	}

	return index
}
