package main

import (
	"fmt"
)

func matrixMul(a, b [][]int64, mod int64) [][]int64 {
	n := len(a)
	m := len(a[0])
	k := len(b[0])
	c := make([][]int64, n)
	for i := 0; i < n; i++ {
		c[i] = make([]int64, k)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < k; j++ {
			for l := 0; l < m; l++ {
				c[i][j] += (a[i][l] * b[l][j]) % mod
			}
			c[i][j] %= mod
		}
	}

	return c
}

func binPow(a [][]int64, n, mod int64) [][]int64 {
	res := make([][]int64, len(a))
	for i := 0; i < len(a); i++ {
		res[i] = make([]int64, len(a))
		res[i][i] = 1
	}

	for n > 0 {
		if n&1 != 0 {
			res = matrixMul(res, a, mod)
		}
		n >>= 1
		a = matrixMul(a, a, mod)
	}

	return res
}

func main() {
	fibMatrix := [][]int64{
		{1, 1},
		{1, 0},
	}
	fibFirst := [][]int64{
		{1, 0}, // {f1, f0}
	}
	f1 := int64(1)
	f0 := int64(0)
	mod := int64(10000)

	for i := 1; i <= 1000; i++ {
		fibFirst = matrixMul(fibFirst, fibMatrix, mod)

		f2 := (f1 + f0) % mod
		f0 = f1
		f1 = f2

		if f1 != fibFirst[0][0] {
			fmt.Println("ERROR:", f1, fibFirst[0][0])
		}

		matrixPow := binPow(fibMatrix, int64(i), mod)
		fibs := matrixMul([][]int64{{1, 0}}, matrixPow, mod)
		if fibs[0][0] == fibFirst[0][0] {
			fmt.Println("RIGHT:", fibs[0][0], fibFirst[0][0])
		}
	}

	matrixPow := binPow(fibMatrix, int64(1000000000000), mod)
	fibs := matrixMul([][]int64{{1, 0}}, matrixPow, mod)
	fmt.Println(fibs[0][0])
}
