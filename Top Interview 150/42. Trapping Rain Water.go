package main

func trap(height []int) int {

	mHeight := 0
	total := 0

	j := 0
	m := -1

	for i := range height {
		if height[i] > m {
			j = i
			m = height[i]
		}
	}

	for i, h := range height {
		if i == j {
			break
		}

		mHeight = max(mHeight, h)

		total += mHeight - h
	}

	mHeight = 0

	for i := len(height) - 1; i != j; i-- {
		h := height[i]

		mHeight = max(mHeight, h)

		total += mHeight - h
	}

	return total
}
