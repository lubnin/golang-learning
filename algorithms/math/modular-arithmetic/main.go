package main

import "fmt"

func binPow(a int64, n int64, mod int64) int64 {
	result := int64(1)
	for n != 0 {
		if n&1 != 0 {
			result *= a
			result %= mod
		}
		n >>= 1
		a *= a
		a %= mod
	}
	return result
}

func main() {
	mod := int64(1000000007)
	fmt.Println(2 * binPow(5, mod-2, mod) % mod)
}
