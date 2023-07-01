package main

import "fmt"

func binPow1(a int, n int) int {
	if n == 0 {
		return 1
	}
	if n%2 == 1 {
		return binPow1(a, n-1) * a
	}
	b := binPow1(a, n/2)
	return b * b
}

func binPow2(a int, n int) int {
	result := 1
	for n != 0 {
		if n&1 != 0 {
			result *= a
		}
		n >>= 1
		a *= a
	}
	return result
}

func main() {
	fmt.Println(binPow1(2, 10) == 1024)
	fmt.Println(binPow1(3, 5) == 243)
	fmt.Println(binPow1(3, 15) == 14348907)
	fmt.Println(binPow1(2, 7) == 64) // false
	fmt.Println(binPow1(2, 7) == 128)

	fmt.Println()

	fmt.Println(binPow2(2, 10) == 1024)
	fmt.Println(binPow2(3, 5) == 243)
	fmt.Println(binPow2(3, 15) == 14348907)
	fmt.Println(binPow2(2, 7) == 64) // false
	fmt.Println(binPow2(2, 7) == 128)
}
