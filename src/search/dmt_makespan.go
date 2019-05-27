// Copyright 2018/2019, University of Freiburg,
// Chair of Artificial Intelligence.
// Authors:
//   - Robert Grönsfeld
//   - Tim Schulte <schultet@informatik.uni-freiburg.de>.
package search

import (
	"encoding/gob"
	"fmt"
	"log"
	"math"
	"math/rand"

	"github.com/schultet/goa/src/heuristic"
	"github.com/schultet/goa/src/opt"
	"github.com/schultet/goa/src/state"
	"github.com/schultet/goa/src/task"
	"github.com/schultet/goa/src/util/floats"
	"github.com/schultet/goa/src/util/ints"
)

const (
	MinInt32 = -1 << 31
)

// DmtMakespanNode implements search.Node
type DmtMakespanNode struct {
	stateID  state.StateID
	parent   *DmtMakespanNode
	children []*DmtMakespanNode
	actions  []*task.Action
	costs    []int
	h        int // heuristic value
	m        int // makespan value
	visits   int
	closed   bool
}

// Action returns the nodes creating action
func (n *DmtMakespanNode) Action() *task.Action {
	if len(n.actions) == 0 {
		return nil
	}
	return n.actions[len(n.actions)-1]
}

// Parent returns the nodes parent node
func (n *DmtMakespanNode) Parent() Node {
	// since we return an interface type, we must explicitly return nil here,
	// even though n.parent == nil it is not the same to return n.parent.
	if n.parent == nil {
		return nil
	}
	return n.parent
}

// ID returns the unique NodeID of the node
func (n *DmtMakespanNode) ID() state.StateID { return n.stateID }

// isRoot returns true iff the node is the root (i.e. has no parent)
func (n *DmtMakespanNode) isRoot() bool { return n.parent == nil }

// String returns a string representation of n
func (n *DmtMakespanNode) String() string {
	return fmt.Sprintf("(sid:%d, actions:%v par:%s)", n.stateID, n.actions, n.parent)
}

// removeChild removes child c from n's children
func (n *DmtMakespanNode) removeChild(c *DmtMakespanNode) {
	for i, child := range n.children {
		if child == c {
			last := n.children[len(n.children)-1]
			n.children[i] = last
			n.children = n.children[:len(n.children)-1]
			return
		}
	}
}

// addChild adds a child to n if its not a forbear of n
func (n *DmtMakespanNode) addChild(c *DmtMakespanNode) {
	n.children = append(n.children, c)
}

// forbearOf checks whether n is a forbear (=ancestor) of other
func (n *DmtMakespanNode) forbearOf(other *DmtMakespanNode) bool {
	for !other.isRoot() {
		if other.parent == n {
			return true
		}
		other = other.parent
	}
	return false
}

// pathToRoot returns all actions along the path to the root
func (n *DmtMakespanNode) pathToRoot() ([]*task.Action, []int) {
	actions := []*task.Action{}
	costs := []int{}
	curNode := n
	for !curNode.isRoot() {
		actions = append(curNode.actions, actions...)
		costs = append(curNode.costs, costs...)
		curNode = curNode.parent
	}
	root := curNode
	actions = append(root.actions, actions...)
	costs = append(root.costs, costs...)
	return actions, costs
}

// DmtMakespan search, implements search.Strategy interface
type DmtMakespan struct {
	*Engine
	heuristic   heuristic.StateEvaluator
	gWeight     float64
	hWeight     float64
	costtype    task.CostType
	c           float64 // exploration coefficient
	root        *DmtMakespanNode
	triallength int

	selectNode     func(*DmtMakespan, []*DmtMakespanNode) *DmtMakespanNode // selection strategy
	initializeNode func(*DmtMakespan, state.State, *DmtMakespanNode)       // initialization strategy
	backupNode     func(*DmtMakespan, *DmtMakespanNode)                    // backup strategy
	tiebreaker     func([]*DmtMakespanNode) *DmtMakespanNode               // tiebreaking strategy
	makespan       func([]*task.Action, []int) int                         // compute makespan
	processMessage func(*DmtMakespan, *StateActionMessage)                 // process state action messages

	transpositionTable map[state.StateID]*DmtMakespanNode
	backupQueue        *BackupQueue

	goalFound bool

	stateActionMessages chan interface{}

	neglect bool
}

