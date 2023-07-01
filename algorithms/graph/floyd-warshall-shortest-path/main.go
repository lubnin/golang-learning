package main

import (
	"fmt"
)

const Inf = 1000000

func dijkstra(graph [][]pair, startVertex int) []int {
	dist := make([]int, len(graph))
	for i := range dist {
		dist[i] = Inf
	}
	dist[startVertex] = 0
	s := make(set, 0)
	s.insert(pair{0, startVertex})

	for !s.empty() {
		currentVertex := s.begin().second
		s.erase(s.begin())

		for i := 0; i < len(graph[currentVertex]); i++ {
			adjacentVertex := graph[currentVertex][i].first
			weight := graph[currentVertex][i].second

			if dist[currentVertex]+weight < dist[adjacentVertex] {
				s.erase(pair{dist[adjacentVertex], adjacentVertex})
				dist[adjacentVertex] = dist[currentVertex] + weight
				s.insert(pair{dist[adjacentVertex], adjacentVertex})
			}
		}
	}
	return dist
}

func floydWarshall(dist [][]int) [][]int {
	n := len(dist)
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				dist[i][j] = min(dist[i][j], dist[i][k]+dist[k][j])
			}
		}
	}
	return dist
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type pair struct {
	first  int
	second int
}

type set []pair

func (s set) empty() bool {
	return len(s) == 0
}

func (s *set) insert(p pair) {
	*s = append(*s, p)
}

func (s set) begin() pair {
	return s[0]
}

func (s *set) erase(p pair) {
	for i, v := range *s {
		if v == p {
			*s = append((*s)[:i], (*s)[i+1:]...)
			break
		}
	}
}

func main() {
	graph := [][]pair{
		{{1, 10}, {5, 5}},
		{{0, 10}, {2, 1}},
		{{1, 1}, {3, 5}, {5, 7}, {6, 10}},
		{{2, 5}, {4, 1}},
		{{3, 1}, {6, 2}},
		{{0, 5}, {2, 7}, {6, 100}, {7, 3}},
		{{2, 10}, {4, 2}, {5, 100}, {7, 8}, {8, 100}},
		{{5, 3}, {6, 8}, {9, 1}},
		{{6, 100}, {9, 1}},
		{{7, 1}, {8, 1}},
	}

	graph2 := make([][]int, len(graph))
	for i := range graph2 {
		graph2[i] = make([]int, len(graph))
		for j := range graph2[i] {
			graph2[i][j] = Inf
		}
	}

	for i := 0; i < len(graph); i++ {
		for j := 0; j < len(graph[i]); j++ {
			adjacentVertex := graph[i][j].first
			weight := graph[i][j].second
			graph2[i][adjacentVertex] = weight
			graph2[adjacentVertex][i] = weight
		}
	}

	for i := 0; i < len(graph); i++ {
		dist := dijkstra(graph, i)
		for j := 0; j < len(dist); j++ {
			fmt.Printf("%d ", dist[j])
		}
		fmt.Println()
	}
	fmt.Println()
	fmt.Println()

	dist := floydWarshall(graph2)
	for i := 0; i < len(dist); i++ {
		for j := 0; j < len(dist[i]); j++ {
			if i == j {
				fmt.Printf("0 ")
			} else {
				fmt.Printf("%d ", dist[i][j])
			}
		}
		fmt.Println()
	}
}
