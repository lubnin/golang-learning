package main

import (
	"fmt"
)

func dfs(graph [][]int, edges map[[2]int]int, cycle *[]int, currentVertex int) {
	for i := 0; i < len(graph[currentVertex]); i++ {
		adjacentVertex := graph[currentVertex][i]
		a := currentVertex
		b := adjacentVertex
		if a > b {
			a, b = b, a
		}
		if edges[[2]int{a, b}] != 0 {
			edges[[2]int{a, b}]--
			dfs(graph, edges, cycle, adjacentVertex)
		}
	}
	*cycle = append(*cycle, currentVertex)
}

func main() {
	graph := [][]int{
		{1, 3, 4},
		{0, 2, 3, 4},
		{1, 3},
		{0, 1, 2, 4},
		{0, 1, 3},
	}

	var oddVertices []int
	for i := 0; i < len(graph); i++ {
		if len(graph[i])%2 == 1 {
			oddVertices = append(oddVertices, i)
		}
	}

	if len(oddVertices) > 2 {
		fmt.Println("No Eulerian cycle and path")
		return
	} else if len(oddVertices) == 2 {
		graph[oddVertices[0]] = append(graph[oddVertices[0]], oddVertices[1])
		graph[oddVertices[1]] = append(graph[oddVertices[1]], oddVertices[0])
	}

	var cycle []int
	edges := make(map[[2]int]int)
	for i := 0; i < len(graph); i++ {
		for j := 0; j < len(graph[i]); j++ {
			if i < graph[i][j] {
				edges[[2]int{i, graph[i][j]}]++
			}
		}
	}

	dfs(graph, edges, &cycle, 0)

	if len(oddVertices) == 2 {
		for i := 1; i < len(cycle); i++ {
			if (cycle[i-1] == oddVertices[0] && cycle[i] == oddVertices[1]) ||
				(cycle[i-1] == oddVertices[1] && cycle[i] == oddVertices[0]) {
				var newCycle []int
				for j := i; j < i+len(cycle); j++ {
					index := j % len(cycle)
					if index == 0 {
						continue
					}
					newCycle = append(newCycle, cycle[index])
				}
				cycle = newCycle
			}
		}
	}

	for i := 0; i < len(cycle); i++ {
		fmt.Print(cycle[i]+1, " ")
	}
}
