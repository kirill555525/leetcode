package main

func jump(nums []int) int {
	n := len(nums)
	jumps := 0
	farthest := 0
	end := 0
	for i := 0; i < n-1; i++ {
		farthest = max(farthest, i+nums[i])
		if farthest >= n-1 {
			jumps++
			return jumps
		}
		if i == end {
			jumps++
			end = farthest
		}

	}
	return jumps
}
