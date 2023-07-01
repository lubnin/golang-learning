package main

import (
	"fmt"
)

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	L := []int{1}

	for i := 1; i <= n; i++ {
		temp := 0
		for j := 0; j < k; j++ {
			if len(L) < (1 + j) {
				break
			}
			temp += L[len(L)-1-j]
			temp %= 10000
		}
		L = append(L, temp)
	}

	for i := 0; i < n; i++ {
		fmt.Print(L[i], " ")
	}
}
