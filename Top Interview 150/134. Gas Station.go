package main

func canCompleteCircuit(gas []int, cost []int) int {
	total := 0
	index := 0
	for i := range gas {
		total += gas[i] - cost[i]
		if total < 0 {
			total = 0
			index = i + 1
		}
	}

	if index != len(gas) {
		for i := 0; i < index; i++ {
			total += gas[i] - cost[i]
			if total < 0 {
				return -1
			}
		}
	} else {
		return -1
	}
	return index
}
