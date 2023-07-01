package main

import "fmt"

func gcd(a int, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func euler(n int) int { // O(n * log(n))
	count := 0
	for i := 1; i <= n; i++ {
		if gcd(i, n) == 1 {
			count++
		}
	}
	return count
}

func eulerFast(n int) int { // O(sqrt(n))
	result := n
	prime := 2
	for n >= prime*prime {
		if n%prime == 0 {
			result = result / prime * (prime - 1) // (1-1/p) = (p-1)/p
			for n%prime == 0 {
				n /= prime
			}
		}
		prime++
	}
	if n != 1 {
		result = result / n * (n - 1)
	}
	return result
}

func main() {
	fmt.Println(euler(9) == 6)
	fmt.Println(eulerFast(9) == 6)
	fmt.Println(euler(36) == 12)
	fmt.Println(eulerFast(36) == 12)
	fmt.Println(euler(84) == 24)
	fmt.Println(eulerFast(84) == 24)
}
