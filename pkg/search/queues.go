package search

import (
	"container/heap"
)

type NodeQueue interface {
	Len() int
	Put(interface{})
	Take() interface{}
}

// A PriorityQueue implements heap.Interface and holds ValueNodes
type PriorityQueue []ValueNode

func NewPriorityQueue() *PriorityQueue {
	q := &PriorityQueue{}
	heap.Init(q)
	return q
}

func (q PriorityQueue) Len() int           { return len(q) }
func (q PriorityQueue) Less(i, j int) bool { return q[i].H() < q[j].H() }
func (q PriorityQueue) Swap(i, j int)      { q[i], q[j] = q[j], q[i] }

// Push adds a new element to the queue
func (q *PriorityQueue) Push(x interface{}) {
	*q = append(*q, x.(ValueNode))
}

// Pop removes the first element from the queue
func (q *PriorityQueue) Pop() interface{} {
	old := *q
	n := len(old)
	item := old[n-1]
	*q = old[0 : n-1]
	return item
}

// Get returns node at index i (NodeGetter interface)
func (q *PriorityQueue) Get(i int) ValueNode { return (*q)[i] }

func (q *PriorityQueue) Put(x interface{}) { heap.Push(q, x) }
func (q *PriorityQueue) Take() interface{} { return heap.Pop(q) }

// Queue for A* search
type AStarQueue struct {
	nodes   []ValueNode
	fValues []int
}

func NewAStarQueue() *AStarQueue {
	q := &AStarQueue{
		nodes:   make([]ValueNode, 0),
		fValues: make([]int, 0),
	}
	heap.Init(q)
	return q
}

func (q *AStarQueue) Put(x interface{}) { heap.Push(q, x) }
func (q *AStarQueue) Take() interface{} { return heap.Pop(q) }

// Get returns node at index i (NodeGetter interface)
func (q *AStarQueue) Get(i int) ValueNode { return q.nodes[i] }

func (q AStarQueue) Len() int           { return len(q.nodes) }
func (q AStarQueue) Less(i, j int) bool { return q.fValues[i] < q.fValues[j] }
func (q AStarQueue) Swap(i, j int) {
	q.fValues[i], q.fValues[j] = q.fValues[j], q.fValues[i]
	q.nodes[i], q.nodes[j] = q.nodes[j], q.nodes[i]
}

// Push adds a new element to the queue
func (q *AStarQueue) Push(x interface{}) {
	n := x.(ValueNode)
	(*q).nodes = append((*q).nodes, n)
	(*q).fValues = append((*q).fValues, n.G()+n.H())
}

// Pop removes the first element from the queue
func (q *AStarQueue) Pop() interface{} {
	old := (*q).nodes
	n := len(old)
	item := old[n-1]
	(*q).nodes = old[0 : n-1]
	(*q).fValues = (*q).fValues[0 : n-1]
	return item
}

// Queue for A* search
type PriorityQueueFIFO struct {
	nodes []ValueNode
	order []int
	cnt   int
}

func NewPriorityQueueFIFO() *PriorityQueueFIFO {
	q := &PriorityQueueFIFO{
		nodes: make([]ValueNode, 0),
		order: make([]int, 0),
	}
	heap.Init(q)
	return q
}

func (q *PriorityQueueFIFO) Put(x interface{}) { heap.Push(q, x) }
func (q *PriorityQueueFIFO) Take() interface{} { return heap.Pop(q) }

// Get returns node at index i (NodeGetter interface)
func (q *PriorityQueueFIFO) Get(i int) ValueNode { return q.nodes[i] }

func (q PriorityQueueFIFO) Len() int { return len(q.nodes) }
func (q PriorityQueueFIFO) Less(i, j int) bool {
	if q.nodes[i].H() == q.nodes[j].H() {
		return q.order[i] < q.order[j]
	}
	return q.nodes[i].H() < q.nodes[j].H()
}
func (q PriorityQueueFIFO) Swap(i, j int) {
	q.order[i], q.order[j] = q.order[j], q.order[i]
	q.nodes[i], q.nodes[j] = q.nodes[j], q.nodes[i]
}

// Push adds a new element to the queue
func (q *PriorityQueueFIFO) Push(n interface{}) {
	(*q).cnt += 1
	(*q).nodes = append((*q).nodes, n.(ValueNode))
	(*q).order = append((*q).order, (*q).cnt)
}

// Pop removes the first element from the queue
func (q *PriorityQueueFIFO) Pop() interface{} {
	old := (*q).nodes
	n := len(old)
	item := old[n-1]
	(*q).nodes = old[0 : n-1]
	(*q).order = (*q).order[0 : n-1]
	return item
}

// BackupQueue implements a PQ for node backups ordered by depth in the tree
type BackupQueue struct {
	nodes  []Node
	depths []int
	cache  map[Node]struct{}
}

// NewBackupQueue returns a new DmtQueue
func NewBackupQueue() *BackupQueue {
	q := &BackupQueue{
		nodes:  make([]Node, 0),
		depths: make([]int, 0),
		cache:  make(map[Node]struct{}, 0),
	}
	heap.Init(q)
	return q
}

func (q *BackupQueue) Put(x interface{}) { heap.Push(q, x) }
func (q *BackupQueue) Take() interface{} { return heap.Pop(q) }

// clear resets and re-initializes a DmtQueue
func (q *BackupQueue) Clear() {
	q.nodes = q.nodes[:0]
	q.depths = q.depths[:0]
	heap.Init(q)
}

// Len returns the number of elements in the queue
func (q BackupQueue) Len() int { return len(q.nodes) }

// TODO: comment
func (q BackupQueue) Less(i, j int) bool { return q.depths[i] > q.depths[j] }

// Swap swaps the position of the elements at the provided indices in the queue
func (q BackupQueue) Swap(i, j int) {
	q.nodes[i], q.nodes[j] = q.nodes[j], q.nodes[i]
	q.depths[i], q.depths[j] = q.depths[j], q.depths[i]

}

// Push adds a new element to the queue
func (q *BackupQueue) Push(x interface{}) {
	if x == nil {
		return
	}
	n := x.(Node)
	_, ok := q.cache[n]
	if !ok {
		(*q).nodes = append((*q).nodes, n)
		(*q).depths = append((*q).depths, DistanceToRoot(n))
		q.cache[x.(Node)] = struct{}{}
	}
}

// Pop removes the first element from the queue
func (q *BackupQueue) Pop() interface{} {
	old := *q
	n := len(old.nodes)
	node := old.nodes[n-1]
	(*q).nodes = old.nodes[0 : n-1]
	(*q).depths = old.depths[0 : n-1]
	delete(q.cache, node)
	return node
}