func newDmtMakespan(e *Engine, opts *opt.OptionSet) *DmtMakespan {
	h, err := heuristic.Get(opts.GetString("heuristic"), e.agentID)
	if err != nil {
		log.Fatalf("ERR: uninitialized heuristic\n")
	}

	//var c float64
	//c, err = strconv.ParseFloat(opts.GetString("expcoeff"), 64)
	c := opts.GetFloat64("expcoeff")
	dmt := &DmtMakespan{
		Engine:             e,
		costtype:           task.CostType(opts.GetInt32("costtype")),
		c:                  c,
		heuristic:          h,
		gWeight:            opts.GetFloat64("g-weight"),
		hWeight:            opts.GetFloat64("h-weight"),
		initializeNode:     defaultInitializeDmtMakespan,
		triallength:        int(opts.GetInt32("triallength")),
		tiebreaker:         randomTiebreakerMakespan,
		transpositionTable: make(map[state.StateID]*DmtMakespanNode, 0),
		backupQueue:        NewBackupQueue(),

		stateActionMessages: make(chan interface{}, 1000000),
	}
	if opts.GetBool("neglect-private") {
		dmt.makespan = makespanApproxPublic(dmt.costtype)
	} else {
		dmt.makespan = makespanApprox
	}
	dmt.server.RegisterMessageChan(&StateActionMessage{}, dmt.stateActionMessages)
	return dmt
}

// NewDmtMakespanBfs returns a greedy dmt strategy
func NewDmtMakespanBfs(e *Engine, opts *opt.OptionSet) Strategy {
	dmt := newDmtMakespan(e, opts)
	dmt.selectNode = greedySelectDmtMakespan
	dmt.backupNode = greedyBackupDmtMakespan
	dmt.initializeNode = greedyInitializeDmtMakespan
	dmt.processMessage = greedyProcessStateActionMessage
	return dmt
}

// NewWeightedDmtMakespan creates a weighted, greedy dmt strategy
func NewWeightedDmtMakespan(e *Engine, opts *opt.OptionSet) Strategy {
	dmt := newDmtMakespan(e, opts)
	dmt.selectNode = weightedMakespanSelect
	dmt.backupNode = greedyBackupDmtMakespan
	dmt.processMessage = makespanProcessStateActionMessage
	return dmt
}

//// NewDmtMakespanGus returns a balanced (uct-like) dmt strategy
//func NewDmtMakespanGus(e *Engine, opts *opt.OptionSet) Strategy {
//	dmt := newDmtMakespan(e, opts)
//	dmt.selectNode = ucb1SelectDmtMakespan
//	dmt.backupNode = greedyBackupDmtMakespan
//	dmt.processMessage = makespanProcessStateActionMessage
//	return dmt
//}

// Tokens returns the tokenarray registered with node n.
func (d *DmtMakespan) Tokens(n Node) state.State {
	return d.stateRegistry.Tokens(d.stateRegistry.Lookup(n.ID()))
}

// HasNode returns the node stored with NodeID id in transmittedNodes and a
// truth value representing whether the id is in the table.
func (d *DmtMakespan) HasNode(sid state.StateID) (Node, bool) {
	n, ok := d.transpositionTable[sid]
	return n, ok
}

