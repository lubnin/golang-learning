package main

import (
	"fmt"
	"reflect"
)

func matrixMul(A, B [][]int) [][]int {
	n := len(A)
	m := len(A[0])
	k := len(B[0])

	C := make([][]int, n)
	for i := 0; i < n; i++ {
		C[i] = make([]int, k)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < k; j++ {
			for k := 0; k < m; k++ {
				C[i][j] += A[i][k] * B[k][j]
			}
		}
	}
	return C
}

func main() {
	A := [][]int{
		{1, 2, 3},
		{3, 1, 2},
	}

	B := [][]int{
		{1, 2},
		{3, 2},
		{1, 2},
	}

	C := [][]int{
		{10, 12},
		{8, 12},
	}

	fmt.Println(reflect.DeepEqual(matrixMul(A, B), C))
}
