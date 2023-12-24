package ref

import (
	"container/heap"
	"fmt"
)

// Print prints all items in the priority queue
func (pq *PriorityQueue) Print() {
	for pq.Len() > 0 {
		item := heap.Pop(pq).(*Item)
		fmt.Printf("%.2d: %s\n", item.Priority, item.Value)
	}
	fmt.Println()
}