// Initialize initializes dmt search
// TODO: testing
func (d *DmtMakespan) Initialize() {
	s0 := append(d.Init, make([]int, len(d.conns))...)
	n := &DmtMakespanNode{
		stateID: d.stateRegistry.Register(s0),
		h:       d.heuristic.Evaluate(s0),
		visits:  1,
	}
	d.transpositionTable[n.stateID] = n
	d.root = n
	d.logger.Println("initialized DMTMakespan search!")
}

// nonclosed returns the slice of non-closed nodes of slice res
func nonclosedMakespan(ns []*DmtMakespanNode) (res []*DmtMakespanNode) {
	for _, n := range ns {
		if !n.closed {
			res = append(res, n)
		}
	}
	return res
}

// allclosed returns true iff all nodes from ns are closed
func allclosedMakespan(ns []*DmtMakespanNode) bool {
	for _, n := range ns {
		if !n.closed {
			return false
		}
	}
	return true
}

// Step executes a single search step (is called repeatedly by search.Engine)
func (d *DmtMakespan) Step() Status {
	d.processMessages() // non-blocking
	d.backup(false)

	if d.root.closed {
		// Wait for messages of other agents that may reopen the root.
		// TODO: Return unsolvable if all agents closed their root.
		return idle
	}
	//d.backupQueue.Clear()
	n := d.root
	for len(n.children) != 0 {
		n = d.selectNode(d, n.children)
	}
	exp := 1
	for ; ; exp++ {
		s := d.stateRegistry.Lookup(n.stateID)
		d.initializeNode(d, s, n)
		d.backupQueue.Put(n)
		if n.Action() != nil && n.Action().Owner == d.agentID && !n.Action().Private {
			d.sendMessages(s, n)
		}
		if allclosedMakespan(n.children) || exp >= d.triallength {
			break
		}
		n = d.selectNode(d, n.children)
	}
	d.backup(false)
	d.info.incExpansions(exp)
	return inProgress
}

// backup implements DMT's backup phase; backupNode is called on all nodes in
// the backupQueue, their parent is added to the queue, until the queue is empty
func (d *DmtMakespan) backup(resetVisits bool) {
	var n *DmtMakespanNode
	for d.backupQueue.Len() > 0 {
		n = d.backupQueue.Take().(*DmtMakespanNode)
		d.backupNode(d, n)
		if resetVisits {
			n.visits = 1
		}
		d.backupQueue.Put(n.Parent())
	}
}

// tiebreaker returns a random node from a node-slice
func randomTiebreakerMakespan(nodes []*DmtMakespanNode) *DmtMakespanNode {
	return nodes[rand.Intn(len(nodes))]
}

// computes sequential cost to the root
func (n *DmtMakespanNode) distance(ct task.CostType) int {
	distance := 0
	curn := n
	for !curn.isRoot() {
		if ct == task.UnitCost {
			for _ = range curn.costs {
				distance += 1
			}
		} else {
			for _, c := range curn.costs {
				distance += c
			}
		}
		curn = curn.parent
	}
	return distance
}

// getPublicCosts of the given action sequence containing private and public
// actions of this agent and only public actions from messages of other agents.
// Assume that private actions of other agents cost nothing and the makespan is
// only affected by the public actions of other agents.
func getPublicCosts(actions []*task.Action, costType task.CostType) []int {
	publicCosts := make([]int, len(actions))
	for i := range actions {
		publicCosts[i] = actions[i].AdjustedCost(costType)
	}
	return publicCosts
}

// makespanPublic computes the makespan value ignoring private actions costs
func makespanApproxPublic(costtype task.CostType) func([]*task.Action, []int) int {
	return func(actions []*task.Action, costs []int) int {
		return makespanApprox(actions, getPublicCosts(actions, costtype))
	}
}

// makespanApprox returns a minimum plan deordering approximating the makespan
func makespanApprox(actions []*task.Action, costs []int) int {
	// empty sequences have zero cost
	if len(actions) == 0 {
		return 0
	}
	makespans := make([]int, len(actions))
	for i := range makespans {
		makespans[i] = costs[i]
	}
	for i, ai := range actions {
		for j := i + 1; j < len(actions); j++ {
			aj := actions[j]
			if nonconcurrent(ai, aj) && makespans[j] < makespans[i]+costs[j] {
				makespans[j] = makespans[i] + costs[j]
			}
		}
	}
	return ints.Max(makespans...)
}

