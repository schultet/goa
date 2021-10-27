package task

import (
	"log"
	"strings"

	"github.com/schultet/goa/pkg/pddl"
)

// CostType specifies actions cost type (e.g. Normal vs Unit)
type CostType byte // TODO: rename to ActionCostType

const (
	// NormalCost is the actions cost from its PDDL definition
	NormalCost CostType = iota
	// UnitCost is one for all actions
	UnitCost
	// Regular action cost, but if an action costs less than one it gets cost 1
	NormalMinOne
)

// Effect is a (variable,value) pair used in action effects
// TODO: when conditional effects are implemented, Effect should become an
// interface and the concrete types should be named atomicEffect and
// conditionalEffect, both implementing the Effect interface
type Effect VariableValuePair

// Condition is a (variable,value) pair used in action preconditions
type Condition VariableValuePair

// Action represents a grounded PDDL action
type Action struct {
	ID            int
	PDDLAction    *pddl.Action
	Args          []*pddl.Object
	Owner         int
	Cost          int
	Preconditions []Condition
	Effects       []Effect
	Marked        bool // TODO: needed?
	Private       bool

	// TODO: this is a hack for now (impl. two action types)
	pre  []Condition // precondition without prevails
	prev []Condition // prevail conditions
}

func (o *Action) IsPublic() bool  { return !o.Private }
func (o *Action) IsPrivate() bool { return o.Private }

// Pre returns the actions precondition (without prevail conditions)
func (o *Action) Pre() []Condition {
	return o.pre
}

// Pre returns the actions prevail conditions
func (o *Action) Prev() []Condition {
	return o.prev
}

// AdjustedCost returns the action cost depending on the costType used.
func (o *Action) AdjustedCost(t CostType) int {
	switch t {
	case NormalCost:
		return o.Cost
	case UnitCost:
		return 1
	case NormalMinOne:
		if o.Cost < 1 {
			return 1
		}
		return o.Cost
	default:
		log.Fatalf("Invalid Action CostType!\n")
	}
	return 0
}

// Projection returns the action projection keeping only variables from vars
func (o *Action) Projection(vars []*Variable) *Action {
	pre := ConditionSet(o.Preconditions).Projection(vars)
	// TODO: implement Projection method for []Effect
	eff := make([]Effect, 0)
	for _, e := range o.Effects {
		good := false
		for _, v := range vars {
			if v.ID == e.Variable {
				good = true
			}
		}
		if good {
			eff = append(eff, e)
		}
	}
	return &Action{
		ID:            o.ID,
		PDDLAction:    o.PDDLAction,
		Args:          o.Args,
		Owner:         o.Owner,
		Cost:          o.Cost,
		Preconditions: pre,
		Effects:       eff,
		Marked:        o.Marked,
		Private:       o.Private,
	}
}

// Name returns the name of the actions underlying (pddl) action
func (o *Action) Name() string { return o.PDDLAction.Name() }

// String returns a string representation of the action
func (o *Action) String() string {
	argNames := make([]string, len(o.Args))
	for i := range o.Args {
		argNames[i] = o.Args[i].Name()
	}
	return o.Name() + "-" + strings.Join(argNames, "-")
}

// hasCondition returns true iff a condition is a precondition of the action
func (o *Action) hasCondition(c Condition) bool {
	for _, cond := range o.Preconditions {
		if cond == c {
			return true
		}
	}
	return false
}

// TODO: testing -> replace with ConditionSet.Satisfies(ConditionSet)
// conditionsSatisfied returns true iff each action precondition is in the set
// of conditions cs.
func (o *Action) conditionsSatisfied(cs []*Condition) bool {
	for i := range o.Preconditions {
		satisfied := false
		for _, match := range cs {
			if *match == o.Preconditions[i] {
				satisfied = true
			}
		}
		if !satisfied {
			return false
		}
	}
	return true
}

// DependsOn determines if action o's preconditions and action n's effects
// have common variables.
func (o *Action) DependsOn(n *Action) bool {
	for _, precondition := range o.Preconditions {
		for _, effect := range n.Effects {
			if precondition.Variable == effect.Variable {
				return true
			}
		}
	}
	return false
}

// ActionList is an array of action objects, that supports some
// filtering methods and implements sort.Interface
type ActionList []*Action

