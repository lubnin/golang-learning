package main

import (
	"fmt"
)

// O(m)
func listOfEdges() {
	var n, m int
	fmt.Scan(&n, &m)
	edges := make([][2]int, m)
	for i := 0; i < m; i++ {
		var a, b int
		fmt.Scan(&a, &b)
		edges[i] = [2]int{a, b}
	}

	fmt.Println("List of edges:")
	for i := 0; i < m; i++ {
		fmt.Println(edges[i][0], edges[i][1])
	}
	fmt.Println()
	fmt.Println()
}

// O(n^2)
func adjacencyMatrix() {
	var n, m int
	fmt.Scan(&n, &m)
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		var a, b int
		fmt.Scan(&a, &b)
		// 1 <= a, b <= n
		a--
		b--
		matrix[a][b] = 1
		matrix[b][a] = 1 // if graph is undirected
	}

	fmt.Println("Adjacency matrix:")
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Print(matrix[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
	fmt.Println()
}

// O(n^2)
func adjacencyList() {
	var n, m int
	fmt.Scan(&n, &m)
	matrix := make([][]int, n)
	for i := 0; i < m; i++ {
		var a, b int
		fmt.Scan(&a, &b)
		a--
		b--
		matrix[a] = append(matrix[a], b)
		matrix[b] = append(matrix[b], a) // if graph is undirected
	}

	fmt.Println("Adjacency list:")
	for i := 0; i < n; i++ {
		fmt.Print(i+1, ": ")
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Print(matrix[i][j]+1, " ")
		}
		fmt.Println()
	}
	fmt.Println()
	fmt.Println()
}

/*
first row is a pair of nodes and edges numbers
next values are a sequence of edges itself
input example:
5 7
1 2
1 3
1 4
1 5
2 3
3 4
4 5
*/

func main() {
	listOfEdges()
	adjacencyMatrix()
	adjacencyList()
}
