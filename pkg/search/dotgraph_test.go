package search

import (
	"testing"

	"github.com/schultet/goa/pkg/pddl"
	"github.com/schultet/goa/pkg/state"
	"github.com/schultet/goa/pkg/task"
)

func TestDotNodeString(t *testing.T) {
	root := &DmtMakespanNode{
		stateID: state.StateID(0),
		h:       100,
	}
	n := &DmtMakespanNode{
		stateID: state.StateID(35),
		parent:  root,
		h:       77,
		costs:   []int{1, 2, 3},
		actions: []*task.Action{
			&task.Action{
				PDDLAction: pddl.DummyAction("move"),
				Args:       []*pddl.Object{pddl.DummyObject("a")},
			},
			&task.Action{
				PDDLAction: pddl.DummyAction("load"),
				Args:       []*pddl.Object{pddl.DummyObject("x"), pddl.DummyObject("y")},
			},
		},
	}
	root.children = []*DmtMakespanNode{n}
	DotGraph(root, "bla")
	//t.Errorf("\n%v\n", DotNodeString(n))
}
