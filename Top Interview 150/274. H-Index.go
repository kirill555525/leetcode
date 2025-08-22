package main

func hIndex(citations []int) int {
	m := make(map[int]int)

	total := 0
	maxTotal := 0

	for _, citation := range citations {
		if citation > total {
			m[citation]++
			total++
			x := total - 1
			total -= m[x]
			delete(m, x)
			maxTotal = max(maxTotal, total)
		}
	}

	return maxTotal
}
