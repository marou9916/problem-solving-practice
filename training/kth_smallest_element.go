package training

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h IntHeap) Len() int { return len(h) }

func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }

func (h IntHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x

}

// func KthSmallestElement(arr []int, k int) int {
// 	priorityQueue := &IntHeap{}

// 	heap.Init(priorityQueue)

// 	for _, value := range arr {
// 		heap.Push(priorityQueue, value)
// 	}

// 	var kth int

// 	for i := 0; i < k; i++ {
// 		kth = heap.Pop(priorityQueue).(int)

// 	}

// 	return kth
// }

func KthSmallestElement(arr []int, k int) int {
	maxHeap := &IntHeap{}

	heap.Init(maxHeap)



	for _, value := range arr {
		fmt.Printf("max heap state:%d\n", *maxHeap)
		fmt.Printf("current Value: %d\n\n", value)


		
		if len(*maxHeap) < k {
			heap.Push(maxHeap, value)
			} else if value < (*maxHeap)[0] {	
				heap.Pop(maxHeap)
				heap.Push(maxHeap, value)
		}
	}

	return (*maxHeap)[0]
}