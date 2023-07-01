package main

import "fmt"

func bfs(graph [][]int, startVertex int) {
	used := make([]int, len(graph))
	for i := range used {
		used[i] = -1
	}
	used[startVertex] = 0
	queue := []int{startVertex}

	for len(queue) > 0 {
		currentVertex := queue[0]
		queue = queue[1:]
		fmt.Print(currentVertex+1, " ")

		for i := 0; i < len(graph[currentVertex]); i++ {
			adjacentVertex := graph[currentVertex][i]
			if used[adjacentVertex] == -1 {
				queue = append(queue, adjacentVertex)
				used[adjacentVertex] = used[currentVertex] + 1
			}
		}
	}

	fmt.Println()
	for i := 0; i < len(used); i++ {
		fmt.Println(i+1, used[i])
	}
}

func main() {
	graph := [][]int{
		{1, 2},
		{0, 3, 6},
		{0, 3},
		{1, 2, 4, 5, 6},
		{3, 7},
		{3, 7},
		{1, 3, 7},
		{4, 5, 6, 8},
		{7},
	}

	bfs(graph, 6)
}
