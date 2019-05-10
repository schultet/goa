package task

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"path"
	"sort"
	"strings"

	"github.com/schultet/goa/src/pddl"
	"github.com/schultet/goa/src/state"
)

// Variable represents a multi-valued state variable of a planning task
type Variable struct {
	ID        int
	Name      string
	DomRange  int32
	FactName  []string
	IsPrivate bool
}

// VariableValuePair contains a variable ID (int) and an intended value (byte)
type VariableValuePair struct {
	Variable int
	Value    int
}

// VariableAssignment is a value assignment over all variables. The index
// corresponds to the variable ID, the value is the assigned value.
type VariableAssignment []int

// MutexGroup is a list of Variable-Value Pairs
type MutexGroup []*VariableValuePair

func (mg MutexGroup) String() string {
	s := "["
	for _, m := range mg {
		s += fmt.Sprintf("{%d,%d},", m.Variable, m.Value)
	}
	return s[:len(s)-1] + "]"
}

// TODO: sas.Task interface
//type Interface interface {
//	NumVariables() int
//	VariableName(v int) string
//	VariableRange(v int) int
//	ActionCost(index int) int
//	ActionName(index int) string
//	NumActions() int
//	GetAction(index int) *Action
//	InitialStateValues() []int
//	//FactName(Fact) string
//	//AreFactsMutex(Fact, Fact) bool
//	//ConvertStateValues(values []int, task *TTask)
//
//	// type Fact VariableValuePair
//}

// Task contains all the information of a planning task.
type Task struct {
	pddlActions map[string]*pddl.Action // TODO: why are these pointers?
	pddlObjects map[string]*pddl.Object // ^
	varIndex    map[int]int

	AgentID     int
	Vars        []*Variable
	Init        state.State
	Goal        ConditionSet
	Actions     ActionMap
	ActionsByID map[int]*Action
	MutexGroups []MutexGroup // [Variable.ID]MutexGroup
}

func (t *Task) GetActions() ActionList {
	return t.Actions[t.AgentID]
}

func (t *Task) GetAction(actionID int) *Action {
	return t.ActionsByID[actionID]
}

func (t *Task) ShallowCopy() *Task {
	return &Task{
		pddlActions: t.pddlActions,
		pddlObjects: t.pddlObjects,
		varIndex:    t.varIndex,

		AgentID:     t.AgentID,
		Vars:        t.Vars,
		Init:        t.Init,
		Goal:        t.Goal,
		Actions:     t.Actions,
		MutexGroups: t.MutexGroups,
	}
}

func (t *Task) SetAgentActions(agentID int, os ActionList) {
	t.Actions[agentID] = os
}

func (t *Task) AddAction(agentID int, o *Action) {
	t.Actions[agentID] = append(t.Actions[agentID], o)
}

func (t *Task) Variables() []*Variable {
	return t.Vars
}

// VariableRanges returns the domainSize/range of each variable in a vector.
// Index corresponds to variable.ID
func (t *Task) VariableRanges() []int {
	ranges := make([]int, len(t.Vars))
	for i, v := range t.Vars {
		ranges[i] = int(v.DomRange)
	}
	return ranges
}

// BuildMacroProjections returns an ActionList containing the sets of macro
// projections for each public action.
//func (t *Task) BuildMacroProjections() ActionList {
//	ops := t.Actions[t.AgentID]
//	res := make(ActionList, 0)
//	varsPublic := t.PublicVars()
//	for _, o := range ops {
//		if o.Private {
//			continue
//		}
//		regr := NewRegressionSearch(o, ops, -1, varsPublic)
//		regr.Search()
//		n := len(res)
//		res = append(res, regr.GetMacroProjections()...)
//		if len(res) == n {
//			res = append(res, o.Projection(varsPublic))
//		}
//	}
//	return res
//}

// PublicVars returns the list of variables that are public
func (t *Task) PublicVars() []*Variable {
	return t.GetVarsF(func(v *Variable) bool { return !v.IsPrivate })
}

// PrivateVars returns the list of variables that are private
func (t *Task) PrivateVars() []*Variable {
	return t.GetVarsF(func(v *Variable) bool { return v.IsPrivate })
}

// GetVarsF returns the list of variables for which p returns true
func (t *Task) GetVarsF(p func(*Variable) bool) []*Variable {
	res := make([]*Variable, 0, len(t.Vars))
	for i := range t.Vars {
		if p(t.Vars[i]) {
			res = append(res, t.Vars[i])
		}
	}
	return res
}