// nonconcurrent returns true iff two actions cannot be executed in parallel
func nonconcurrent(a, b *task.Action) bool {
	return a.Owner == b.Owner ||
		prodcons(a, b) ||
		prodthreat(a, b) ||
		consthreat(a, b)
}

func prodcons(a, b *task.Action) bool {
	prod := a.Effects
	cons := b.Preconditions
	for i, j := 0, 0; i < len(prod) && j < len(cons); {
		if prod[i].Variable < cons[j].Variable {
			i++
		} else if prod[i].Variable > cons[j].Variable {
			j++
		} else {
			if prod[i].Value == cons[j].Value {
				return true
			}
			i++
			j++
		}
	}
	return false
}

func prodthreat(a, b *task.Action) bool {
	prod := a.Effects
	threat := b.Effects
	for i, j := 0, 0; i < len(prod) && j < len(threat); {
		if prod[i].Variable < threat[j].Variable {
			i++
		} else if prod[i].Variable > threat[j].Variable {
			j++
		} else {
			if prod[i].Value != threat[j].Value {
				return true
			}
			i++
			j++
		}
	}
	return false
}

func consthreat(a, b *task.Action) bool {
	cons := a.Preconditions
	threat := b.Effects
	for i, j := 0, 0; i < len(cons) && j < len(threat); {
		if cons[i].Variable < threat[j].Variable {
			i++
		} else if cons[i].Variable > threat[j].Variable {
			j++
		} else {
			if cons[i].Value != threat[j].Value {
				return true
			}
			i++
			j++
		}
	}
	return false
}

// greedy successor selection
func greedySelectDmtMakespan(d *DmtMakespan, nodes []*DmtMakespanNode) *DmtMakespanNode {
	bestNodes := make([]*DmtMakespanNode, 0, len(nodes))
	bestH := heuristic.MaxCost
	for _, n := range nodes {
		if n.closed {
			continue
		}
		if n.h < bestH {
			bestNodes = append(bestNodes[:0], n)
			bestH = n.h
		} else if n.h == bestH {
			bestNodes = append(bestNodes, n)
		}
	}
	return d.tiebreaker(bestNodes)
}

// weightedMakespanSelect selects successors based on their (weighted) makespan
// and their (weighted) distance to the goal.
func weightedMakespanSelect(d *DmtMakespan, nodes []*DmtMakespanNode) *DmtMakespanNode {
	bestNodes := make([]*DmtMakespanNode, 0, len(nodes))
	fBest := heuristic.MaxCost
	var f int
	for _, n := range nodes {
		if n.closed {
			continue
		}

		f = int(d.gWeight*float64(n.m) + d.hWeight*float64(n.h))
		if f < fBest {
			bestNodes = append(bestNodes[:0], n)
			fBest = f
		} else if f == fBest {
			bestNodes = append(bestNodes, n)
			//d1 := n.distance(d.costtype)
			//d2 := bestNodes[0].distance(d.costtype)
			//if d1 < d2 {
			//	bestNodes = append(bestNodes[:0], n)
			//} else if d1 == d2 {
			//	bestNodes = append(bestNodes, n)
			//}
		}
	}
	return d.tiebreaker(bestNodes)
}

// propagate the minimum h-value among all children
func greedyBackupDmtMakespan(d *DmtMakespan, n *DmtMakespanNode) {
	n.closed = true
	n.h = heuristic.MaxCost
	n.m = ints.MaxValue
	for _, child := range n.children {
		if child.closed {
			continue
		}
		n.closed = false
		// set node h-value to smallest children h-value
		if child.h < n.h {
			n.h = child.h
		}
		// set node makespan to smallest children makespan
		if child.m < n.m {
			n.m = child.m
		}
	}
	n.visits++
}

