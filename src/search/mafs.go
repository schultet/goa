package search

import (
	"encoding/gob"
	"fmt"
	"log"
	"math/rand"

	"github.com/schultet/goa/src/heuristic"
	"github.com/schultet/goa/src/opt"
	"github.com/schultet/goa/src/state"
	"github.com/schultet/goa/src/task"
)

// ClosedList is a hashmap containing all nodes expanded during search
type ClosedList map[state.StateID]*MafsNode

func (cl ClosedList) Add(n *MafsNode) {
	tp, ok := cl[n.stateID]
	if !ok || n.g < tp.g {
		cl[n.stateID] = n
	}
}

// MafsMessage contains a search state, and other information
type MafsMessage struct {
	SenderID int
	State    []int // SharedState
	G        int
	H        int
}

// MafsNode is the node type for mafs search implementing search.Node interface
type MafsNode struct {
	g       int
	h       int
	action  *task.Action
	parent  *MafsNode
	stateID state.StateID
	deadend bool
}

// H is the h-value accessor to satisfy ValueNode interface
func (n *MafsNode) H() int { return n.h }

// G is the g-value accessor to satisfy ValueNode interface
func (n *MafsNode) G() int { return n.g }

// Action returns the nodes creating action
func (n *MafsNode) Action() *task.Action { return n.action }

// Parent returns the nodes parent node
func (n *MafsNode) Parent() Node { return n.parent }

// ID returns the stateID the node is associated with uniquely
func (n *MafsNode) ID() state.StateID { return n.stateID }

func (n *MafsNode) distance() int {
	d := 0
	for n != nil {
		d++
		n = n.parent
	}
	return d
}

func (n *MafsNode) String() string {
	a := "nil"
	if n.action != nil {
		a = n.action.String()
	}
	return fmt.Sprintf("(g: %d, h: %d, a: %s)",
		n.g, n.h, a)
}

// NodeGetter instances can access an element at index i with Get(i)
type NodeGetter interface {
	Get(int) ValueNode
}

// Mafs (Multi agent forward search) implements search.Strategy interface
type Mafs struct {
	*Engine
	heuristic    heuristic.StateEvaluator
	costtype     task.CostType
	open         NodeQueue
	closed       ClosedList
	triallength  int
	reopenClosed bool
	lazy         bool

	selectNode func(*Mafs, []*MafsNode) *MafsNode

	stateMessages chan interface{}
}

// NewMafs creates a new mafs (multi-agent forward search) strategy
func NewMafs(e *Engine, opts *opt.OptionSet) Strategy {
	h, err := heuristic.Get(opts.GetString("heuristic"), e.agentID)
	if err != nil {
		log.Fatalf("ERR heuristic not initialized\n")
	}

	//e.ActionPruner = getActionPruner(opts, e.Task)
	evaluators := map[string]NodeQueue{
		"greedy":     NewPriorityQueue(),
		"astar":      NewAStarQueue(),
		"greedyfifo": NewPriorityQueueFIFO(),
	}
	opts.GetString("evaluator")
	mafs := &Mafs{
		Engine:        e,
		heuristic:     h,
		costtype:      task.CostType(opts.GetInt32("costtype")),
		open:          evaluators[opts.GetString("evaluator")],
		closed:        make(ClosedList, 0),
		triallength:   int(opts.GetInt32("triallength")),
		reopenClosed:  opts.GetBool("reopen_closed"),
		lazy:          opts.GetBool("lazy"),
		stateMessages: make(chan interface{}, 1000000),
	}
	switch opts.GetString("rollout-policy") {
	case "greedy":
		mafs.selectNode = (*Mafs).selectGreedy
	case "random":
		mafs.selectNode = (*Mafs).selectRandom
	}
	mafs.server.RegisterMessageChan(&MafsMessage{}, mafs.stateMessages)
	return mafs
}

// NewNode returns a new MafsNode
func (bfs *Mafs) NewNode(sid state.StateID, par *MafsNode, a *task.Action, g, h int) *MafsNode {
	return &MafsNode{
		stateID: sid,
		parent:  par,
		action:  a,
		g:       g,
		h:       h,
		deadend: false,
	}
}

