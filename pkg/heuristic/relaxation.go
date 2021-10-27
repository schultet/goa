package heuristic

import (
	"container/heap"

	"github.com/schultet/goa/pkg/comm"
	"github.com/schultet/goa/pkg/state"
	"github.com/schultet/goa/pkg/task"
)

type relaxedFact struct {
	variable       int
	value          int
	isGoal         bool
	preconditionOf []*relaxedAction // actions, this fact is a precondition of
	cost           int
	reachedBy      *relaxedAction
	marked         bool
}

type relaxedAction struct {
	src            *task.Action // original (unrelaxed) action
	id             int
	pre            []int // set of facts (indices)
	eff            []int // set of facts (indices)
	cost           int
	unsatisfiedPre int // number of preconditions not satisfied
	tmpCost        int // used hMax or hAdd cost, incl. action cost
	marked         bool
}

// Priority Queue for relaxed plan exploration
type explorationQueue []*factDistancePair

type factDistancePair struct {
	fact     *relaxedFact
	distance int
}

func (p explorationQueue) Len() int { return len(p) }

func (p explorationQueue) Less(i, j int) bool {
	return p[i].distance < p[j].distance
}
func (p explorationQueue) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

func (p *explorationQueue) Push(x interface{}) {
	*p = append(*p, x.(*factDistancePair))
}

func (p *explorationQueue) Pop() interface{} {
	old := *p
	n := len(old)
	item := old[n-1]
	*p = old[0 : n-1]
	return item
}

// relaxation heuristics solve a relaxed planning task to compute h-values
type relaxationHeuristic struct {
	agentID     int
	numFacts    int
	goal        []int
	facts       []relaxedFact
	factIndex   [][]int // factIndex[var][val] -> index of corresp. fact
	actions     []relaxedAction
	queue       explorationQueue
	relaxedPlan []bool
	prefOps     []*task.Action
	costtype    task.CostType
}

// TODO: doc
type distributedFFHeuristic struct {
	server     comm.Server
	dispatcher comm.Dispatcher
	// TODO: other members
}

// TODO: doc
func DistributedFFHeuristic(t *task.Task, s comm.Server, d comm.Dispatcher) *distributedFFHeuristic {
	return &distributedFFHeuristic{
		server:     s,
		dispatcher: d,
		// initialize other members
	}
	// TODO: implement
}

// TODO: doc
func (h *distributedFFHeuristic) Evaluate(s state.State) int {
	// TODO: implement
	//h.dispatcher.Send(AgentId, m)
	return -1
}

func RelaxationHeuristic(t *task.Task, costtype task.CostType) *relaxationHeuristic {
	r := &relaxationHeuristic{
		agentID:  t.AgentID,
		queue:    make(explorationQueue, 0),
		prefOps:  make([]*task.Action, 0),
		costtype: costtype,
	}
	// TODO: rt := NewRelaxedTask(t)
	heap.Init(&r.queue)
	r.computeFacts(t) // including goal-facts
	r.numFacts = len(r.facts)
	r.computeRelaxedActions(t)
	r.relaxedPlan = make([]bool, len(r.actions))
	return r
}

// computes the list of fdr facts (v=d) and the factIndex
func (r *relaxationHeuristic) computeFacts(t *task.Task) {
	numFacts := 0
	for _, v := range t.Vars {
		numFacts += int(v.DomRange)
	}
	curFact := 0
	r.facts = make([]relaxedFact, numFacts)
	r.factIndex = make([][]int, len(t.Vars))
	for i, variable := range t.Vars {
		r.factIndex[i] = make([]int, variable.DomRange)
		for j := 0; j < int(variable.DomRange); j++ {
			r.factIndex[i][j] = curFact
			r.facts[curFact] = relaxedFact{
				variable: variable.ID,
				value:    j,
				cost:     -1,
			}
			curFact += 1
		}
	}
	r.goal = make([]int, len(t.Goal))
	for i, c := range t.Goal {
		r.goal[i] = r.factIndex[c.Variable][c.Value]
		r.facts[r.goal[i]].isGoal = true
	}
}