// initializes a node by creating all successor nodes and computing its h value
func defaultInitializeDmtMakespan(d *DmtMakespan, s state.State, n *DmtMakespanNode) {
	var (
		succ     state.State
		succID   state.StateID
		h        int
		makespan int
	)
	applicableActions := d.actionPruner.Prune(&s, d.succ.ApplicableActions(s))
	for _, o := range applicableActions {
		succ = Successor(s, o)
		succID = d.stateRegistry.Register(succ)
		h = d.heuristic.Evaluate(succ)

		// 0) check for deadend
		if h == heuristic.MaxCost {
			continue // TODO: incr deadends
		}

		// compute makespan
		actions, costs := n.pathToRoot()
		actions = append(actions, o)
		costs = append(costs, o.AdjustedCost(d.costtype))
		makespan = d.makespan(actions, costs)

		// prune node if it is worse than the best found plan (regarding makespan)
		if makespan >= d.planRegistry.bestMakespan {
			continue
		}

		// 1) check for transposition node
		tp, ok := d.transpositionTable[succID]
		if ok {
			if !tp.closed {
				if makespan < tp.m {
					d.backupQueue.Put(tp.parent)
					tp.parent.removeChild(tp)
					tp.m = makespan
					tp.parent = n
					n.addChild(tp)
					tp.actions = []*task.Action{o}
					tp.costs = []int{o.AdjustedCost(d.costtype)}
				}
			}
			continue
		}

		// 3) no transposition -> create new node
		newNode := &DmtMakespanNode{
			stateID: succID,
			parent:  n,
			actions: []*task.Action{o},
			costs:   []int{o.AdjustedCost(d.costtype)},
			h:       h,
			m:       makespan,
			visits:  1,
		}

		n.addChild(newNode)
		d.transpositionTable[succID] = newNode

		// 4) found a goal -> register & extract
		if Satisfies(succ, d.Goal) {
			newNode.closed = true
			d.goalNode = newNode
			d.planRegistry.Start(d.goalNode, d.costtype)
			d.backup(false)
			//DotGraph(d.root, fmt.Sprintf("%v.dot", d.agentID))
			//DotPlan(newNode, fmt.Sprintf("%v-plan.dot", d.agentID))
			//log.Fatalf("ende")
		}

	}
	d.info.setHeuristicValue(h) // update progress
}

// initializes a node by creating all successor nodes and computing its h value
func greedyInitializeDmtMakespan(d *DmtMakespan, s state.State, n *DmtMakespanNode) {
	var (
		succ     state.State
		succID   state.StateID
		h        int
		makespan int
	)
	applicableActions := d.actionPruner.Prune(&s, d.succ.ApplicableActions(s))
	for _, o := range applicableActions {
		succ = Successor(s, o)
		succID = d.stateRegistry.Register(succ)
		h = d.heuristic.Evaluate(succ)

		// 0) check for deadend
		if h == heuristic.MaxCost {
			continue // TODO: incr deadends
		}

		// compute makespan
		actions, costs := n.pathToRoot()
		actions = append(actions, o)
		costs = append(costs, o.AdjustedCost(d.costtype))
		makespan = d.makespan(actions, costs)
		// compute g-value
		gValue := n.distance(d.costtype) + o.AdjustedCost(d.costtype)

		// prune node if it is worse than the best found plan (regarding gValue)
		if gValue >= d.planRegistry.bestCost {
			continue
		}

		// 1) check for transposition node
		tp, ok := d.transpositionTable[succID]
		if ok {
			if !tp.closed {
				if gValue < tp.distance(d.costtype) {
					d.backupQueue.Put(tp.parent)
					tp.parent.removeChild(tp)
					tp.m = makespan
					tp.parent = n
					n.addChild(tp)
					tp.actions = []*task.Action{o}
					tp.costs = []int{o.AdjustedCost(d.costtype)}
				}
			}
			continue
		}

		// 3) no transposition -> create new node
		newNode := &DmtMakespanNode{
			stateID: succID,
			parent:  n,
			actions: []*task.Action{o},
			costs:   []int{o.AdjustedCost(d.costtype)},
			h:       h,
			m:       makespan,
			visits:  1,
		}

		n.addChild(newNode)
		d.transpositionTable[succID] = newNode

		// 4) found a goal -> register & extract
		if Satisfies(succ, d.Goal) {
			newNode.closed = true
			d.goalNode = newNode
			d.planRegistry.Start(d.goalNode, d.costtype)
			d.backup(false)
			//DotGraph(d.root, fmt.Sprintf("%v.dot", d.agentID))
			//DotPlan(newNode, fmt.Sprintf("%v-plan.dot", d.agentID))
			//log.Fatalf("ende")
		}

	}
	d.info.setHeuristicValue(h) // update progress
}

