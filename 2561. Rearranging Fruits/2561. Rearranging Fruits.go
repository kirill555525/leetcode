package main

import "sort"

func minCost(basket1 []int, basket2 []int) int64 {
	m := make(map[int]int)

	smallest := basket1[0]

	for i := 0; i < len(basket1); i++ {
		m[basket1[i]]++
		m[basket2[i]]--

		smallest = min(smallest, basket1[i], basket2[i])
	}

	temp1 := []int{}
	temp2 := []int{}

	for key, value := range m {
		if value%2 == 1 || value%2 == -1 {
			return -1
		}
		if value > 0 {
			for value != 0 {
				temp1 = append(temp1, key)
				value -= 2
			}
		} else if value < 0 {
			for value != 0 {
				temp2 = append(temp2, key)
				value += 2
			}
		}

	}

	if len(temp1) != len(temp2) {
		return -1
	}

	sort.Slice(temp1, func(i, j int) bool {
		return temp1[i] < temp1[j]
	})

	sort.Slice(temp2, func(i, j int) bool {
		return temp2[i] < temp2[j]
	})

	var cost int64

	for i, j := 0, len(temp2)-1; i < len(temp1); i, j = i+1, j-1 {
		cost += min(int64(temp1[i]), int64(temp2[j]), int64(2*smallest))
	}

	return cost

}
