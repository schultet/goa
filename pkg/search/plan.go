package search

import (
	"fmt"
	"log"

	"github.com/schultet/goa/pkg/comm"
	"github.com/schultet/goa/pkg/state"
	"github.com/schultet/goa/pkg/task"
	"github.com/schultet/goa/pkg/util/ints"
)

// Move represents the application of an action at a certain point in time
type Move struct {
	T      int // discrete time
	Action *task.Action
}

// Plan consists of Moves of an agent
type Plan []Move

// Cost returns the sum of all action costs of the plan
func (p Plan) Cost(ct task.CostType) int {
	cost := 0
	for _, m := range p {
		cost += m.Action.AdjustedCost(ct)
	}
	return cost
}

// Reverse reverses a plan given the total number of steps among all agents
// (each agents plan might be shorter).
func (p Plan) Reversed(maxStep int) Plan {
	res := make(Plan, len(p))
	for i, m := range p {
		res[len(res)-i-1] = Move{maxStep - m.T, m.Action}
	}
	return res
}

//func (p Plan) Valid(e *Engine) bool {
//	return false
//}

func (p Plan) String() string {
	s := ""
	for _, m := range p {
		s += fmt.Sprintf("(%d:%s)\n", m.T, m.Action)
	}
	return s
}

// getMakespan of a plan with move costs according to the given cost type.
func (p Plan) getMakespan(c task.CostType) int {
	var actions []*task.Action
	var costs []int
	for _, move := range p {
		actions = append(actions, move.Action)
		costs = append(costs, move.Action.AdjustedCost(c))
	}
	return makespanApprox(actions, costs)
}

// XMessage = plan extraction message
type XMessage struct {
	SenderID      int
	Descriptor    string // = AgentID+GoalSID of extraction initiator
	CurrentState  []int
	CurrentStep   int
	TotalCost     int
	TotalMakespan int
	Msg           string
}

type PlanRegistry struct {
	e            *Engine
	plans        map[string]*PlanExtractor
	bestMakespan int
	bestCost     int
}

func NewPlanRegistry(e *Engine) *PlanRegistry {
	return &PlanRegistry{
		e:            e,
		plans:        make(map[string]*PlanExtractor),
		bestMakespan: ints.MaxValue,
		bestCost:     ints.MaxValue,
	}
}

func (pr *PlanRegistry) NumPlans() int {
	n := 0
	for _, plan := range pr.plans {
		if plan.done {
			n += 1
		}
	}
	return n
}

func (pr *PlanRegistry) BestCostPlan() Plan {
	for _, px := range pr.plans {
		if px.totalCost == pr.bestCost {
			return px.plan
		}
	}
	return nil
}

func (pr *PlanRegistry) BestMakespanPlan() (string, Plan) {
	for _, px := range pr.plans {
		if px.totalMakespan == pr.bestMakespan {
			return px.descriptor, px.plan
		}
	}
	return "", nil
}

// Start extracting a plan leading to the goal node n.
func (pr *PlanRegistry) Start(n Node, costtype task.CostType) {
	descriptor := fmt.Sprintf("%d-%d", pr.e.agentID, n.ID())

	_, ok := pr.plans[descriptor]
	if ok {
		// FIXME: Equal states have the same descriptors when pruning is enabled
		//log.Fatalf("ERROR: Multiple plans to state %s found\n", descriptor)
		return
	}

	// TODO: type switch should not be necessary
	// there has to be a more generic way...
	c := -1
	switch v := n.(type) {
	case *MafsNode:
		c = v.g
	case *DmtMakespanNode:
		c = v.distance(costtype)
	case *DmtNode:
		c = v.distance(costtype)
	}

	m := -1
	switch v := n.(type) {
	case *DmtMakespanNode:
		m = v.m
	}

	// create new PlanExtractor
	log.Printf("Start plan extraction: %s (%v)\n", descriptor, m)
	pr.plans[descriptor] = &PlanExtractor{
		e:             pr.e,
		plan:          make(Plan, 0),
		descriptor:    descriptor,
		totalCost:     c,
		totalMakespan: m,
	}
	makespan := pr.plans[descriptor].totalMakespan
	cost := pr.plans[descriptor].totalCost
	if makespan < pr.bestMakespan {
		pr.bestMakespan = makespan
	}
	if cost < pr.bestCost {
		pr.bestCost = cost
	}
	pr.plans[descriptor].Extract(n)
}