// Tokens returns the tokenarray registered with node n.
func (bfs *Mafs) Tokens(n Node) state.State {
	return bfs.stateRegistry.Tokens(bfs.stateRegistry.Lookup(n.ID()))
}

// HasNode returns the node associated with sid in closed list
func (bfs *Mafs) HasNode(sid state.StateID) (Node, bool) {
	n, ok := bfs.closed[sid]
	if ok {
		return n, ok
	}
	// if the node wasn't expanded -> look in `open`
	for i := 0; i < bfs.open.Len(); i++ {
		node := bfs.open.(NodeGetter).Get(i).(*MafsNode)
		//for _, node := range bfs.open {
		if node.stateID == sid {
			return node, true
		}
	}
	return nil, false
}

// Initialize initializes Mafs search
func (bfs *Mafs) Initialize() {
	s0 := append(bfs.Init, make([]int, len(bfs.conns))...)
	initNode := &MafsNode{
		h:       bfs.heuristic.Evaluate(s0),
		stateID: bfs.stateRegistry.Register(s0),
	}
	//bfs.open = append(bfs.open, initNode)
	bfs.open.Put(initNode)
	bfs.logger.Println("initialized MAFS search!")
}

// Step removes and expands best node from open list and adds it to closed list.
// Returns idle when open list is empty, solved when a goal is found, and
// inProgress otherwise.
func (bfs *Mafs) Step() Status {
	bfs.processMessages() // process all incoming messages
	if bfs.open.Len() == 0 {
		return idle // TODO: if all idle return unsolvable
	}
	n := bfs.open.Take().(*MafsNode)

	// repeatedly expand nodes and add successors to openlist
	var s state.State
	exp := 1
	for ; ; exp++ {
		s = bfs.stateRegistry.Lookup(n.stateID)
		bfs.closed.Add(n)
		if n.Action() != nil && n.Action().Owner == bfs.agentID && !n.Action().Private {
			bfs.sendMessages(s, n)
		}
		n = bfs.selectNode(bfs, bfs.expand(s, n))

		if n == nil || exp >= bfs.triallength {
			break
		}
	}

	bfs.info.incExpansions(exp)
	return inProgress
}

// returns a slice of BfsNodes that are successors to n
func (bfs *Mafs) expand(s state.State, n *MafsNode) []*MafsNode {
	bfs.info.setHeuristicValue(n.h) // update smallest h-value

	applicableActions := bfs.actionPruner.Prune(&s, bfs.succ.ApplicableActions(s))

	if len(applicableActions) == 0 { // TODO: detect deadends
		n.deadend = true
	}

	successors := make([]*MafsNode, 0, len(applicableActions))
	var (
		succState state.State
		succSID   state.StateID
		succNode  *MafsNode
		h         int
	)
	if bfs.lazy {
		h = bfs.heuristic.Evaluate(s)
		n.h = h
	}
	for _, o := range applicableActions {
		succState = Successor(s, o)
		succSID = bfs.stateRegistry.Register(succState)
		g := n.g + o.AdjustedCost(bfs.costtype)
		if !bfs.lazy {
			h = bfs.heuristic.Evaluate(succState)
		}

		// (1) prune if worse than or equal to best found plan
		if g >= bfs.planRegistry.bestCost {
			continue
		}

		// (2) if its a goal state, extract plan
		if Satisfies(succState, bfs.Goal) {
			succNode = bfs.NewNode(succSID, n, o, g, bfs.heuristic.Evaluate(succState))
			bfs.goalNode = succNode
			bfs.planRegistry.Start(succNode, bfs.costtype)
			continue
		}

		// (3) check if node is closed
		tp, ok := bfs.closed[succSID]
		if ok {
			if g < tp.g && !tp.deadend {
				tp.parent = n
				tp.g = g
				if bfs.reopenClosed {
					bfs.open.Put(tp)
					successors = append(successors, tp)
				}
				continue
			}
			continue
		}

		// (4) its a new node -> add to open
		succNode = bfs.NewNode(succSID, n, o, g, h)
		bfs.open.Put(succNode)
		successors = append(successors, succNode)
	}
	return successors
}

// selectRandom returns random node from `nodes`
func (bfs *Mafs) selectRandom(nodes []*MafsNode) *MafsNode {
	return nodes[rand.Intn(len(nodes))]
}

