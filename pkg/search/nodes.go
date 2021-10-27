package search

import "github.com/schultet/goa/pkg/state"
import "github.com/schultet/goa/pkg/task"

// Node interface all search nodes should implement
type Node interface {
	ID() state.StateID
	Parent() Node
	Action() *task.Action
}

type ValueNode interface {
	Node
	H() int
	G() int
}

// NodeAccessor contains methods to access nodes
type NodeAccessor interface {
	HasNode(state.StateID) (Node, bool)
	Tokens(Node) state.State
}

// DistanceToRoot computes how many actions away a node is from the root
func DistanceToRoot(n Node) int {
	distance := 0
	for n.Parent() != nil { // n is not root
		distance++
		n = n.Parent()
	}
	return distance
}