// selection strategy similar to ucb1 for cp
func ucb1SelectDmtMakespan(d *DmtMakespan, nodes []*DmtMakespanNode) *DmtMakespanNode {
	// TODO: this only works when greedy backup is used -> FIX!!!
	minH := float64(nodes[0].parent.h)
	parentVisits := float64(nodes[0].parent.visits)
	visitsLog := math.Log(parentVisits)
	maxH := 0.0 // compute best H score
	for _, n := range nodes {
		if !n.closed && float64(n.h) > maxH {
			maxH = float64(n.h)
		}
	}
	if maxH-minH == 0 {
		return d.tiebreaker(nonclosedMakespan(nodes))
	}
	bestNodes := make([]*DmtMakespanNode, 0, len(nodes))
	best, ucb1Score := math.MaxFloat64, float64(0)
	for _, n := range nodes {
		if n.closed {
			continue
		}
		if n.h == 0 {
			return n
		}
		ucb1Score = (float64(n.h)-minH)/(maxH-minH) -
			d.c*math.Sqrt(visitsLog/float64(n.visits))
		if floats.IsSmaller(ucb1Score, best) {
			bestNodes = append(bestNodes[:0], n)
			best = ucb1Score
		} else if floats.AlmostEquals(ucb1Score, best) {
			bestNodes = append(bestNodes, n)
		}
	}

	if len(bestNodes) == 0 {
		cs := make([]DmtMakespanNode, len(nodes))
		for i, n := range nodes {
			cs[i] = *n
		}
		log.Fatalf("ERROR: No best nodes: %v, best:%g, maxH:%g "+
			"parentH:%g\n", cs, best, maxH, minH)
	}
	return d.tiebreaker(bestNodes)
}

// perAgentCost computes for each agent the cost from n to the
// ancestor node that contains a different token for the agent.
func (d *DmtMakespan) sendMessages(s state.State, n *DmtMakespanNode) {
	sharedState := d.stateRegistry.SharedState(s, n.stateID)
	agents := d.succ.ApplicableAgents(d.stateRegistry.Public(s))
	costs := make([]int, 0)
	actions := make([]int, 0)
	var action *task.Action

	curNode := n
	for !curNode.isRoot() && len(agents) != 0 {
		// private action:
		// add cost of private action to the cost of the first public action
		// (the one that was prepended last) of
		// the same agent, then continue with the next parent node
		if curNode.actions[len(curNode.actions)-1].Private {
			costs[0] += curNode.costs[0]
			curNode = curNode.parent
			continue
		}
		// public action(s):
		for j := len(curNode.actions) - 1; j >= 0; j-- {
			action = curNode.actions[j]
			for i, agent := range agents {
				if action.Owner == agent {
					agents = append(agents[:i], agents[i+1:]...) // remove agent
					d.sendMessage(agent, sharedState, actions, costs, n.h)
					break
				}
			}
			actions = prepend(action.ID, actions)
			costs = prepend(curNode.costs[j], costs)
		}
		// continue with parent
		curNode = curNode.parent
	}
	for _, agent := range agents {
		d.sendMessage(agent, sharedState, actions, costs, n.h)
	}
}

