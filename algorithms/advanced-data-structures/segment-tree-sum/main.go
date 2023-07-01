package main

import (
	"fmt"
)

var tree []int

const n = 8

func buildTree(a []int, v, l, r int) {
	if l == r {
		tree[v] = a[l]
		return
	}
	m := (l + r) / 2
	buildTree(a, v*2, l, m)
	buildTree(a, v*2+1, m+1, r)
	tree[v] = tree[2*v] + tree[2*v+1]
}

func sum(i, j, v, l, r int) int {
	if i > j {
		return 0
	}
	if i == l && j == r {
		return tree[v]
	}
	m := (l + r) / 2
	return sum(i, min(m, j), v*2, l, m) + sum(max(m+1, i), j, v*2+1, m+1, r)
}

func update(i, x, v, l, r int) {
	if l == r {
		tree[v] = x
		return
	}
	m := (l + r) / 2
	if i <= m {
		update(i, x, v*2, l, m)
	} else {
		update(i, x, v*2+1, m+1, r)
	}
	tree[v] = tree[2*v] + tree[2*v+1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8}
	tree = make([]int, 4*n)
	buildTree(a, 1, 0, n-1)
	fmt.Println(sum(1, 6, 1, 0, n-1))
	update(3, 10, 1, 0, n-1)
	fmt.Println(sum(1, 6, 1, 0, n-1))
}
