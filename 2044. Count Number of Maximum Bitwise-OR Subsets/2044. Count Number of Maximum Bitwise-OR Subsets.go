package main

func combinations(temp []int, nums []int) [][]int {
	res := make([][]int, 0)
	tempCopy := make([]int, len(temp))
	copy(tempCopy, temp)
	res = append(res, tempCopy)
	for i, num := range nums {
		newTemp := append(temp, num)
		res = append(res, combinations(newTemp, nums[i+1:])...)
	}

	return res
}

func countMaxOrSubsets(nums []int) int {
	total := 0
	for _, num := range nums {
		total |= num
	}

	data := combinations([]int{}, nums)
	count := 0
	for i, v := range data {
		check := 0
		for j := range v {
			check |= data[i][j]
		}
		if check == total {
			count++
		}
	}

	return count
}
