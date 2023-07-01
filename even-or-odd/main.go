package main

import "fmt"

func main() {

	var a = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var s string

	for _, n := range a {
		s = ""
		if n%2 == 0 {
			s = "even"
		} else {
			s = "odd"
		}
		fmt.Printf("element %d is %s \n", n, s)
	}
}