// ActionMap contains each agents list of actions (keys are agentIDs)
type ActionMap map[int]ActionList

func (aom ActionMap) Len() int {
	n := 0
	for _, v := range aom {
		n += len(v)
	}
	return n
}

func (aom ActionMap) Actions() ActionList {
	res := make(ActionList, aom.Len())
	for _, v := range aom {
		res = append(res, v...)
	}
	return res
}

// ActionList implements sort.Interface
func (ls ActionList) Len() int      { return len(ls) }
func (ls ActionList) Swap(i, j int) { ls[i], ls[j] = ls[j], ls[i] }

// Less returns true iff action at index i comes in lexicographic order before
// action at index j. TODO: fix slow down due to recomputation of o.String()
// method
func (ls ActionList) Less(i, j int) bool {
	return ls[i].String() < ls[j].String()
}

// Satisfacted returns two ActionLists. The first contains all actions
// of which all preconditions exist in the condition set cs, the second contains
// those actions where this is not the case.
func (ls ActionList) Satisfacted(cs []*Condition) (app, unapp ActionList) {
	for _, o := range ls {
		if o.conditionsSatisfied(cs) {
			app = append(app, o)
		} else {
			unapp = append(unapp, o)
		}
	}
	return
}

// WithCondition returns all actions in ActionList that contain condition c in
// their precondition and those who don't as two separate lists.
func (ls ActionList) WithCondition(c Condition) (res ActionList, rest ActionList) {
	for _, o := range ls {
		if o.hasCondition(c) {
			res = append(res, o)
		} else {
			rest = append(rest, o)
		}
	}
	return
}

// Projection returns an ActionList containing all actions projections
func (ls ActionList) Projection(vars []*Variable) ActionList {
	projections := make(ActionList, len(ls))
	for i, o := range ls {
		projections[i] = o.Projection(vars)
	}
	return projections
}

// Copy creates a copy of an ActionList
func (ls ActionList) Copy() ActionList {
	res := make(ActionList, len(ls))
	copy(res, ls)
	return res
}

// String returns a string representation of an ActionList
func (ls ActionList) String() string {
	s := ""
	for i := range ls {
		s += ls[i].String() + "\n"
	}
	return s
}

// ConditionSet is a list of Conditions with set-like behaviour
type ConditionSet []Condition

// Implement sort.Interface
func (q ConditionSet) Len() int           { return len(q) }
func (q ConditionSet) Less(i, j int) bool { return q[i].Variable < q[j].Variable }
func (q ConditionSet) Swap(i, j int)      { q[i], q[j] = q[j], q[i] }

//// ConsistentWith checks whether conditions in the set are consistent with the
//// mutexgroups and trivial mutexes (one variable must not have multiple values)
//func (cs ConditionSet) ConsistentWith(mg []MutexGroup) bool {
//	//vars := make(map[int]int, len(cs))
//	// TODO: implement
//	return true
//}

// Equal returns true iff two conditionsets are structurally equal
func (cs ConditionSet) Equals(other ConditionSet) bool {
	if len(cs) != len(other) {
		return false
	}
	seen := make(map[Condition]struct{})
	for _, c := range cs {
		seen[c] = struct{}{}
	}
	for _, c := range other {
		_, ok := seen[c]
		if !ok {
			return false
		}
	}
	return true
}

// Consistent returns true iff no two conditions refer to the same variable
func (cs ConditionSet) Consistent() bool {
	duplicates := make(map[int]struct{}, len(cs))
	for _, c := range cs {
		_, ok := duplicates[c.Variable]
		if ok {
			return false
		}
		duplicates[c.Variable] = struct{}{}
	}
	return true
}

// SubsetOf checks whether all conditions of a ConditionSet are contained in
// another ConditionSet
// TODO: rename Subsumes (referring to partial state subsumption in sas+)
// Complexity:  O(n), where n = max{len(other), len(cs)}
func (cs ConditionSet) SubsetOf(other ConditionSet) bool {
	visited := make(map[Condition]struct{}, 0)
	for _, c := range other {
		visited[c] = struct{}{}
	}
	for _, c := range cs {
		_, ok := visited[c]
		if !ok {
			return false
		}
	}
	return true
}

