package search

import (
	"testing"

	"github.com/schultet/goa/src/pddl"
	"github.com/schultet/goa/src/state"
	"github.com/schultet/goa/src/task"
	"github.com/schultet/goa/src/util/slice"
)

func TestApplicableOps(t *testing.T) {
	vars := []*task.Variable{
		{ID: 0, DomRange: 2},
		{ID: 1, DomRange: 2},
		{ID: 2, DomRange: 2},
	}
	ops := task.ActionList([]*task.Action{
		{PDDLAction: pddl.DummyAction("A"), Preconditions: []task.Condition{}},
		{PDDLAction: pddl.DummyAction("B"), Preconditions: []task.Condition{
			{Variable: 2, Value: 1}}}, //
		{PDDLAction: pddl.DummyAction("C"), Preconditions: []task.Condition{
			{Variable: 0, Value: 1}}},
		{PDDLAction: pddl.DummyAction("D"), Preconditions: []task.Condition{
			{Variable: 0, Value: 1}}},
		{PDDLAction: pddl.DummyAction("E"), Preconditions: []task.Condition{
			{Variable: 0, Value: 1},
			{Variable: 1, Value: 0}}},
		{PDDLAction: pddl.DummyAction("F"), Preconditions: []task.Condition{
			{Variable: 0, Value: 0},
			{Variable: 1, Value: 0}}},
		{PDDLAction: pddl.DummyAction("G"), Preconditions: []task.Condition{
			{Variable: 0, Value: 0},
			{Variable: 1, Value: 0},
			{Variable: 2, Value: 0}}}, //
	})
	sg := newActionGraph(vars, ops, []*task.Condition{})
	//sg.printInorder()
	testStates := []state.State{
		{0, 0, 1},
		{0, 1, 0},
		{1, 1, 1},
		{1, 0, 1},
		{1, 0, 0},
	}
	expRes := [][]string{
		{"A", "B", "F"},
		{"A"},
		{"A", "B", "C", "D"},
		{"A", "B", "C", "D", "E"},
		{"A", "C", "D", "E"},
	}
	for j, s := range testStates {
		appOps := sg.ApplicableActions(s)
		opNames := make([]string, len(appOps))
		for i, o := range appOps {
			opNames[i] = o.Name()
		}
		if !slice.ContainSameStrings(opNames, expRes[j]) {
			t.Errorf("wrong applicable ops generated!\nexp:%s\nwas:%s\n", expRes[j], opNames)
		}
	}
}
