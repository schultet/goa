package search

import (
	"fmt"
	"log"
	"math"
	"math/rand"

	"github.com/schultet/goa/src/heuristic"
	"github.com/schultet/goa/src/opt"
	"github.com/schultet/goa/src/state"
	"github.com/schultet/goa/src/task"
	"github.com/schultet/goa/src/util/floats"
)

// DmtMessage contains the search state and other information
type DmtMessage struct {
	SenderID int
	State    []int // SharedState
	Action   int
	G        int
	H        int
}

// DmtNode implements search.Node
type DmtNode struct {
	stateID   state.StateID
	parent    *DmtNode
	children  []*DmtNode
	action    *task.Action
	h         int
	visits    int
	closed    bool
	cost      int
	disregard byte
}

// NewDmtNode creates a new DmtNode
func NewDmtNode(sid state.StateID, par *DmtNode, op *task.Action, cost, h, v int) *DmtNode {
	return &DmtNode{stateID: sid, parent: par, action: op, cost: cost, h: h, visits: v}
}

// Action returns the nodes creating action
func (n *DmtNode) Action() *task.Action { return n.action }

// Parent returns the nodes parent node
func (n *DmtNode) Parent() Node {
	// since we return an interface type, we must explicitly return nil here,
	// even though n.parent == nil it is not the same to return n.parent.
	if n.parent == nil {
		return nil
	}
	return n.parent
}

// ID returns the unique NodeID of the node
func (n *DmtNode) ID() state.StateID { return n.stateID }

// isRoot returns true iff the node is the root (i.e. has no parent)
func (n *DmtNode) isRoot() bool { return n.parent == nil }

// String returns a string representation of n
func (n *DmtNode) String() string {
	//return fmt.Sprintf("(id: %d, sid:%d, tokens:%d, op:%s par:%s)",
	//	n.id, n.sid, n.tid, n.op, n.parent)
	return fmt.Sprintf("(sid:%d, op:%s par:%s)", n.stateID, n.action, n.parent)
}