// Intersection returns the intersection of two ConditionSets
func (cs ConditionSet) Intersection(other ConditionSet) ConditionSet {
	inFirst := make(map[Condition]struct{}, len(cs))
	for _, c := range cs {
		inFirst[c] = struct{}{}
	}
	res := make(ConditionSet, 0)
	for _, c := range other {
		_, ok := inFirst[c]
		if ok {
			res = append(res, c)
		}
	}
	return res
}

// Add adds a condition to ConditionSet cs if it is not already contained in cs
// and returns the resulting set
func (cs ConditionSet) Add(c Condition) (ConditionSet, bool) {
	for _, cp := range cs {
		if cp.Variable == c.Variable {
			if cp.Value == c.Value {
				return cs, true
			}
			return cs, false
		}
	}
	return append(cs, c), true
}

func (cs ConditionSet) Except(vs ...int) ConditionSet {
	res := make(ConditionSet, 0, len(cs))
	for _, c := range cs {
		good := true
		for _, v := range vs {
			if v == c.Variable {
				good = false
			}
		}
		if good {
			res = append(res, c)
		}
	}
	return res
}

func (cs ConditionSet) Ints() []int {
	res := make([]int, 0, len(cs)*2)
	for _, c := range cs {
		res = append(res, c.Variable, int(c.Value))
	}
	return res
}

func (cs ConditionSet) Ints32() []int32 {
	res := make([]int32, 0, len(cs)*2)
	for _, c := range cs {
		res = append(res, int32(c.Variable), int32(c.Value))
	}
	return res
}

func (cs ConditionSet) IsSorted() bool {
	if len(cs) < 2 {
		return true
	}
	for i, j := 0, 1; j < len(cs); i, j = i+1, j+1 {
		if cs[j].Variable > cs[i].Variable {
			return false
		}
	}
	return true
}

// Projection returns the ConditionSet that only contains variables from vars
func (cs ConditionSet) Projection(vars []*Variable) ConditionSet {
	res := make(ConditionSet, 0)
	for ic, iv := 0, 0; ic < len(cs) && iv < len(vars); {
		if cs[ic].Variable < vars[iv].ID {
			ic += 1
		} else if cs[ic].Variable == vars[iv].ID {
			res = append(res, cs[ic])
			ic += 1
			iv += 1
		} else {
			iv += 1
		}
	}
	return res
}

// TODO: these methods are obsolete due to Project-method above
func (cs ConditionSet) PrivateConditions(t *Task) ConditionSet {
	res := make(ConditionSet, 0)
	for _, c := range cs {
		if t.Vars[c.Variable].IsPrivate {
			res = append(res, c)
		}
	}
	return res
}

func (cs ConditionSet) PublicConditions(t *Task) ConditionSet {
	res := make(ConditionSet, 0)
	for _, c := range cs {
		if !t.Vars[c.Variable].IsPrivate {
			res = append(res, c)
		}
	}
	return res
}

func CreateConditionSet(a []int) ConditionSet {
	res := make(ConditionSet, 0, len(a)/2+1)
	for i := 0; i < len(a); i += 2 {
		res = append(res, Condition{a[i], a[i+1]})
	}
	return res
}

func ConditionSetToIntSlice(cs ConditionSet) []int {
	res := make([]int, 0, len(cs)*2)
	for _, c := range cs {
		res = append(res, c.Variable, int(c.Value))
	}
	return res
}

func CreateEffects(a []int) []Effect {
	res := make([]Effect, 0, len(a)/2+1)
	for i := 0; i < len(a); i += 2 {
		res = append(res, Effect{a[i], a[i+1]})
	}
	return res
}

func EffectsToIntSlice(cs []Effect) []int {
	res := make([]int, 0, len(cs)*2)
	for _, c := range cs {
		res = append(res, c.Variable, int(c.Value))
	}
	return res
}

//func containsEffect(es []Effect, e Effect) bool {
//	for _, ep := range es {
//		if ep.Variable == e.Variable && ep.Value == e.Value {
//			return true
//		}
//	}
//	return false
//}
//
//func containsEffectFunc(es []Effect, p func(Effect) bool) bool {
//	for _, e := range es {
//		if p(e) {
//			return true
//		}
//	}
//	return false
//}
