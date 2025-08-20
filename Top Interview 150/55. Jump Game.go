package main

func canJump(nums []int) bool {
	jumps := 0
	for i, num := range nums {

		if i == len(nums)-1 {
			return true
		}

		jumps = max(num, jumps)

		if jumps == 0 {
			return false
		}
		jumps--
	}

	return true

}
