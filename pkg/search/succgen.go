package search

import (
	"fmt"
	//"log"

	"github.com/schultet/goa/pkg/state"
	"github.com/schultet/goa/pkg/task"
)

// An ActionEvaluator can compute applicable actions
type ActionEvaluator interface {
	ApplicableActions(s state.State) task.ActionList
	ApplicableAgents(s state.State) []int
}

// TODO: actionEvaluator can be optimized: instead of having an actionGraph for
// each agent, have one actionGraph containing all agents actions.
// Computing ApplicableAgents then needs only one sweep rather than
// n (one for each agent)
type actionEvaluator struct {
	id        int
	succGens  map[int]*actionGraph
	agSuccGen *actionGraph
}

// NewActionEvaluator creates and returns an appTester
func NewActionEvaluator(vars []*task.Variable, am task.ActionMap, id int) ActionEvaluator {
	a := actionEvaluator{
		id:       id,
		succGens: make(map[int]*actionGraph),
	}
	for i, actions := range am {
		if i != id {
			a.succGens[i] = newActionGraph(vars, actions, []*task.Condition{})
		} else {
			a.agSuccGen = newActionGraph(vars, actions, []*task.Condition{})
		}
	}
	return &a
}

// returns a list of agentIDs for agents that have applicable actions in state
// s, excluding the planning agent.
func (eval actionEvaluator) ApplicableAgents(s state.State) []int {
	var ids []int
	for id, succgen := range eval.succGens {
		if succgen.ApplicableActionsExist(s) {
			ids = append(ids, id)
		}
	}
	return ids
}

func (eval actionEvaluator) ApplicableActions(s state.State) task.ActionList {
	return eval.agSuccGen.ApplicableActions(s)
}

// pruningEvaluator combines evaluation of applicable actions with pruning
type pruningEvaluator struct {
	ActionEvaluator
	ap ActionPruner
}

func (e pruningEvaluator) ApplicableActions(s state.State) task.ActionList {
	return e.ap.Prune(&s, e.ApplicableActions(s))
}

func NewPruningEvaluator(am ActionEvaluator, ap ActionPruner) ActionEvaluator {
	return &pruningEvaluator{am, ap}
}

// actionGraph is a graph structure to efficiently compute applicable
// actions for a given state
type actionGraph struct {
	variable int
	appOps   task.ActionList
	edges    []*actionGraph
}

// newActionGraph recursively builds a successor generator graph. To build
// the structure, a list of variables and a list of actions is required.
func newActionGraph(vars []*task.Variable, os task.ActionList,
	cs []*task.Condition) *actionGraph {
	if 0 == len(vars) || 0 == len(os) {
		app, _ := os.Satisfacted(cs)
		if len(app) == 0 {
			return nil
		}
		return &actionGraph{-1, app, make([]*actionGraph, 0)}
	}
	app, unapp := os.Satisfacted(cs)
	v, vs := vars[0], vars[1:]

	sg := actionGraph{
		variable: v.ID,
		appOps:   app,
		edges:    make([]*actionGraph, v.DomRange+1),
	}

	for i := 0; i < int(v.DomRange); i++ {
		newCs := make([]*task.Condition, len(cs)+1)
		copy(newCs, cs)
		newCs[len(cs)] = &task.Condition{Variable: v.ID, Value: i}
		app, unapp = unapp.WithCondition(task.Condition{
			Variable: v.ID, Value: i})
		sgg := newActionGraph(vs, app, newCs)
		sg.edges[i] = sgg
	}
	sg.edges[v.DomRange] = newActionGraph(vs, unapp, cs)

	return &sg
}

// printInorder prints a successorGenerator to stdout
func (sg *actionGraph) printInorder() {
	fmt.Println(*sg)
	for _, s := range sg.edges {
		if s != nil {
			s.printInorder()
		}
	}
}

// ApplicableActions returns the list of actions that are applicable in state s
func (sg *actionGraph) ApplicableActions(s state.State) (ops task.ActionList) {
	if sg == nil {
		ops = make(task.ActionList, 0)
		return
	}
	ops = make(task.ActionList, len(sg.appOps))
	copy(ops, sg.appOps)
	if sg.variable == -1 || len(sg.edges) == 0 {
		return
	}

	val := s[sg.variable]
	if len(sg.edges) > val && sg.edges[val] != nil {
		ops = append(ops, sg.edges[val].ApplicableActions(s)...)
	}
	if sg.edges[len(sg.edges)-1] != nil {
		ops = append(ops, sg.edges[len(sg.edges)-1].ApplicableActions(s)...)
	}

	return
}

// ApplicableActionsExist returns true iff at least one applicable actions exists for s
func (sg *actionGraph) ApplicableActionsExist(s state.State) bool {
	//fmt.Printf("sg = %+v\n", sg)
	if len(sg.appOps) > 0 {
		return true
	}
	if sg.variable == -1 || len(sg.edges) == 0 {
		return false
	}

	val := s[sg.variable]
	app := false
	if len(sg.edges) > int(val) && sg.edges[val] != nil {
		app = sg.edges[val].ApplicableActionsExist(s)
		if app {
			return true
		}
	}
	if sg.edges[len(sg.edges)-1] != nil {
		return sg.edges[len(sg.edges)-1].ApplicableActionsExist(s)
	}
	return false
}
