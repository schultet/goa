package search

import (
	"fmt"
	"log"
	"sort"

	"github.com/schultet/goa/src/opt"
	"github.com/schultet/goa/src/state"
	"github.com/schultet/goa/src/task"
)

// ActionPruner interface type to implement various pruning techniques (like
// stubborn sets) for pruning actions
type ActionPruner interface {
	Prune(*state.State, []*task.Action) []*task.Action
}

// PhonyPruner does not prune anything
type PhonyPruner struct{}

// NewPhonyPruner returns a new PhonyPruner instance
func NewPhonyPruner(t *task.Task) *PhonyPruner {
	return &PhonyPruner{}
}

// Prune does not prune anything (its just the identity of ops)
func (*PhonyPruner) Prune(s *state.State, ops []*task.Action) []*task.Action {
	return ops
}

// StubbornSetsSimple implements simple stubborn set computation
type StubbornSetsSimple struct {
	*StubbornSets

	interferenceRelation [][]int
}

// NewStubbornSetsSimple builds and returns a new StubbornSetsSimple object
func NewStubbornSetsSimple(t *task.Task) *StubbornSetsSimple {
	ss := &StubbornSetsSimple{
		StubbornSets: NewStubbornSets(t),
	}
	ss.computeInterferenceRelation()
	fmt.Println("pruning method: stubborn sets simple")
	return ss
}

// applicablePublic returns all action IDs that are applicable in the public
// projection of state s. Preconditions on private variables are ignored.
func (ss *StubbornSets) applicablePublic(s *state.State) []int {
	app := make([]int, 0)
	spub := (*s)[:ss.numPublicVars]
Outer:
	for opID, preconditions := range ss.sortedOpPreconditions {
		if ss.opByID[opID].Private {
			continue
		}
	Inner:
		for _, pre := range preconditions {
			if pre.Variable >= ss.numPublicVars {
				break Inner // can do this b/c preconditions are sorted
			}
			//if !pre.Satisfies(spub) {
			if int(pre.Value) != spub[pre.Variable] {
				continue Outer // don't add opID to app
			}
		}
		for _, eff := range ss.sortedOpEffects[opID] {
			if int(eff.Value) != (*s)[eff.Variable] {
				app = append(app, opID)
				break
			}
		}
	}
	return app
}

func (ss *StubbornSetsSimple) addNES(f FactPair) {
	for _, opID := range ss.achievers[f.Variable][f.Value] {
		ss.markAsStubborn(opID)
	}
}
func (ss *StubbornSetsSimple) addInterfering(opID int) {
	for _, intOpID := range ss.interferenceRelation[opID] {
		ss.markAsStubborn(intOpID)
	}
}

func (ss *StubbornSetsSimple) computeInterferenceRelation() {
	ss.interferenceRelation = make([][]int, ss.numActions)

	for i := 0; i < ss.numActions; i++ {
		interferenceOp1 := make([]int, 0) // TODO: may init. with cap=numOps
		for j := 0; j < ss.numActions; j++ {
			if i != j && ss.interfere(i, j) {
				interferenceOp1 = append(interferenceOp1, j)
			}
		}
		ss.interferenceRelation[i] = interferenceOp1
	}
}

func (ss *StubbornSetsSimple) initializeStubbornSet(s *state.State) {
	//unsatisfiedGoal := ss.findUnsatisfiedGoal(s)
	//if unsatisfiedGoal.noFact() {
	//	log.Fatalf("OH NOES")
	//}
	//ss.addNES(unsatisfiedGoal)

	// compute all public actions applicable in the public projection of state s
	//ops := ss.applicablePublic(s)
	//if len(ops) == 0 {
	//	return
	//}
	//ss.markAsStubborn(ops[0])
	//ss.markAsStubborn(ops[rand.Intn(len(ops))])
	for _, o := range ss.applicablePublic(s) {
		ss.markAsStubborn(o)
	}
}

// handleStubbornAction adds a NES for an unsatisfied precondition of the
// action or (if all conditions are satisfied) all actions that interfere
// with the action.
func (ss *StubbornSetsSimple) handleStubbornAction(s *state.State, opID int) {
	unsatisfiedPrecondition := ss.findUnsatisfiedPrecondition(opID, s)
	if unsatisfiedPrecondition.noFact() {
		// all preconditions are satisfied -> add interfering action
		ss.addInterfering(opID)
	} else {
		// at least one precondition is not satisfied -> add NES for the first
		// unsatisfied precondition of the action
		ss.addNES(unsatisfiedPrecondition)
	}
}

