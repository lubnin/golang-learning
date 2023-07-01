package main

import "fmt"

func sieve(n int) []int {
	mark := make([]bool, n)
	for i := range mark {
		mark[i] = true
	}
	var primes []int
	primes = append(primes, 2)
	for i := 3; i*i <= n; i += 2 {
		if mark[i] {
			primes = append(primes, i)
			for j := i * i; j < n; j += i {
				mark[j] = false
			}
		}
	}
	for i := 3; i < n; i += 2 {
		if mark[i] {
			primes = append(primes, i)
		}
	}
	return primes
}

func main() {
	primes := sieve(100)
	for _, v := range primes {
		fmt.Printf("%d ", v)
	}
}