// selectGreedy returns best (= lowest h-value) node from `nodes`
func (bfs *Mafs) selectGreedy(nodes []*MafsNode) *MafsNode {
	if len(nodes) == 0 {
		return nil
	}
	best := []*MafsNode{nodes[0]}
	minh := nodes[0].h
	for _, n := range nodes[1:] {
		if n.h < minh {
			minh = n.h
			best = []*MafsNode{n}
		} else if n.h == minh {
			best = append(best, n)
		}
	}
	return best[rand.Intn(len(best))]
}

// sendMessages sends state-messages to all relevant agents
func (bfs *Mafs) sendMessages(s state.State, n *MafsNode) {
	interestedAgents := bfs.succ.ApplicableAgents(bfs.stateRegistry.Public(s))
	if len(interestedAgents) > 0 {
		msg := &MafsMessage{bfs.agentID,
			[]int(bfs.stateRegistry.SharedState(s, n.stateID)), n.g, n.h}
		for _, toAgentID := range interestedAgents {
			bfs.dispatcher.Send(toAgentID, msg)
		}
		bfs.info.incMessagesOut(len(interestedAgents))
	}
}

// processMessages handles incoming messages depending on type/content
func (bfs *Mafs) processMessages() {
	var m interface{}
	for {
		select {
		case m = <-bfs.xMessages:
			bfs.planRegistry.processMessage(m.(*XMessage))
			continue
		default: // do not block, to process other message types
		}

		select {
		case m = <-bfs.stateMessages:
			bfs.processStateMessage(m.(*MafsMessage))
		default:
			return // inProgress
		}
	}
}

func (bfs *Mafs) processStateMessage(m *MafsMessage) {
	bfs.info.incMessagesIn(1)

	newState := bfs.stateRegistry.GlobalState(m.State)
	newSID := bfs.stateRegistry.Register(newState)
	a := bfs.Actions[int(m.SenderID)][0]

	// prune if state has worse cost than best found plans cost
	if m.G >= bfs.planRegistry.bestCost {
		return
	}

	tp, ok := bfs.closed[newSID]
	if !ok || m.G < tp.g {
		//h := int(ints.Max(bfs.heuristic.Evaluate(newState), m.H))
		h := bfs.heuristic.Evaluate(newState)
		newNode := bfs.NewNode(newSID, nil, a, m.G, h)
		bfs.open.Put(newNode)

		if Satisfies(newState, bfs.Goal) {
			bfs.goalNode = newNode
		}
	}
}

// register mafs strategies
func init() {
	// we need to tell gob about the concrete message type used by the search
	gob.Register(&MafsMessage{})
	// cmd options for MAFS strategies
	mafsOptions := func() *opt.OptionSet {
		opts := opt.NewOptionSet()
		opts.Add(opt.Option{Type: opt.String, Name: "heuristic", Short: 'h',
			DefaultValue: "ff", Description: "heuristic"})
		opts.Add(opt.Option{Type: opt.Int32, Name: "triallength", Short: 'l',
			DefaultValue: 1, Description: "the trial length"})
		opts.Add(opt.Option{Type: opt.Int32, Name: "costtype", Short: 0,
			DefaultValue: 2, Description: "NormalCost=0|UnitCost=1|NormalMinOne=2"})
		opts.Add(opt.NewOption(opt.Bool, "reopen_closed", 'r', false,
			"reopen closed nodes?"))
		opts.Add(opt.NewOption(opt.String, "evaluator", 'e', "greedy",
			"evaluator type, e.g. 'greedy' for gbfs or 'astar' for astar search"))
		opts.Add(opt.NewOption(opt.Bool, "lazy", 'z', false,
			"lazy heuristic evaluation?"))
		opts.Add(opt.NewOption(opt.String, "rollout-policy", 'p', "greedy",
			"successor node selection, e.g. greedy|random"))
		return opts
	}
	strategyRegistry = append(strategyRegistry, &StrategyInfo{
		Name:        "mafs-g",
		Description: "greedy multi-agent forward search (MAFS)",
		NewStrategy: NewMafs,
		Options:     mafsOptions(),
	})
}

// TODO: option for deferred vs immediate sending of state messages
// TODO: option for Max/Min/First/Second(heur.Evaluate(s), m.H)
