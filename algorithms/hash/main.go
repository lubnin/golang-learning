package main

import (
	"fmt"
)

func main() {
	s := "asdadqwezxcasdfqcsdfsf"
	hashes := make([]int64, len(s)+1)
	hashes[0] = 0
	primes := make([]int64, len(s)+1)
	primes[0] = 1

	p := int64(301)
	m := int64(1e9 + 7)

	for i := 0; i < len(s); i++ {
		hashes[i+1] = (hashes[i]*p + int64(s[i]-'a'+1)) % m
		primes[i+1] = (primes[i] * p) % m
	}

	i, j := 3, 7 // adqwe

	substrHash := (hashes[j+1] - (hashes[i]*primes[j-i+1])%m) % m
	if substrHash < 0 {
		substrHash += m
	}

	fmt.Println(substrHash)
}
