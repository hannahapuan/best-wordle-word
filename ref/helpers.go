package ref

import (
	"container/heap"
	"fmt"
)

// Print prints all items in the priority queue
func (pq *PriorityQueue) Print() {
	for pq.Len() > 0 {
		item := heap.Pop(pq).(*Item)
		fmt.Printf("%s\t%.2d\n", item.Value, item.Priority)
	}
	fmt.Println()
}
