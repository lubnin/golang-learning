package main

import (
	"fmt"
)

// O(n!)
func queens(row int, col, diag1, diag2 []int) int {
	n := len(col)
	if row == n {
		return 1
	}

	count := 0
	for column := 0; column < n; column++ {
		if col[column] == 0 && diag1[column+row] == 0 && diag2[row-column+(n-1)] == 0 {
			col[column] = 1
			diag1[column+row] = 1
			diag2[row-column+(n-1)] = 1

			count += queens(row+1, col, diag1, diag2)

			col[column] = 0
			diag1[column+row] = 0
			diag2[row-column+(n-1)] = 0
		}
	}

	return count
}

func main() {
	for n := 1; n <= 12; n++ {
		col := make([]int, n)
		diag1 := make([]int, 2*n-1)
		diag2 := make([]int, 2*n-1)
		fmt.Println(n, "=", queens(0, col, diag1, diag2))
	}
}
