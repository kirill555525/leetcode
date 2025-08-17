package main

import "slices"

func rotate(nums []int, k int) {

	n := len(nums)

	k = k % n

	slices.Reverse(nums)

	slices.Reverse(nums[:k])
	slices.Reverse(nums[k:])

}
