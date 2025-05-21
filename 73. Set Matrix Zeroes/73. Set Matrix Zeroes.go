package main

func setZeroes(matrix [][]int) {
	flag := false

	for i := 0; i < len(matrix); i++ {
		if matrix[i][0] == 0 {
			flag = true
		}
		for j := 1; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	for j := 1; j < len(matrix[0]); j++ {
		if matrix[0][j] == 0 {
			for i := range matrix {
				matrix[i][j] = 0
			}
		}
	}

	for i := range matrix {
		if matrix[i][0] == 0 {
			for j := range matrix[i] {
				matrix[i][j] = 0
			}
		}

		if flag {
			matrix[i][0] = 0
		}
	}

}
