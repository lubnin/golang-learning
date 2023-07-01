package main

import "fmt"

func dfs(graph [][]int, used []int, currentVertex int) {
	fmt.Print(currentVertex+1, " ")
	used[currentVertex] = 1
	for i := 0; i < len(graph[currentVertex]); i++ {
		adjacentVertex := graph[currentVertex][i]
		if used[adjacentVertex] != 1 {
			dfs(graph, used, adjacentVertex)
		}
	}
}

func main() {
	graph := [][]int{
		{4, 5},
		{3},
		{3},
		{1, 2, 9},
		{0},
		{0},
		{},
		{8},
		{7},
		{3},
	}

	used := make([]int, len(graph))

	for i := 0; i < len(graph); i++ {
		if used[i] != 1 {
			dfs(graph, used, i)
			fmt.Println()
		}
	}
}
