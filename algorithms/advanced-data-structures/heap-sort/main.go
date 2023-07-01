package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"sort"
)

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

func empty(heap []int) bool {
	return len(heap) == 1
}

func heapSort(arr *[]int) {
	heap := []int{0}
	for _, v := range *arr {
		add(&heap, v)
	}
	*arr = []int{}
	for !empty(heap) {
		*arr = append(*arr, top(heap))
		erase(&heap)
	}
}

func main() {
	var arr []int
	n := rand.Intn(100)
	for i := 0; i < n; i++ {
		arr = append(arr, rand.Intn(100))
	}
	heapSort(&arr)

	arrCopy := append([]int(nil), arr...)
	sort.Ints(arrCopy)

	fmt.Println(arr)
	fmt.Println(arrCopy)
	fmt.Println(reflect.DeepEqual(arrCopy, arr))
}
