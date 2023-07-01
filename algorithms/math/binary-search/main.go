package main

import "fmt"

func binarySearch(a []int, x int) bool {
	l := 0
	r := len(a) - 1
	for l <= r {
		m := (l + r) / 2
		if a[m] == x {
			return true
		}
		if a[m] < x {
			l = m + 1
		} else if a[m] > x {
			r = m - 1
		}
	}
	return false
}

func main() {
	arr := []int{1, 2, 4, 5, 7, 9, 12, 15, 19, 23}
	fmt.Println(binarySearch(arr, 7))  // true
	fmt.Println(binarySearch(arr, 12)) // true
	fmt.Println(binarySearch(arr, 23)) // true
	fmt.Println(binarySearch(arr, 8))  // false
	fmt.Println(binarySearch(arr, 3))  // false
	fmt.Println(binarySearch(arr, -1)) // false
}
