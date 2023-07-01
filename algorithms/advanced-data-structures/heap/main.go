package main

import "fmt"

func add(heap *[]int, x int) { // O(log(n))
	*heap = append(*heap, x)
	ind := len(*heap) - 1
	for ind != 1 && (*heap)[ind] < (*heap)[ind/2] {
		(*heap)[ind], (*heap)[ind/2] = (*heap)[ind/2], (*heap)[ind]
		ind /= 2
	}
}

func erase(heap *[]int) { // O(log(n))
	(*heap)[1], (*heap)[len(*heap)-1] = (*heap)[len(*heap)-1], (*heap)[1]
	*heap = (*heap)[:len(*heap)-1]
	ind := 1
	for ind*2 < len(*heap) && (*heap)[ind] > (*heap)[ind*2] || ind*2+1 < len(*heap) && (*heap)[ind] > (*heap)[ind*2+1] {
		if ind*2+1 >= len(*heap) || (*heap)[ind*2] < (*heap)[ind*2+1] {
			(*heap)[ind], (*heap)[ind*2] = (*heap)[ind*2], (*heap)[ind]
			ind *= 2
		} else {
			(*heap)[ind], (*heap)[ind*2+1] = (*heap)[ind*2+1], (*heap)[ind]
			ind *= 2
			ind++
		}
	}
}

func top(heap []int) int { // O(1)
	return heap[1]
}

func main() {
	heap := []int{0}
	add(&heap, 2)
	add(&heap, 5)
	add(&heap, 3)
	add(&heap, 10)
	add(&heap, 6)
	add(&heap, 9)
	add(&heap, 5)
	add(&heap, 12)
	add(&heap, 40)
	fmt.Println(heap)

	fmt.Println(top(heap))
	erase(&heap)
	fmt.Println(top(heap))
	erase(&heap)
	fmt.Println(top(heap))
	erase(&heap)
	fmt.Println(top(heap))
	erase(&heap)
	fmt.Println(top(heap))
	erase(&heap)
	fmt.Println(top(heap))
	erase(&heap)
	fmt.Println(top(heap))
	erase(&heap)
	fmt.Println(top(heap))
	erase(&heap)
	fmt.Println(top(heap))
}
