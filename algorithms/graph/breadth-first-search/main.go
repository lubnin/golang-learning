package main

import (
	"fmt"
)

func bfsList(graph [][]int) {
	startVertex := 0
	used := make([]int, len(graph))
	used[startVertex] = 1
	queue := []int{startVertex}

	for len(queue) > 0 {
		currentVertex := queue[0]
		queue = queue[1:]
		fmt.Print(currentVertex+1, " ")

		for i := 0; i < len(graph[currentVertex]); i++ {
			adjacentVertex := graph[currentVertex][i]
			if used[adjacentVertex] == 0 {
				queue = append(queue, adjacentVertex)
				used[adjacentVertex] = 1
			}
		}
	}
}

func bfsMatrix(graph [][]int) {
	startVertex := 0
	used := make([]int, len(graph))
	used[startVertex] = 1
	queue := []int{startVertex}

	for len(queue) > 0 {
		currentVertex := queue[0]
		queue = queue[1:]
		fmt.Print(currentVertex+1, " ")

		for i := 0; i < len(graph[currentVertex]); i++ {
			if graph[currentVertex][i] == 1 {
				adjacentVertex := i
				if used[adjacentVertex] == 0 {
					queue = append(queue, adjacentVertex)
					used[adjacentVertex] = 1
				}
			}
		}
	}
}

func main() {
	adjacencyList := [][]int{
		{1, 3},
		{0, 5, 6},
		{6},
		{0, 4},
		{3, 7},
		{1, 6},
		{1, 2, 5},
		{4},
	}

	adjacencyMatrix := [][]int{
		{0, 1, 0, 1, 0, 0, 0, 0},
		{1, 0, 0, 0, 0, 1, 1, 0},
		{0, 0, 0, 0, 0, 0, 1, 0},
		{1, 0, 0, 0, 1, 0, 0, 0},
		{0, 0, 0, 1, 0, 0, 0, 1},
		{0, 1, 0, 0, 0, 0, 1, 0},
		{0, 1, 1, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 1, 0, 0, 0},
	}

	bfsList(adjacencyList)
	fmt.Println()
	bfsMatrix(adjacencyMatrix)
}
