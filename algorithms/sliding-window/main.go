package main

import "fmt"

func main() {
	var arr = []int{1, 7, 5, 3, 2, 3, 8, 9}
	var k = 4
	var sum = 0
	for i := 0; i < k; i++ {
		sum += arr[i]
	}
	fmt.Println(sum)
	// O(n)
	for i := 1; i <= len(arr)-k; i++ {
		sum = sum - arr[i-1] + arr[i+k-1]
		fmt.Println(sum)
	}

	fmt.Println()

	// O(n * k)
	for i := 0; i <= len(arr)-k; i++ {
		sum := 0
		for j := 0; j < k; j++ {
			sum += arr[i+j]
		}
		fmt.Println(sum)
	}
}