func (ss *StubbornSetsSimple) interfere(op1ID, op2ID int) bool {
	return ss.canDisable(op1ID, op2ID) ||
		ss.canConflict(op1ID, op2ID) ||
		ss.canDisable(op2ID, op1ID)
}

// Prune takes a state and a set of actions and returns a subset of these
// actions. Actions not contained in that set can be pruned during search.
func (ss *StubbornSetsSimple) Prune(s *state.State, actions []*task.Action) []*task.Action {
	ss.numUnprunedSuccessorsGenerated += int64(len(actions))
	// clear stubborn set from previous call
	ss.stubborn = make([]bool, len(ss.stubborn))
	ss.initializeStubbornSet(s) // TODO: call from `subclass`
	for len(ss.stubbornQueue) > 0 {
		opID := ss.stubbornQueue[len(ss.stubbornQueue)-1]
		ss.stubbornQueue = ss.stubbornQueue[:len(ss.stubbornQueue)-1]
		ss.handleStubbornAction(s, opID) // TODO: call from `subclass`
	}

	remainingOps := make([]*task.Action, 0, len(actions))
	for _, o := range actions {
		if ss.stubborn[ss.opID(o)] {
			remainingOps = append(remainingOps, o)
		}
	}
	ss.numPrunedSuccessorsGenerated += int64(len(remainingOps))
	return remainingOps
}

// StubbornSets is the `superclass` type, that is embedded by all concrete
// StubbornSets types (like StubbornSetsSimple or StubbornSetsEC). It provides
// all common methods (and instance variables) between the different
// implementations of stubborn sets.
type StubbornSets struct {
	numUnprunedSuccessorsGenerated int64
	numPrunedSuccessorsGenerated   int64
	numActions                     int

	sortedOpPreconditions [][]FactPair
	sortedOpEffects       [][]FactPair
	sortedGoals           []FactPair
	achievers             [][][]int

	stubborn      []bool
	stubbornQueue []int

	numPublicVars int
	relativeOpID  map[int]int
	opByID        map[int]*task.Action
}

// NewStubbornSets creates a new StubbornSets object
func NewStubbornSets(t *task.Task) *StubbornSets {
	ss := &StubbornSets{
		numActions:                     len(t.GetActions()),
		numUnprunedSuccessorsGenerated: 0,
		numPrunedSuccessorsGenerated:   0,
		sortedOpPreconditions:          make([][]FactPair, len(t.GetActions())),
		sortedOpEffects:                make([][]FactPair, len(t.GetActions())),
		stubborn:                       make([]bool, len(t.GetActions())),
		stubbornQueue:                  make([]int, 0),
		relativeOpID:                   make(map[int]int, 0),
		opByID:                         make(map[int]*task.Action, 0),
		numPublicVars:                  len(t.PublicVars()),
	}
	// add action preconditions and effects
	for i, o := range t.GetActions() {
		ss.sortedOpPreconditions[i] = make([]FactPair, len(o.Preconditions))
		for j, p := range o.Preconditions {
			ss.sortedOpPreconditions[i][j] = FactPair(p)
		}
		ss.sortedOpEffects[i] = make([]FactPair, len(o.Effects))
		for j, e := range o.Effects {
			ss.sortedOpEffects[i][j] = FactPair(e)
		}
		ss.relativeOpID[o.ID] = i
		ss.opByID[i] = o
	}

	// TODO: test goals are sorted
	ss.sortedGoals = make([]FactPair, len(t.Goal))
	for i := range ss.sortedGoals {
		ss.sortedGoals[i] = FactPair(t.Goal[i])
	}
	sort.Slice(ss.sortedGoals, func(i, j int) bool {
		return ss.sortedGoals[i].Variable < ss.sortedGoals[j].Variable
	})

	ss.computeSortedActions(t)
	ss.computeAchievers(t)
	return ss
}

func (ss *StubbornSets) opID(o *task.Action) int {
	return ss.relativeOpID[o.ID]
}

