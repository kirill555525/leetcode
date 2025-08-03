package main

import "slices"

func maxTotalFruits(fruits [][]int, startPos int, k int) int {
	total := 0
	var arrLeft [][]int
	var arrRight [][]int
	amountRight := 0
	for _, arr := range fruits {
		if arr[0] == startPos {
			total += arr[1]
			continue
		}
		if arr[0] > startPos {
			amountRight += arr[1]
			arrRight = append(arrRight, []int{arr[0] - startPos, amountRight})
		} else {
			arrLeft = append(arrLeft, []int{startPos - arr[0], arr[1]})
		}

	}

	slices.Reverse(arrLeft)

	for i := 1; i < len(arrLeft); i++ {
		arrLeft[i][1] += arrLeft[i-1][1]
	}

	amount1 := 0
	amount2 := 0
	maxAmount := 0

	j := len(arrRight) - 1

	for _, left := range arrLeft {
		temp := k - left[0]
		if temp < 0 {
			break
		}

		amount1 = left[1]
		temp -= left[0]
		for j != -1 {
			if temp-arrRight[j][0] < 0 {
				j--
			} else {
				amount2 = arrRight[j][1]
				break
			}
		}

		maxAmount = max(maxAmount, amount1+amount2)
		amount2 = 0
	}

	amount1 = 0
	amount2 = 0

	j = len(arrLeft) - 1
	for _, right := range arrRight {
		temp := k - right[0]
		if temp < 0 {
			break
		}

		amount1 = right[1]
		temp -= right[0]

		for j != -1 {
			if temp-arrLeft[j][0] < 0 {
				j--
			} else {
				amount2 = arrLeft[j][1]
				break
			}
		}

		maxAmount = max(maxAmount, amount1+amount2)
		amount2 = 0
	}

	return maxAmount + total

}