// computes relaxed actions
func (r *relaxationHeuristic) computeRelaxedActions(t *task.Task) {
	numActions := 0
	for _, ops := range t.Actions {
		numActions += len(ops)
	}
	r.actions = make([]relaxedAction, numActions)
	var (
		curFact int
		curOp   int = 0
	)
	for _, ops := range t.Actions {
		for i, o := range ops {
			r.actions[curOp] = relaxedAction{
				src:            ops[i],
				id:             curOp,
				pre:            make([]int, len(o.Preconditions)),
				eff:            make([]int, len(o.Effects)),
				cost:           o.AdjustedCost(r.costtype),
				unsatisfiedPre: len(o.Preconditions),
			}
			for j, f := range o.Preconditions {
				curFact = r.factIndex[f.Variable][f.Value]
				r.actions[curOp].pre[j] = curFact
				r.facts[curFact].preconditionOf = append(
					r.facts[curFact].preconditionOf, &r.actions[curOp])
			}
			for j, f := range o.Effects {
				r.actions[curOp].eff[j] = r.factIndex[f.Variable][f.Value]
			}
			curOp += 1
		}
	}
}

func (h *relaxationHeuristic) enqueueIfNecessary(f *relaxedFact, cost int,
	op *relaxedAction) {
	if f.cost == -1 || f.cost > cost {
		f.cost = cost
		f.reachedBy = op
		heap.Push(&h.queue, &factDistancePair{fact: f, distance: cost})
	}
}

func (h *relaxationHeuristic) increaseCost(cost *int, amount int) {
	*cost += amount
	if *cost > MaxCost {
		*cost = MaxCost
	}
}

func (h *relaxationHeuristic) setupExplorationQueue() {
	h.queue = h.queue[:0]
	heap.Init(&h.queue)

	for i := range h.facts {
		h.facts[i].cost = -1
		h.facts[i].marked = false
	}
	for i := range h.actions {
		h.actions[i].unsatisfiedPre = len(h.actions[i].pre)
		h.actions[i].tmpCost = h.actions[i].cost
		h.actions[i].marked = false

		if h.actions[i].unsatisfiedPre == 0 {
			for _, eff := range h.actions[i].eff {
				h.enqueueIfNecessary(&h.facts[eff], h.actions[i].cost,
					&h.actions[i])
			}
		}
	}
}

func (h *relaxationHeuristic) setupExplorationQueueState(s state.State) {
	//fmt.Printf("%v\n%v\n%v\n%d/%d/%d\n", s, h.facts, h.factIndex, len(s),
	// len(h.facts), len(h.factIndex))
	for v, val := range s[:len(h.factIndex)] {
		//fmt.Printf("{%v,%v}\n", v, val)
		h.enqueueIfNecessary(&h.facts[h.factIndex[v][val]], 0, nil)
	}
}

func (h *relaxationHeuristic) explore() {
	unsolvedGoals := len(h.goal)
	for h.queue.Len() != 0 {
		fdpair := heap.Pop(&h.queue).(*factDistancePair)
		d := fdpair.distance
		f := fdpair.fact
		factCost := f.cost
		if factCost < d {
			continue
		}
		if f.isGoal {
			unsolvedGoals -= 1
			if unsolvedGoals == 0 {
				return
			}
		}
		triggeredOps := f.preconditionOf
		for _, o := range triggeredOps {
			h.increaseCost(&o.tmpCost, factCost)
			o.unsatisfiedPre -= 1
			if o.unsatisfiedPre == 0 {
				for _, eff := range o.eff {
					h.enqueueIfNecessary(&h.facts[eff], o.tmpCost, o)
				}
			}
		}
	}
}

func (h *relaxationHeuristic) computeAddAndFF(s state.State) int {
	h.setupExplorationQueue()
	h.setupExplorationQueueState(s)
	h.explore()

	totalCost := 0
	for _, i := range h.goal {
		factCost := h.facts[i].cost
		if factCost == -1 {
			return DeadEnd
		}
		h.increaseCost(&totalCost, factCost)
	}
	return totalCost
}

// actionApplicable returns true iff action a is applicable in state s
func actionApplicable(a *task.Action, s state.State) bool {
	for _, c := range a.Preconditions {
		if s[c.Variable] != int(c.Value) {
			return false
		}
	}
	return true
}

