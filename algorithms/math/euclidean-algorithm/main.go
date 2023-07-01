package main

import "fmt"

func gcd(a int, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func lcm(a int, b int) int {
	return a / gcd(a, b) * b
}

func main() {
	fmt.Println(gcd(15, 24) == 3)
	fmt.Println(gcd(7, 14) == 7)
	fmt.Println(gcd(14, 7) == 7)
	fmt.Println(gcd(15, 28) == 1)

	fmt.Println(lcm(15, 24) == 120)
	fmt.Println(lcm(7, 14) == 14)
	fmt.Println(lcm(14, 7) == 14)
	fmt.Println(lcm(15, 28) == 15*28)
}
