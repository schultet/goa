package search

import (
	"encoding/gob"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"strings"
	"time"

	"github.com/schultet/goa/src/comm"
	"github.com/schultet/goa/src/heuristic"
	"github.com/schultet/goa/src/opt"
	"github.com/schultet/goa/src/state"
	"github.com/schultet/goa/src/task"
	"github.com/schultet/goa/src/util/ints"
)

const printInfoInterval = 13 // print search progress info every k seconds

// StateActionMessage contains a search state, actions, and other information
type StateActionMessage struct {
	SenderID int
	State    []int // SharedState
	Actions  []int
	Costs    []int
	H        int
}

func (m *StateActionMessage) G() int {
	res := 0
	for _, v := range m.Costs {
		res += v
	}
	return res
}

// TextMessage is the message type for sending strings
type TextMessage struct {
	SenderID int
	Msg      string
}

// PlanExtractionMessage is used for distributed plan extraction procedure. The message means
// that during planning, the sender (with senderID) has received a state from
// the other agent containing stateID in the tokenlist. The other agent can
// lookup the stateID and tokens to continue plan extraction from there.
type PlanExtractionMessage struct {
	SenderID int
	Step     int
	Tokens   []int
	Msg      string
}

// Engine holds all the data structures needed for performing the actual search.
// It implements a high level Search algorithm, that repetitively calls the
// step() method of the Strategy object.
type Engine struct {
	Strategy
	*task.Task

	agentID       int
	conns         comm.ConnList
	succ          ActionEvaluator
	actionPruner  ActionPruner
	stateRegistry state.StateRegistry
	server        comm.Server
	dispatcher    comm.Dispatcher
	goalNode      Node
	info          *ProgressInfo
	logger        *log.Logger
	planLimit     int

	planRegistry *PlanRegistry

	textMessages    chan interface{}
	extractMessages chan interface{}
	xMessages       chan interface{}
}

// NewEngine creates a new search engine.
func NewEngine(t *task.Task, server comm.Server, dispatcher comm.Dispatcher,
	conns comm.ConnList, opts *opt.OptionSet) *Engine {
	//rand.Seed(opts.GetInt32("seed")) // TODO: << make option
	agentID := t.AgentID
	tokenRanges := make([]int, len(conns)-1)
	for i := range tokenRanges {
		tokenRanges[i] = ints.MaxValue //int(^uint32(0) >> 1) //ints.MaxValue
	}
	ranges := append(t.VariableRanges(), tokenRanges...)
	packer := state.NewIntPacker(ranges)

	var w io.Writer = os.Stdout
	if logf := opts.GetString("logfile"); logf != "" {
		filename := path.Join(path.Dir(logf), fmt.Sprintf("%d-%s", agentID, path.Base(logf)))
		f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err == nil {
			w = f
		}
	}
	logger := log.New(w, fmt.Sprintf("@Agent%2d: ", agentID), 0)
	logger.Println("Config:", opts.GetString("search"))

	e := &Engine{
		Task:            t,
		agentID:         agentID,
		conns:           conns,
		server:          server,
		dispatcher:      dispatcher,
		succ:            NewActionEvaluator(t.Vars, t.Actions, agentID),
		stateRegistry:   state.NewGStateRegistry(len(t.PublicVars()), len(t.PrivateVars()), agentID, packer),
		info:            NewProgressInfo(agentID, conns, logger),
		logger:          logger,
		planLimit:       int(opts.GetInt32("planlimit")),
		textMessages:    make(chan interface{}, 100),
		extractMessages: make(chan interface{}, 100),
		xMessages:       make(chan interface{}, 100),
	}
	e.planRegistry = NewPlanRegistry(e)
	e.server.RegisterMessageChan(&TextMessage{}, e.textMessages)
	e.server.RegisterMessageChan(&PlanExtractionMessage{}, e.extractMessages)
	e.server.RegisterMessageChan(&XMessage{}, e.xMessages)

	h := opts.GetString("heuristic")
	heuristic.Initialize(h, t) // only if uninitialized
	e.Strategy = NewStrategy(e, strings.Fields(opts.GetString("search")))
	e.actionPruner = getActionPruner(opts, e.Task)

	return e
}

// Search is the engine's main search routine, calling the implemented
// Strategy's step() function until status changes to solved, failed, or timeout
func (e *Engine) Search(searchTimeout time.Duration) {
	e.logger.Println("starts search.")
	e.info.Print()
	e.Initialize() // initialize search explicitly, init engine earlier..
	defer e.dispatcher.Quit()
	defer e.server.Quit()

	tstart := time.Now()
	status := inProgress
	for status == inProgress || status == idle { //|| status == solved {
		status = e.Step()
		if time.Since(e.info.lastPrint).Seconds() > printInfoInterval {
			e.info.PrintCurrent()
		}
		if time.Since(tstart) > searchTimeout {
			status = timeout
		}
		if e.planRegistry.NumPlans() >= e.planLimit {
			break
		}
	}

	terminateSearch(e, status)
}

// terminate terminates the search accordingly based on status
func terminateSearch(e *Engine, status Status) {
	switch status {
	case solved, timeout:
		// TODO: solved is always true because of the following two function
		// calls -> bug
	case failed:
		e.logger.Println("FAILED")
	default:
		e.logger.Println("DEFAULT termination...")
		e.info.Print()
	}

	// print pruning statistics
	e.logger.Println("PRUNING:")
	switch pruner := e.actionPruner.(type) {
	case *StubbornSetsSimple:
		pruner.printStatistics()
	default:
		fmt.Println("no pruning done")
	}
	// this is a hack to make sure that everything written to stdout is
	// written to the resp. log-file before the job terminates (GKI GRID).
	time.Sleep(2 * time.Second)
}

// SetStrategy sets the engine's search strategy to s
func (e *Engine) SetStrategy(s Strategy) { e.Strategy = s }

// SetDispatcher sets the engine's message dispatcher to d
func (e *Engine) SetDispatcher(d comm.Dispatcher) { e.dispatcher = d }

// SetServer sets the engine's server to s
func (e *Engine) SetServer(s comm.Server) { e.server = s }

// SetComm sets the engine's communication server and dispatcher
func (e *Engine) SetComm(s comm.Server, d comm.Dispatcher) {
	e.server = s
	e.dispatcher = d
}

// Satisfies returns true iff the state satisfies all conditions in cs
func Satisfies(s state.State, cs task.ConditionSet) bool {
	for i := range cs {
		if s[cs[i].Variable] != int(cs[i].Value) {
			return false
		}
	}
	return true
}

// Successor computes the successor state that results from applying action o
// to state s. TODO: move somewhere more suitable (maybe mpt.StateRegistry
func Successor(s state.State, o *task.Action) state.State {
	succ := s.Copy()
	for _, eff := range o.Effects {
		succ[eff.Variable] = int(eff.Value)
	}
	return succ
}

func init() {
	// we have to register all of the used message types
	gob.Register(&TextMessage{})
	gob.Register(&XMessage{})
	gob.Register(&StateActionMessage{})
}
