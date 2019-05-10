package search

import (
	"fmt"
	"testing"

	"github.com/schultet/goa/src/state"
	"github.com/schultet/goa/src/task"
)

func newSimpleTask() *task.Task {
	init := state.State{0, 0, 0}
	goal := []task.Condition{{Variable: 0, Value: 1}}
	vars := []*task.Variable{
		{ID: 0, Name: "pos", DomRange: 2},
		{ID: 1, Name: "left", DomRange: 2, IsPrivate: true},
		{ID: 2, Name: "right", DomRange: 2, IsPrivate: true},
	}
	ops := task.ActionList{
		{ID: 0, Preconditions: []task.Condition{{1, 0}}, Effects: []task.Effect{{1, 1}}, Private: true},
		{ID: 1, Preconditions: []task.Condition{{2, 0}}, Effects: []task.Effect{{2, 1}}, Private: true},
		{ID: 2, Preconditions: []task.Condition{{0, 0}, {2, 1}, {1, 1}}, Effects: []task.Effect{{0, 1}}},
	}
	return &task.Task{
		AgentID: 0,
		Actions: task.ActionMap{0: ops},
		Vars:    vars,
		Init:    init,
		Goal:    goal,
	}
}

func TestNewStubbornSets(t *testing.T) {
	task := newSimpleTask()
	s := NewStubbornSetsSimple(task)
	fmt.Printf("Precons: %v\n", s.sortedOpPreconditions)
	fmt.Printf("Effects: %v\n", s.sortedOpEffects)
	fmt.Printf("OpIDs: %v\n", s.relativeOpID)
	fmt.Printf("Achievers: %v\n", s.achievers)
	fmt.Printf("applicable = %+v\n", s.applicablePublic(&task.Init))
	fmt.Printf("numPublicVars = %+v\n", s.numPublicVars)

}
