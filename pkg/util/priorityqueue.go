// a simple implementation of a priority queue based on golangs heap structure
package util

import (
	"container/heap"
)

type Item struct {
	value    string
	priority int
	index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

// add new item to the priority queue
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

// remove and return item with highest priority
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// updates an items value and priority
func (pq *PriorityQueue) update(item *Item, value string, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index) // re-establishes heap ordering
}

/* Example Usage: -------------------------------------------------------------

// create some items
items := map[string]int {"a":35, "b":70, "c":118}

// create pq
pq := make(PriorityQueue, len(items))

// initialize queue with items
i := 0
for value, priority := range items {
	pq[i] = &Item{value, priority, i,}
	i++
}
heap.Init(&pq)

// insert new item
item := &Item{value:"foo", priority:1,}
heap.Push(&pq, item)

// update the priority of an item
pq.update(item, item.value, 5)

// Take items out; they arrive in increasing priority order.
for pq.Len() > 0 {
	item := heap.Pop(&pq).(*Item)
	fmt.Printf("%.2d:%s ", item.priority, item.value)
}
*/
