package main

import "fmt"

var parent []int
var sizes []int

func findSet(v int) int {
	if v == parent[v] {
		return v
	}
	parent[v] = findSet(parent[v])
	return parent[v]
}

func unionSet(a, b int) {
	a = findSet(a)
	b = findSet(b)
	if a != b {
		if sizes[a] < sizes[b] {
			a, b = b, a
		}
		parent[b] = a
		sizes[a] += sizes[b]
	}
}

func main() {
	parent = make([]int, 10)
	sizes = make([]int, 10)
	for i := 0; i < len(parent); i++ {
		parent[i] = i
		sizes[i] = 1
	}

	fmt.Println(findSet(0))
	fmt.Println(findSet(1))
	unionSet(0, 1)
	fmt.Println(findSet(0))
	fmt.Println(findSet(1))

	unionSet(1, 2)
	unionSet(2, 3)
	unionSet(3, 4)
	unionSet(4, 5)

	fmt.Println(findSet(0))
	fmt.Println(findSet(1))
	fmt.Println(findSet(2))
	fmt.Println(findSet(3))
	fmt.Println(findSet(4))
	fmt.Println(findSet(5))
	fmt.Println(findSet(6))
}