// GetVarIDsF returns the list of IDs of variables for which p returns true
func (t *Task) GetVarIDsF(p func(*Variable) bool) []int {
	res := make([]int, 0, len(t.Vars))
	for i := range t.Vars {
		if p(t.Vars[i]) {
			res = append(res, t.Vars[i].ID)
		}
	}
	return res
}

// GetPrivateVarIDs returns the list of IDs of private variables
func (t *Task) GetPrivateVarIDs() []int {
	return t.GetVarIDsF(func(v *Variable) bool { return v.IsPrivate })
}

func (t *Task) String() string {
	vars := make([]string, len(t.Vars))
	for i, v := range t.Vars {
		vars[i] = fmt.Sprintf("%+v\n", v)
	}
	return fmt.Sprintf("AGENT:%d\nGOAL:%+v\nINIT:%v\nOPS:%+v\nVARS:%+v\n"+
		"MUTEXGROUPS:%+v\n",
		t.AgentID, t.Goal, t.Init, t.Actions, vars, t.MutexGroups)
}

// PrintState prints state s
func (t *Task) PrintState(s VariableAssignment) {
	for i, v := range s {
		fmt.Printf("%d:%s: %s\n", i, t.Vars[i].Name, t.Vars[i].FactName[v])
	}
}

// NewTaskFromJSON converts a byte-slice of the supported json format into the
// corresponding planning Task representation
func NewTaskFromJSON(data []byte) *Task {
	var jsonTask interface{}
	err := json.Unmarshal(data, &jsonTask)
	if err != nil {
		log.Fatalf("Error unmarshalling json data: %s\n", err)
	}

	content := jsonTask.(map[string]interface{})
	id := int(content["AgentID"].(float64))
	//log.Println(id, content["AgentID"].(float64), content["AgentName"].(string))
	task := &Task{
		pddlActions: make(map[string]*pddl.Action),
		pddlObjects: make(map[string]*pddl.Object),
		varIndex:    make(map[int]int),
		AgentID:     id,
	}
	// NOTE: variables must be translated first to set the varIndex
	task.translateVars(content["Variables"].([]interface{}))
	task.translateInit(content["InitialState"].([]interface{}))
	task.translateGoal(content["Goal"].([]interface{}))
	task.translateOps(id,
		content["Operators"].([]interface{}),
		content["ProjectedOperators"].([]interface{}))
	task.translateMutexGroups(content["MutexGroups"].([]interface{}))

	task.ActionsByID = make(map[int]*Action)
	for _, actionList := range task.Actions {
		for _, a := range actionList {
			task.ActionsByID[a.ID] = a
		}
	}

	// TODO: make this optional
	for _, action := range task.ActionsByID {
		sortAction(action)
	}

	return task
}

func sortAction(a *Action) {
	sort.Slice(a.Preconditions, func(i, j int) bool {
		return a.Preconditions[i].Variable < a.Preconditions[j].Variable
	})
	sort.Slice(a.Effects, func(i, j int) bool {
		return a.Effects[i].Variable < a.Effects[j].Variable
	})
}

// NewTaskFromFile reads a problem file and returns the corresponding planning
// Task if the file format (e.g. .json) is supported
func NewTaskFromFile(f string) (*Task, error) {
	data, err := ioutil.ReadFile(f)
	if err != nil {
		return nil, fmt.Errorf("Error reading json file: %s\n", err)
	}
	if ext := path.Ext(f); ext == ".json" {
		return NewTaskFromJSON(data), nil
	} else {
		return nil, fmt.Errorf("Unsupported file format: %s\n", ext)
	}
}