// FactPair represents a variable value-assignment (its basically the same as
// task.VariableValuePair or task.Condition)
type FactPair struct {
	Variable int
	Value    int
}

func (f FactPair) noFact() bool { return f.Variable == -1 }

// TODO: test
func (ss *StubbornSets) computeSortedActions(t *task.Task) {
	for opID := range ss.sortedOpPreconditions {
		sort.Slice(ss.sortedOpPreconditions[opID], func(i, j int) bool {
			return ss.sortedOpPreconditions[opID][i].Variable < ss.sortedOpPreconditions[opID][j].Variable
		})
	}
	for opID := range ss.sortedOpEffects {
		sort.Slice(ss.sortedOpEffects[opID], func(i, j int) bool {
			return ss.sortedOpEffects[opID][i].Variable < ss.sortedOpEffects[opID][j].Variable
		})
	}
}

func (ss *StubbornSets) computeAchievers(t *task.Task) {
	ss.achievers = make([][][]int, len(t.Vars))
	for i := range ss.achievers {
		ss.achievers[t.Vars[i].ID] = make([][]int, t.Vars[i].DomRange)
	}
	for _, o := range t.GetActions() {
		for _, eff := range o.Effects {
			ss.achievers[eff.Variable][eff.Value] = append(
				ss.achievers[eff.Variable][eff.Value], ss.opID(o))
		}
	}
}

func (ss *StubbornSets) canDisable(op1ID, op2ID int) bool {
	return containConflictingFact(ss.sortedOpEffects[op1ID], ss.sortedOpPreconditions[op2ID])
}

func (ss *StubbornSets) canConflict(op1ID, op2ID int) bool {
	return containConflictingFact(ss.sortedOpEffects[op1ID], ss.sortedOpEffects[op2ID])
}

func (ss *StubbornSets) findUnsatisfiedGoal(s *state.State) FactPair {
	return findUnsatisfiedCondition(ss.sortedGoals, s)
}

// findUnsatisfiedPrecondition returns the first action precondition that is
// unsatisfied in s (preconditions are sorted by variable ID)
func (ss *StubbornSets) findUnsatisfiedPrecondition(opID int, s *state.State) FactPair {
	return findUnsatisfiedCondition(ss.sortedOpPreconditions[opID], s)
}

func (ss *StubbornSets) markAsStubborn(opID int) bool {
	if !ss.stubborn[opID] {
		ss.stubborn[opID] = true
		ss.stubbornQueue = append(ss.stubbornQueue, opID)
		return true
	}
	return false
}

func (ss *StubbornSets) printStatistics() {
	fmt.Print("total successors before partial-order reduction: ")
	fmt.Printf("%d\n", ss.numUnprunedSuccessorsGenerated)
	fmt.Print("total successors after partial-order reduction: ")
	fmt.Printf("%d\n", ss.numPrunedSuccessorsGenerated)
}

func containConflictingFact(facts1, facts2 []FactPair) bool {
	for i, j := 0, 0; i < len(facts1) && j < len(facts2); {
		if facts1[i].Variable < facts2[j].Variable {
			i++
		} else if facts1[i].Variable > facts2[j].Variable {
			j++
		} else {
			if facts1[i].Value != facts2[j].Value {
				return true
			}
			i++
			j++
		}
	}
	return false
}

// findUnsatisfiedCondition returns the first FactPair (condition) from
// conditions, that is unsatisfied in state s. If all conditions are satisfied,
// a `phony` FactPair (with variable -1) is returned
func findUnsatisfiedCondition(conditions []FactPair, s *state.State) FactPair {
	for _, c := range conditions {
		if (*s)[c.Variable] != int(c.Value) {
			return c
		}
	}
	return FactPair{-1, 0} // TODO: no fact -> handle properly
}

func getActionPruner(opts *opt.OptionSet, t *task.Task) ActionPruner {
	// select pruning method based on cmd argument
	var pruner ActionPruner
	switch opts.GetString("pruning") {
	case "none":
		pruner = NewPhonyPruner(t)
	case "sss":
		pruner = NewStubbornSetsSimple(t)
	default:
		log.Fatalf("%s\n", opts.GetString("pruning"))
	}
	return pruner
}
