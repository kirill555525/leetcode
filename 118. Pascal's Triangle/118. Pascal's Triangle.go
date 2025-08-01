package main

func generate(numRows int) [][]int {
	res := make([][]int, 1)
	res[0] = append(res[0], 1)
	if numRows == 1 {
		return res
	}

	res = append(res, []int{1, 1})

	for i := 2; i < numRows; i++ {
		temp := []int{1}
		for j := 1; j < len(res[i-1]); j++ {
			temp = append(temp, res[i-1][j]+res[i-1][j-1])
		}
		temp = append(temp, 1)
		res = append(res, temp)
	}

	return res
}
