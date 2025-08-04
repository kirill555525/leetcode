package main

func totalFruit(fruits []int) int {
	x := fruits[0]
	y := -1

	length := 1
	maxLength := 0
	counter := 0

	for i := 1; i < len(fruits); i++ {
		if x == fruits[i] {
			if fruits[i-1] == x {
				counter++
			} else {
				counter = 1
			}
			length++
			continue
		}

		if y == fruits[i] {
			if fruits[i-1] == y {
				counter++
			} else {
				counter = 1
			}
			length++
			continue
		}

		if y == -1 {
			y = fruits[i]
			counter = 1
			length++
			continue
		}

		x = fruits[i]
		y = fruits[i-1]
		maxLength = max(maxLength, length)
		length = 1 + counter
		counter = 1

	}

	maxLength = max(maxLength, length)

	return maxLength
}