func (pr *PlanRegistry) processMessage(m *XMessage) {
	_, ok := pr.plans[m.Descriptor]
	if !ok {
		pr.plans[m.Descriptor] = &PlanExtractor{
			e:             pr.e,
			plan:          make(Plan, 0),
			descriptor:    m.Descriptor,
			totalCost:     m.TotalCost,
			totalMakespan: m.TotalMakespan,
		}
		makespan := pr.plans[m.Descriptor].totalMakespan
		cost := pr.plans[m.Descriptor].totalCost
		if makespan < pr.bestMakespan {
			pr.bestMakespan = makespan
		}
		if cost < pr.bestCost {
			pr.bestCost = cost
		}
	}
	pr.plans[m.Descriptor].processMessage(m)
}

type PlanExtractor struct {
	e             *Engine
	plan          Plan
	descriptor    string
	state         state.SharedState
	step          int
	totalCost     int
	totalMakespan int
	done          bool
}

func (x *PlanExtractor) Msg() *XMessage {
	return &XMessage{
		SenderID:      x.e.agentID,
		CurrentStep:   x.step,
		CurrentState:  x.state,
		Descriptor:    x.descriptor,
		TotalCost:     x.totalCost,
		TotalMakespan: x.totalMakespan,
	}
}

func (x *PlanExtractor) processMessage(m *XMessage) {
	if m.Msg == "done" {
		x.done = true
		x.plan = x.plan.Reversed(m.CurrentStep)

		x.e.info.SetPlan(x.e.planRegistry.plans[x.descriptor])
		x.e.logger.Print(x.e.info.SummaryJSON())

		return
	}

	sid := x.e.stateRegistry.Token(m.CurrentState) // TODO: rename Token(...) to GetStateID(...)
	n, _ := x.e.Strategy.(NodeAccessor).HasNode(sid)
	x.step = m.CurrentStep
	x.Extract(n)
}

// Extract the plan steps of this agent, then ask other agents for their parts.
func (x *PlanExtractor) Extract(n Node) {
	// extract locally
	for n.Action() != nil && n.Action().Owner == x.e.agentID {
		x.plan = append(x.plan, Move{T: x.step, Action: n.Action()})
		n, x.step = n.Parent(), x.step+1
	}
	x.state = x.e.stateRegistry.SharedState(x.e.stateRegistry.Lookup(n.ID()), n.ID())

	// reached the root -> inform others -> done.
	if n.Action() == nil {
		msg := x.Msg()
		msg.Msg = "done"

		go comm.Broadcast(
			x.e.dispatcher,
			msg,
			x.e.conns.Except(int(x.e.agentID)),
		)

		x.done = true
		x.plan = x.plan.Reversed(x.step)

		x.e.info.SetPlan(x.e.planRegistry.plans[x.descriptor])
		x.e.logger.Print(x.e.info.SummaryJSON())

		return
	}

	// someone else shall continueth
	if n.Action().Owner != x.e.agentID {
		x.e.dispatcher.Send(n.Action().Owner, x.Msg())
		return
	}
	log.Fatalf("ERROR: Plan extraction most likely broken.")
}

func (e *Engine) Tokens(n Node) []int { //state.TokenArray {
	switch nodet := n.(type) {
	case *MafsNode:
		return e.stateRegistry.(*state.GstateReg).Tokens(e.stateRegistry.Lookup(nodet.stateID))
	case *DmtNode:
		return e.stateRegistry.(*state.GstateReg).Tokens(e.stateRegistry.Lookup(nodet.stateID))
	case *DmtMakespanNode:
		return e.stateRegistry.(*state.GstateReg).Tokens(e.stateRegistry.Lookup(nodet.stateID))
		//.Add( int32(nodet.ID()), e.agentID)
		// TODO: ^ this is strange, check
	}
	return nil
}
