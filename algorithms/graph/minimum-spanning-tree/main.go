package main

import (
	"fmt"
	"sort"
)

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

type edge struct {
	a, b, c int
}

type byCost []edge

func (e byCost) Len() int           { return len(e) }
func (e byCost) Less(i, j int) bool { return e[i].c < e[j].c }
func (e byCost) Swap(i, j int)      { e[i], e[j] = e[j], e[i] }

// O(M*logN)
func main() {
	parent = make([]int, 7)
	sizes = make([]int, 7)
	for i := 0; i < len(parent); i++ {
		parent[i] = i
		sizes[i] = 1
	}

	edges := []edge{
		{1, 2, 5},
		{1, 3, 9},
		{1, 5, 1},
		{1, 6, 3},
		{2, 4, 8},
		{2, 6, 3},
		{3, 5, 4},
		{4, 6, 7},
		{5, 6, 2},
	}

	sort.Sort(byCost(edges))

	allCost := 0

	for i := 0; i < len(edges); i++ {
		a := edges[i].a
		b := edges[i].b
		c := edges[i].c
		if findSet(a) != findSet(b) {
			fmt.Println(a, b)
			allCost += c
			unionSet(a, b)
		}
	}
	fmt.Println(allCost)
}
