package main

func productExceptSelf(nums []int) []int {
	total := 1
	n := len(nums)

	answer := make([]int, n)

	answer[0] = 1
	for i := 1; i < n; i++ {
		answer[i] = answer[i-1] * nums[i-1]
	}

	total = 1

	for j := n - 2; j >= 0; j-- {
		total *= nums[j+1]
		answer[j] *= total
	}

	return answer
}
