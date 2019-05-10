package search

import (
	"container/heap"
	"testing"

	"github.com/schultet/goa/src/state"
	"github.com/schultet/goa/src/task"
)

// TestNode is a generic data type satisfying the Node interface
type TestNode struct {
	id     int
	g      int
	h      int
	parent *TestNode
}

func (n *TestNode) ID() state.StateID    { return state.StateID(n.id) }
func (n *TestNode) G() int               { return n.g }
func (n *TestNode) H() int               { return n.h }
func (n *TestNode) Action() *task.Action { return nil }
func (n *TestNode) Parent() Node {
	if n.parent == nil {
		return nil
	}
	return n.parent
}

// createTestNodes creates a set of nodes
func createTestNodes() []*TestNode {
	n0 := &TestNode{id: 0}
	n1 := &TestNode{parent: n0, id: 1}
	n2 := &TestNode{parent: n1, id: 2}
	n3 := &TestNode{parent: n1, id: 3}
	n4 := &TestNode{parent: n3, id: 4}

	return []*TestNode{n0, n1, n2, n3, n4}
}

type TestCase struct {
	insertionOrder []int
	expectedOrder  []int
}

func runTests(t *testing.T, nodes []*TestNode, tests []TestCase, queue heap.Interface) {
	for i, tc := range tests {
		for _, v := range tc.insertionOrder {
			heap.Push(queue, nodes[v])
		}

		for _, n := range tc.expectedOrder {
			if nn := heap.Pop(queue).(Node); int(nn.ID()) != n {
				t.Errorf("TC: %v\nEXP: %v\nWAS: %v\n", i, n, nn)
			}
		}
	}
}

func TestBackupQueue(t *testing.T) {
	nodes := createTestNodes()
	tests := []TestCase{
		{[]int{1, 2, 4, 0}, []int{4, 2, 1, 0}},
		{[]int{4, 2, 1, 0}, []int{4, 2, 1, 0}},
		{[]int{1, 4, 0, 2}, []int{4, 2, 1, 0}},
		{[]int{1, 1, 2, 2}, []int{2, 1}},
		{[]int{1, 3, 4}, []int{4, 3, 1}},
	}

	q := NewBackupQueue()
	heap.Init(q)
	runTests(t, nodes, tests, q)
}

func TestPriorityQueue(t *testing.T) {
	nodes := []*TestNode{
		{id: 0, h: -1},
		{id: 1, h: 0},
		{id: 2, h: 22},
		{id: 3, g: 1, h: 32},
		{id: 4, g: 0, h: 33},
		{id: 5, g: 3, h: 34},
		{id: 6, g: 2, h: 35},
		{id: 7, h: 77},
	}
	tests := []TestCase{
		{[]int{0, 1, 2, 3, 4, 5, 6, 7}, []int{0, 1, 2, 3, 4, 5, 6, 7}},
		{[]int{1, 1, 1, 1, 1, 1, 1, 1}, []int{1, 1, 1, 1, 1, 1, 1, 1}},
		{[]int{0, 1, 2, 3, 0, 1, 2, 3}, []int{0, 0, 1, 1, 2, 2, 3, 3}},
		{[]int{0}, []int{0}},
		{[]int{2, 1}, []int{1, 2}},
	}

	q := NewPriorityQueue()
	heap.Init(q)
	runTests(t, nodes, tests, q)
}

func TestAStarQueue(t *testing.T) {
	nodes := []*TestNode{
		{id: 0, g: 6, h: 0},  // f= 6
		{id: 1, g: 4, h: 4},  // f= 8
		{id: 2, g: 3, h: 22}, // f=25
		{id: 3, g: 1, h: 32}, // f=33
		{id: 4, g: 0, h: 34}, // f=34
		{id: 5, g: 0, h: 35}, // f=35
		{id: 6, g: 5, h: 35}, // f=40
	}
	tests := []TestCase{
		{[]int{6, 5, 4, 3, 2, 1, 0}, []int{0, 1, 2, 3, 4, 5, 6}},
		{[]int{0, 1, 2, 3, 4, 5, 6}, []int{0, 1, 2, 3, 4, 5, 6}},
		{[]int{0}, []int{0}},
		{[]int{0, 0, 0}, []int{0, 0, 0}},
		{[]int{1, 0}, []int{0, 1}},
		{[]int{1, 2, 1, 0}, []int{0, 1, 1, 2}},
	}

	q := NewAStarQueue()
	heap.Init(q)
	runTests(t, nodes, tests, q)
}
