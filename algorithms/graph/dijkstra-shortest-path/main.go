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

	dist := dijkstra(graph, 0)
	for _, v := range dist {
		fmt.Printf("%d ", v) // 0 10 11 16 17 5 16 8 10 9
	}
}