// prepend prepends an int to a slice of ints
// TODO: -> implement a more efficient way
func prepend(v int, vs []int) []int {
	return append([]int{v}, vs...)
}

func (d *DmtMakespan) sendMessage(toAgentID int, sharedState state.SharedState,
	actions []int, costs []int, hValue int) {

	d.dispatcher.Send(
		toAgentID,
		&StateActionMessage{
			SenderID: d.agentID,
			State:    sharedState,
			Actions:  actions,
			Costs:    costs,
			H:        hValue,
		},
	)
	d.info.incMessagesOut(1)
}

// TODO: move somewhere appropriate, find suitable name...
func actionIDs(actions []*task.Action) []int {
	ids := make([]int, len(actions))
	for i, a := range actions {
		ids[i] = a.ID
	}
	return ids
}

// TODO: move somewhere appropriate, find suitable name...
func (dmt *DmtMakespan) idsToActions(ids []int) []*task.Action {
	actions := make([]*task.Action, len(ids))
	for i, id := range ids {
		actions[i] = dmt.GetAction(id)
	}
	return actions
}

// processMessages handles incoming messages depending on type/content
func (d *DmtMakespan) processMessages() {
	var m interface{}
	for {
		// XMessages have higher priority, so check them out first
		select {
		case m = <-d.xMessages:
			d.planRegistry.processMessage(m.(*XMessage))
			continue
		default: // do not block, to process other message types
		}

		select {
		case m = <-d.stateActionMessages:
			d.processMessage(d, m.(*StateActionMessage))
		default:
			return
		}
	}
}

// processStateActionMessage integrates a received state into the tree
func makespanProcessStateActionMessage(d *DmtMakespan, m *StateActionMessage) {
	d.info.incMessagesIn(1)

	parent := d.transpositionTable[d.stateRegistry.Token(m.State)]
	newState := d.stateRegistry.GlobalState(m.State)
	newSID := d.stateRegistry.Register(newState)

	// compute makespan
	actions, costs := parent.pathToRoot()
	actions = append(actions, d.idsToActions(m.Actions)...)
	costs = append(costs, m.Costs...)
	makespan := d.makespan(actions, costs)

	// prune if state has worse makespan than best found plans makespan
	if makespan >= d.planRegistry.bestMakespan {
		return
	}

	// handle transpositions
	tp, ok := d.transpositionTable[newSID]
	if ok {
		if !tp.closed {
			if makespan < tp.m {
				tp.parent.removeChild(tp)
				d.backupQueue.Put(tp.parent)
				tp.costs = m.Costs
				tp.actions = d.idsToActions(m.Actions)
				tp.m = makespan
				parent.addChild(tp)
				tp.parent = parent
				d.backupQueue.Put(parent)
			}
		}
		return
	}

	// add new node
	//h := int(ints.Min(d.heuristic.Evaluate(newState), m.H))
	h := d.heuristic.Evaluate(newState)
	newNode := &DmtMakespanNode{
		stateID: newSID,
		parent:  parent,
		actions: d.idsToActions(m.Actions), // not a nice solution
		costs:   m.Costs,
		h:       h,
		m:       makespan,
		visits:  1,
	}
	if Satisfies(newState, d.Goal) {
		newNode.closed = true
	}
	parent.addChild(newNode)
	d.backupQueue.Put(parent)
	d.transpositionTable[newSID] = newNode
}

