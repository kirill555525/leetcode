package main

import (
	"fmt"
	"log"
)

func multiplyMatrices(a, b [][]int) [][]int {
	n := len(a)
	result := make([][]int, n)
	for i := range result {
		result[i] = make([]int, n)
		for j := 0; j < n; j++ {
			sum := 0
			for k := 0; k < n; k++ {
				sum = (sum + a[i][k]*b[k][j]) % (1_000_000_000 + 7)
			}
			result[i][j] = sum
		}
	}
	return result
}

func MultiplyMatrixVector(A [][]int, x []int) []int {
	m := len(A)
	if m == 0 {
		return nil
	}
	n := len(A[0])
	if len(x) != n {
		log.Fatalf("Размер вектора (%d) не совпадает с числом столбцов матрицы (%d)", len(x), n)
	}

	result := make([]int, m)
	for i := 0; i < m; i++ {
		sum := 0
		for j := 0; j < n; j++ {
			sum = (sum + A[j][i]*x[j]) % (1_000_000_000 + 7)
		}
		result[i] = sum
	}
	return result
}

func lengthAfterTransformations(s string, t int, nums []int) int {

	matrix := make([][]int, 26)
	for i := range matrix {
		matrix[i] = make([]int, 26)
	}

	for i, v := range nums {
		for j := i + 1; v != 0; v, j = v-1, j+1 {
			if j == 26 {
				j = 0
			}
			matrix[i][j] = 1
		}
	}

	v := make([]int, 26)
	for i := 0; i < len(s); i++ {
		v[s[i]-'a']++
	}

	for t != 0 {
		if t&1 == 1 {
			v = MultiplyMatrixVector(matrix, v)
		}
		matrix = multiplyMatrices(matrix, matrix)
		t >>= 1
	}

	total := 0

	for _, value := range v {
		total = (total + value) % (1_000_000_000 + 7)
	}
	return total

}

func main() {
	s := "k"
	t := 2
	fmt.Println(lengthAfterTransformations(s, t, []int{2, 2, 1, 3, 1, 1, 2, 3, 3, 2, 1, 2, 2, 1, 1, 3, 1, 2, 2, 1, 3, 3, 3, 2, 2, 1}))

}