func (t *Task) translateOp(op map[string]interface{}) *Action {
	fields := strings.Fields(op["Name"].(string))
	action := pddl.DummyAction(fields[0])
	if a, ok := t.pddlActions[fields[0]]; ok {
		action = a
	} else {
		t.pddlActions[fields[0]] = action
	}

	args := make([]*pddl.Object, len(fields)-1)
	for i, name := range fields[1:] {
		dob := pddl.DummyObject(name)
		if obj, ok := t.pddlObjects[name]; ok {
			args[i] = obj
		} else {
			t.pddlObjects[name] = dob
			args[i] = dob
		}
	}

	id := int(op["GlobalID"].(float64))
	prevail := op["Prevail"].([]interface{})
	prepost := op["Prepost"].([]interface{})
	preconditions := make([]Condition, 0, len(prevail)+len(prepost))
	prev := make([]Condition, 0, len(prevail))
	for _, p := range prevail {
		cond := p.(map[string]interface{})
		c := Condition{
			Variable: t.varIndex[int(cond["Var"].(float64))],
			Value:    int(cond["Val"].(float64)),
		}
		preconditions = append(preconditions, c)
		prev = append(prev, c)
	}
	pre := make([]Condition, 0, len(prepost))
	for _, p := range prepost {
		cond := p.(map[string]interface{})
		if val := cond["Pre"].(float64); val >= 0 { // is pre != -1 ?
			c := Condition{
				Variable: t.varIndex[int(cond["Var"].(float64))],
				Value:    int(val),
			}
			preconditions = append(preconditions, c)
			pre = append(pre, c)
		}
	}
	effects := make([]Effect, len(prepost))
	for i, e := range prepost {
		eff := e.(map[string]interface{})
		effects[i] = Effect{
			Variable: t.varIndex[int(eff["Var"].(float64))],
			Value:    int(eff["Post"].(float64)),
		}
	}

	return &Action{
		ID:         id,
		PDDLAction: action,
		Args:       args,
		Owner:      int(op["Owner"].(float64)),
		//Cost:       1, //int(op["Cost"].(float64)),
		Cost:          int(op["Cost"].(float64)),
		Preconditions: preconditions,
		Effects:       effects,
		Marked:        false,
		Private:       op["IsPrivate"].(bool),

		pre:  pre,
		prev: prev,
	}
}

func (t *Task) translateOps(
	agentID int, ops, projectedOps []interface{}) {
	t.Actions = make(ActionMap)

	for i := range ops {
		o := ops[i].(map[string]interface{})
		owner := int(o["Owner"].(float64))
		if _, ok := t.Actions[owner]; !ok {
			t.Actions[owner] = make(ActionList, 0)
		}
		t.Actions[owner] = append(t.Actions[owner], t.translateOp(o))
	}

	for i := range projectedOps {
		o := projectedOps[i].(map[string]interface{})
		owner := int(o["Owner"].(float64))
		if owner != agentID { // skip agents own projections
			if _, ok := t.Actions[owner]; !ok {
				t.Actions[owner] = make(ActionList, 0)
			}
			t.Actions[owner] = append(t.Actions[owner], t.translateOp(o))
		}
	}
}

func (t *Task) translateMutexGroups(mgs []interface{}) {
	t.MutexGroups = make([]MutexGroup, 0, len(mgs))
	for _, mutexgroup := range mgs {
		mg := make(MutexGroup, 0)
		for _, mutex := range mutexgroup.([]interface{}) {
			m := mutex.(map[string]interface{})
			mg = append(mg, &VariableValuePair{
				Variable: t.varIndex[int(m["Var"].(float64))],
				Value:    int(m["Val"].(float64)),
			})
		}
		t.MutexGroups = append(t.MutexGroups, mg)
	}
}

func (t *Task) translateVars(vars []interface{}) {
	t.Vars = make([]*Variable, 0, len(vars))
	var priv []*Variable
	for i, variable := range vars {
		v := variable.(map[string]interface{})
		private := v["IsPrivate"].(bool)
		vp := &Variable{
			ID:        i,
			Name:      v["Name"].(string),
			DomRange:  int32(v["Range"].(float64)),
			IsPrivate: private,
		}
		factname := v["Factname"].([]interface{})
		vp.FactName = make([]string, len(factname))
		for i := range factname {
			vp.FactName[i] = factname[i].(string)
		}
		if private {
			priv = append(priv, vp)
		} else {
			t.Vars = append(t.Vars, vp)
		}
	}
	t.Vars = append(t.Vars, priv...) // ensure private vars are last
	for i, v := range t.Vars {       // set variable index (pos of var in state)
		t.varIndex[v.ID] = i
	}
}

func (t *Task) translateInit(init []interface{}) {
	t.Init = make(state.State, len(init))
	for i, val := range init {
		t.Init[t.varIndex[i]] = int(val.(float64))
	}
}

func (t *Task) translateGoal(goal []interface{}) {
	t.Goal = make(ConditionSet, len(goal))
	for i := range goal {
		cond := goal[i].(map[string]interface{})
		t.Goal[i] = Condition{
			Variable: t.varIndex[int(cond["Var"].(float64))],
			Value:    int(cond["Val"].(float64)),
		}
	}
}