// greedyProcessStateActionMessage integrates a received state into the tree
func greedyProcessStateActionMessage(d *DmtMakespan, m *StateActionMessage) {
	d.info.incMessagesIn(1)

	parent := d.transpositionTable[d.stateRegistry.Token(m.State)]
	newState := d.stateRegistry.GlobalState(m.State)
	newSID := d.stateRegistry.Register(newState)

	// compute g-value
	gValue := parent.distance(d.costtype)
	for _, cost := range m.Costs {
		gValue += cost
	}

	// prune if state has worse cost than best found plans cost
	if gValue >= d.planRegistry.bestCost {
		return
	}

	// handle transpositions
	tp, ok := d.transpositionTable[newSID]
	if ok {
		if !tp.closed {
			if gValue < tp.distance(d.costtype) {
				tp.parent.removeChild(tp)
				d.backupQueue.Put(tp.parent)
				tp.costs = m.Costs
				tp.actions = d.idsToActions(m.Actions)
				parent.addChild(tp)
				tp.parent = parent
				tp.m = d.makespan(tp.pathToRoot())
				d.backupQueue.Put(parent)
			}
		}
		return
	}

	// add new node
	//h := int(ints.Min(d.heuristic.Evaluate(newState), m.H))
	h := d.heuristic.Evaluate(newState)
	newNode := &DmtMakespanNode{
		stateID: newSID,
		parent:  parent,
		actions: d.idsToActions(m.Actions), // not a nice solution
		costs:   m.Costs,
		h:       h,
		visits:  1,
	}
	newNode.m = d.makespan(newNode.pathToRoot())
	if Satisfies(newState, d.Goal) {
		newNode.closed = true
	}
	parent.addChild(newNode)
	d.backupQueue.Put(parent)
	d.transpositionTable[newSID] = newNode
}

// init registers dmt strategies
func init() {
	// we need to tell gob about the concrete message type used by the search
	gob.Register(&StateActionMessage{})
	// options to be parsed for all dmt strategies
	dmtMakespanOptions := func() *opt.OptionSet {
		opts := opt.NewOptionSet()
		opts.Add(opt.NewOption(opt.String, "heuristic", 'h', "ff", "heuristic fct"))
		opts.Add(opt.NewOption(opt.Float64, "expcoeff", 'c', "1.4142", "explorati"))
		opts.Add(opt.NewOption(opt.Int32, "costtype", 0, 2,
			"NormalCost=0|UnitCost=1|NormalMinOne=2"))
		opts.Add(opt.NewOption(opt.Int32, "triallength", 'l', 1, "trial length"))
		opts.Add(opt.NewOption(opt.Float64, "g-weight", 'w', "1.0", "weight of g-value"))
		opts.Add(opt.NewOption(opt.Float64, "h-weight", 'W', "1.0", "weight of h-value"))
		opts.Add(opt.NewOption(opt.Bool, "neglect-private", 'n', false, "neglect private actions of other agents"))
		return opts
	}
	strategyRegistry = append(strategyRegistry,
		&StrategyInfo{
			Name:        "dmtmakespan-bfs",
			Description: "DMT Makespan Greedy BFS",
			NewStrategy: NewDmtMakespanBfs,
			Options:     dmtMakespanOptions(),
		},
		&StrategyInfo{
			Name:        "dmtmakespan-weighted",
			Description: "DMT Makespan Greedy Weighted",
			NewStrategy: NewWeightedDmtMakespan,
			Options:     dmtMakespanOptions(),
		},
		// TODO: just a shortcut b/c I am tired of looking up the name
		// `dmtmakespan-weighted`
		&StrategyInfo{
			Name:        "dmtm",
			Description: "DMT Makespan Greedy Weighted",
			NewStrategy: NewWeightedDmtMakespan,
			Options:     dmtMakespanOptions(),
		},
		&StrategyInfo{
			Name:        "dmtg",
			Description: "DMT Greedy BFS",
			NewStrategy: NewDmtMakespanBfs,
			Options:     dmtMakespanOptions(),
		},
	)
}
