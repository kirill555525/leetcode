package main

func maxSum(nums []int) int {
	m := make(map[int]struct{})
	total := 0
	flag := -100
	for _, elem := range nums {
		if elem <= 0 {
			if elem > flag {
				flag = elem
			}
			continue
		}
		if _, ok := m[elem]; !ok {
			m[elem] = struct{}{}
			total += elem
		}
	}

	if total == 0 {
		total = flag
	}

	return total
}