func (h *relaxationHeuristic) markPrefOpsAndRelaxedPlan(s state.State, goal *relaxedFact) {
	if goal.marked {
		return
	}
	goal.marked = true
	o := goal.reachedBy
	if o == nil {
		return
	}
	for _, pre := range o.pre {
		h.markPrefOpsAndRelaxedPlan(s, &h.facts[pre])
	}
	if o.id != -1 {
		h.relaxedPlan[o.id] = true
		if o.cost == o.tmpCost && !o.marked && actionApplicable(o.src, s) {
			o.marked = true
			h.prefOps = append(h.prefOps, o.src)
		}
	}
}

// Bellman-Ford computation of hadd
func (h *relaxationHeuristic) Evaluate(s state.State) int {
	t := make([]int, len(h.facts)+len(h.actions))
	for i := 0; i < len(t); i++ {
		t[i] = MaxCost
	}
	for v, val := range s[:len(h.factIndex)] {
		t[h.factIndex[v][val]] = 0
	}

	changed := true
	for changed {
		changed = false
		for i, op := range h.actions {
			xo := SumOfInts(t, op.pre...)
			if xo != t[h.numFacts+i] {
				t[h.numFacts+i] = xo
				for _, p := range op.eff {
					if 1+xo < 0 || t[p] < 1+xo {
						t[p] = t[p]
					} else {
						t[p] = 1 + xo // TODO: use action cost & costtype
					}
				}
				changed = true
			}
		}
	}
	return SumOfInts(t, h.goal...)
}

// additive heuristic (hAdd) with dijkstra computation
type additiveHeuristic struct {
	*relaxationHeuristic
}

func AdditiveHeuristic(t *task.Task) StateEvaluator {
	return &additiveHeuristic{relaxationHeuristic: RelaxationHeuristic(t, task.NormalMinOne)}
}

func (h *additiveHeuristic) Evaluate(s state.State) int {
	h.prefOps = h.prefOps[:0] // clear preferred actions
	hvalue := h.computeAddAndFF(s)
	if hvalue == DeadEnd {
		return MaxCost
	}
	return hvalue
}

// additive heuristic (hAdd) with bellman ford computation
func AdditiveBFHeuristic(t *task.Task) StateEvaluator {
	return RelaxationHeuristic(t, task.NormalMinOne)
}

// ff heuristic
type ffHeuristic struct {
	*relaxationHeuristic
}

func FFHeuristic(t *task.Task) StateEvaluator {
	return &ffHeuristic{relaxationHeuristic: RelaxationHeuristic(t, task.NormalMinOne)}
}

func (h *ffHeuristic) Evaluate(s state.State) int {
	h.prefOps = h.prefOps[:0]      // clear preferred actions
	hvalue := h.computeAddAndFF(s) // compute additive heuristic value
	if hvalue == DeadEnd {
		return MaxCost
	}

	for _, g := range h.goal {
		h.markPrefOpsAndRelaxedPlan(s, &h.facts[g])
	}

	hvalue = 0
	for i, v := range h.relaxedPlan {
		if v {
			h.relaxedPlan[i] = false    // clean up for next comp
			hvalue += h.actions[i].cost // TODO: adjusted cost
		}
	}
	return hvalue
}

// preferredActions returns preferred actions if the heuristic supports them
// otherwise an empty list is returned.
func preferredActions(h StateEvaluator, agentID int) []*task.Action {
	switch heur := h.(type) {
	case *ffHeuristic:
		agentPrefOps := make([]*task.Action, 0, len(heur.prefOps))
		for _, op := range heur.prefOps {
			if op.Owner == agentID {
				agentPrefOps = append(agentPrefOps, op)
			}
		}
		return agentPrefOps
	default:
		return []*task.Action{}
	}
}

// SumOfInts computes the sum of all t[i] where i in gs
func SumOfInts(t []int, gs ...int) int {
	s := 0
	for _, g := range gs {
		s += t[g]
		if s > MaxCost {
			return MaxCost
		}
	}
	return s
}

// register relaxation heuristics add, ff, addbf (bellman-ford)
func init() {
	Register(&EvaluatorInfo{"add", "additive heuristic", AdditiveHeuristic})
	Register(&EvaluatorInfo{"addbf", "additive heuristic (bellman-ford)",
		AdditiveBFHeuristic})
	Register(&EvaluatorInfo{"ff", "ff-heuristic", FFHeuristic})
}