// removeChild removes child c from n's children
func (n *DmtNode) removeChild(c *DmtNode) {
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
func (n *DmtNode) addChild(c *DmtNode) {
	n.children = append(n.children, c)
}

// forbearOf checks whether n is a forbear (=ancestor) of other
func (n *DmtNode) forbearOf(other *DmtNode) bool {
	for !other.isRoot() {
		if other.parent == n {
			return true
		}
		other = other.parent
	}
	return false
}

// Dmt search, implements search.Strategy interface
type Dmt struct {
	*Engine
	heuristic   heuristic.StateEvaluator
	costtype    task.CostType
	c           float64 // exploration coefficient
	root        *DmtNode
	triallength int

	selectNode     func(*Dmt, []*DmtNode) *DmtNode   // selection strategy
	initializeNode func(*Dmt, state.State, *DmtNode) // initialization strategy
	backupNode     func(*Dmt, *DmtNode)              // backup strategy
	tiebreaker     func([]*DmtNode) *DmtNode

	transpositionTable map[state.StateID]*DmtNode
	backupQueue        *BackupQueue

	goalFound bool
	lazy      bool // evaluate heuristics lazily
	disregard byte

	stateMessages chan interface{}
}

func newDmt(e *Engine, opts *opt.OptionSet) *Dmt {
	h, err := heuristic.Get(opts.GetString("heuristic"), e.agentID)
	if err != nil {
		log.Fatalf("ERR: uninitialized heuristic\n")
	}

	//var c float64
	//c, err = strconv.ParseFloat(opts.GetString("expcoeff"), 64)
	c := opts.GetFloat64("expcoeff")
	dmt := &Dmt{
		Engine:             e,
		costtype:           task.CostType(opts.GetInt32("costtype")),
		c:                  c,
		heuristic:          h,
		initializeNode:     (*Dmt).defaultInitialize,
		triallength:        int(opts.GetInt32("triallength")),
		tiebreaker:         randomTiebreaker,
		transpositionTable: make(map[state.StateID]*DmtNode),
		backupQueue:        NewBackupQueue(),
		lazy:               opts.GetBool("lazy"),
		disregard:          byte(opts.GetInt32("disregard")),

		stateMessages: make(chan interface{}, 1000000),
	}
	dmt.server.RegisterMessageChan(&DmtMessage{}, dmt.stateMessages)
	return dmt
}

// NewDmtBfs returns a greedy dmt strategy
func NewDmtBfs(e *Engine, opts *opt.OptionSet) Strategy {
	dmt := newDmt(e, opts)
	dmt.selectNode = greedySelectDmt
	dmt.backupNode = greedyBackupDmt
	return dmt
}

// NewDmtGus returns a balanced (uct-like) dmt strategy
func NewDmtGus(e *Engine, opts *opt.OptionSet) Strategy {
	dmt := newDmt(e, opts)
	dmt.selectNode = ucb1SelectDmt
	dmt.backupNode = greedyBackupDmt
	return dmt
}

// Tokens returns the tokenarray registered with node n.
func (d *Dmt) Tokens(n Node) state.State {
	return d.stateRegistry.Tokens(d.stateRegistry.Lookup(n.ID()))
}

// HasNode returns the node stored with NodeID id in transmittedNodes and a
// truth value representing whether the id is in the table.
func (d *Dmt) HasNode(sid state.StateID) (Node, bool) {
	n, ok := d.transpositionTable[sid]
	return n, ok
}

// Initialize initializes dmt search
// TODO: testing
func (d *Dmt) Initialize() {
	s0 := append(d.Init, make([]int, len(d.conns))...)
	n := &DmtNode{
		stateID: d.stateRegistry.Register(s0),
		h:       d.heuristic.Evaluate(s0),
		visits:  1,
	}
	d.transpositionTable[n.stateID] = n
	d.root = n
	d.logger.Println("initialized DMT search!")
}

// nonclosed returns the slice of non-closed nodes of slice res
func nonclosed(ns []*DmtNode) (res []*DmtNode) {
	for _, n := range ns {
		if !n.closed {
			res = append(res, n)
		}
	}
	return res
}

// allclosed returns true iff all nodes from ns are closed
func allclosed(ns []*DmtNode) bool {
	for _, n := range ns {
		if !n.closed {
			return false
		}
	}
	return true
}

// Step executes a single search step (is called repeatedly by search.Engine)
func (d *Dmt) Step() Status {
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
		//if n.Action() != nil && n.Action().Owner == d.agentID && !n.Action().Private {
		//	d.sendMessages(s, n)
		//}
		if allclosed(n.children) || exp >= d.triallength {
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
func (d *Dmt) backup(resetVisits bool) {
	var n *DmtNode
	for d.backupQueue.Len() > 0 {
		n = d.backupQueue.Take().(*DmtNode)
		d.backupNode(d, n)
		d.backupQueue.Put(n.Parent())

		// send messages and/or reduce disregard
		if n.disregard == 1 {
			if n.Action() != nil && !n.Action().Private {
				d.sendMessages(d.stateRegistry.Lookup(n.stateID), n)
			}
		}
		if n.disregard > 0 {
			n.disregard -= 1
		}
		if resetVisits {
			n.visits = 1
		}
	}
}

// tiebreaker returns a random node from a node-slice
func randomTiebreaker(nodes []*DmtNode) *DmtNode {
	return nodes[rand.Intn(len(nodes))]
}

// computes cost to the root
func (n *DmtNode) distance(ct task.CostType) int {
	distance := 0
	curn := n
	for !curn.isRoot() {
		if ct == task.UnitCost {
			distance++
		} else {
			distance += curn.cost
		}
		curn = curn.parent
	}
	return distance
}

// greedy successor selection
func greedySelectDmt(d *Dmt, nodes []*DmtNode) *DmtNode {
	bestNodes := make([]*DmtNode, 0, len(nodes))
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

// propagate the minimum h-value among all children
func greedyBackupDmt(d *Dmt, n *DmtNode) {
	n.closed = true
	n.h = heuristic.MaxCost
	for _, child := range n.children {
		if child.closed {
			continue
		}
		n.closed = false
		// set node h-value to smallest children h-value
		if child.h < n.h {
			n.h = child.h
		}
	}
	if n.closed && n.disregard > 0 {
		n.disregard = 1 // transmit this state in next backup
	}
	n.visits++
}

// initializes a node by creating all successor nodes and computing its h value
func (d *Dmt) defaultInitialize(s state.State, n *DmtNode) {
	var (
		succ   state.State
		succID state.StateID
		h      int
		gValue int
	)
	if d.lazy {
		h = d.heuristic.Evaluate(s)
		n.h = h
	}
	applicableActions := d.actionPruner.Prune(&s, d.succ.ApplicableActions(s))
	for _, o := range applicableActions {
		succ = Successor(s, o)
		succID = d.stateRegistry.Register(succ)
		if !d.lazy {
			h = d.heuristic.Evaluate(succ)
		}

		// 0) check for deadend
		if h == heuristic.MaxCost {
			continue // TODO: incr deadends
		}

		// compute g-value
		gValue = n.distance(d.costtype) + o.AdjustedCost(d.costtype)
		// prune node if it is worse than or equal to the best found plan
		// (regarding cost)
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
					tp.parent = n
					n.addChild(tp)
					tp.action = o
					tp.cost = o.AdjustedCost(d.costtype)
				}
			}
			// TODO: what if this is a goal state?
			continue
		}

		// 3) no transposition -> create new node
		newNode := &DmtNode{
			stateID:   succID,
			parent:    n,
			action:    o,
			cost:      o.AdjustedCost(d.costtype),
			h:         h,
			visits:    1,
			disregard: d.disregard,
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
func ucb1SelectDmt(d *Dmt, nodes []*DmtNode) *DmtNode {
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
		return d.tiebreaker(nonclosed(nodes))
	}
	bestNodes := make([]*DmtNode, 0, len(nodes))
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
		cs := make([]DmtNode, len(nodes))
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
func (d *Dmt) sendMessages(s state.State, n *DmtNode) {
	sharedState := d.stateRegistry.SharedState(s, n.stateID)
	agents := d.succ.ApplicableAgents(d.stateRegistry.Public(s))
	cost := 0

	curNode := n
	for !curNode.isRoot() && len(agents) != 0 {
		// private action:
		if curNode.action.Private {
			cost += curNode.cost
			curNode = curNode.parent
			continue
		}
		// public action:
		for i, agent := range agents {
			if curNode.action.Owner == agent {
				agents = append(agents[:i], agents[i+1:]...) // remove agent
				d.sendMessage(agent, sharedState, cost, n.h, n.action)
				break
			}
		}
		cost += curNode.cost
		// continue with parent
		curNode = curNode.parent
	}
	for _, agent := range agents {
		d.sendMessage(agent, sharedState, cost, n.h, n.action)
	}
}

func (d *Dmt) sendMessage(toAgentID int, sharedState state.SharedState,
	cost int, hValue int, action *task.Action) {

	d.dispatcher.Send(
		toAgentID,
		&DmtMessage{
			SenderID: d.agentID,
			State:    sharedState,
			G:        cost,
			H:        hValue,
			Action:   action.ID,
		},
	)
	d.info.incMessagesOut(1)
}

// processMessages handles incoming messages depending on type/content
func (d *Dmt) processMessages() {
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
		case m = <-d.stateMessages:
			processMessageDmt(d, m.(*DmtMessage))
		default:
			return
		}
	}
}

// processMessage integrates a received state into the tree
func processMessageDmt(d *Dmt, m *DmtMessage) {
	d.info.incMessagesIn(1)

	parent := d.transpositionTable[d.stateRegistry.Token(m.State)]
	newState := d.stateRegistry.GlobalState(m.State)
	newSID := d.stateRegistry.Register(newState)

	gValue := parent.distance(d.costtype) + m.G
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
				tp.cost = m.G
				tp.action = d.GetAction(m.Action)
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
	newNode := &DmtNode{
		stateID: newSID,
		parent:  parent,
		action:  d.GetAction(m.Action), // not a nice solution
		cost:    m.G,
		h:       h,
		visits:  1,
	}
	if Satisfies(newState, d.Goal) {
		newNode.closed = true
	}
	parent.addChild(newNode)
	d.backupQueue.Put(parent)
	d.transpositionTable[newSID] = newNode
}

// init registers dmt strategies
func init() {
	// options to be parsed for all dmt strategies
	dmtOptions := func() *opt.OptionSet {
		opts := opt.NewOptionSet()
		opts.Add(opt.NewOption(opt.String, "heuristic", 'h', "ff", "heuristic fct"))
		opts.Add(opt.NewOption(opt.Float64, "expcoeff", 'c', "1.4142", "explorati"))
		opts.Add(opt.NewOption(opt.Int32, "costtype", 0, 2,
			"NormalCost=0|UnitCost=1|NormalMinOne=2"))
		opts.Add(opt.NewOption(opt.Int32, "triallength", 'l', 1, "trial length"))
		opts.Add(opt.NewOption(opt.Bool, "lazy", 'z', false,
			"lazy heuristic evaluation?"))
		opts.Add(opt.NewOption(opt.Int32, "disregard", 'd', 1, "disregard of message sending"))
		return opts
	}
	strategyRegistry = append(strategyRegistry, &StrategyInfo{
		Name:        "dmt-bfs",
		Description: "distributed multi-agent thts bfs",
		NewStrategy: NewDmtBfs,
		Options:     dmtOptions(),
	}, &StrategyInfo{
		Name:        "dmt-gus",
		Description: "distributed multi-agent thts greedy uct star",
		NewStrategy: NewDmtGus,
		Options:     dmtOptions(),
	})
}
