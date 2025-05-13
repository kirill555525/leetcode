package main

func twoSum(nums []int, target int) []int {
	res := make([]int, 2)
	m := make(map[int]int)
	for i, v := range nums {
		if j, ok := m[target-v]; ok {
			res[0] = j
			res[1] = i
			break
		} else {
			m[v] = i
		}
	}

	return res
}
