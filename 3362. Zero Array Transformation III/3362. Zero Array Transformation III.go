package main

import (
	"container/heap"
	"fmt"
	"slices"
)

type Heap []int

func (h *Heap) Len() int {
	return len(*h)
}
func (h *Heap) Less(i, j int) bool {
	return (*h)[i] > (*h)[j]
}
func (h *Heap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}
func (h *Heap) Push(x any) {
	*h = append(*h, x.(int))
}
func (h *Heap) Pop() any {
	n := len(*h)
	x := (*h)[n-1]
	*h = (*h)[:n-1]
	return x
}

func maxRemoval(nums []int, queries [][]int) int {
	pq := &Heap{}
	heap.Init(pq)
	slices.SortFunc(queries, func(a, b []int) int {
		return a[0] - b[0]
	})

	deltaArray := make([]int, len(nums)+1)
	operations := 0

	for i := range nums {
		maxIndex := -1
		for index, value := range queries {
			if i == value[0] {
				heap.Push(pq, value[1])
				maxIndex = index
			} else {
				break
			}
		}
		queries = queries[maxIndex+1:]

		operations += deltaArray[i]

		for operations < nums[i] && len(*pq) > 0 && (*pq)[0] >= i {
			operations++
			deltaArray[heap.Pop(pq).(int)+1] -= 1
		}
		if operations < nums[i] {
			return -1
		}

	}
	return len(*pq)

}

func main() {
	nums := []int{1, 1, 1, 1}
	queries := [][]int{
		{1, 3}, {0, 2}, {1, 3}, {1, 2},
	}
	fmt.Println(maxRemoval(nums, queries))
}
