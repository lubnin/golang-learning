package main

import (
	"fmt"
)

func dfs(graph [][]int, used []int, order *[]int, currentVertex int) {
	used[currentVertex] = 1
	for i := 0; i < len(graph[currentVertex]); i++ {
		adjacentVertex := graph[currentVertex][i]
		if used[adjacentVertex] != 1 {
			dfs(graph, used, order, adjacentVertex)
		}
	}
	*order = append(*order, currentVertex)
}

func main() {
	graph := [][]int{
		{4},
		{0},
		{3, 1},
		{0, 1, 4},
		{},
	}

	used := make([]int, len(graph))
	var order []int

	for i := 0; i < len(graph); i++ {
		if used[i] == 0 {
			dfs(graph, used, &order, i)
		}
	}

	for i := 0; i < len(graph); i++ {
		fmt.Println(i+1, "->", order[i]+1)
	}
}
