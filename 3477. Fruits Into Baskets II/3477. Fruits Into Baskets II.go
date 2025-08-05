package main

func numOfUnplacedFruits(fruits []int, baskets []int) int {
	count := 0
	for i := 0; i < len(fruits); i++ {
		for j := 0; j < len(baskets); j++ {
			if baskets[j] >= fruits[i] {
				count++
				baskets[j] = 0
				break
			}
		}
	}

	return len(fruits) - count
}
