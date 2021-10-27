package pddl

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

// Action data type represents a PDDL action
// implements Component interface
type Action struct {
	name         string
	parameters   []Variable
	precondition FolFormula
	effect       FolFormula
	cost         int
}

func (a Action) ParameterNames() []string {
	names := make([]string, len(a.parameters))
	for i := range a.parameters {
		names[i] = a.parameters[i].name
	}
	return names
}

func (a Action) PreconditionLiterals() LiteralList {
	return a.precondition.Conjuncts()
}

func (a Action) EffectLiterals() LiteralList {
	return a.effect.Conjuncts()
}

// GroundCopy returns a grounded copy of an action
func (a *Action) GroundCopy(args []*Object, objects ObjectList) *GroundAction {
	gc := &GroundAction{
		action:       a,
		args:         args,
		precondition: a.PreconditionLiterals(),
		effect:       a.EffectLiterals(),
	}
	paramToArg := make(map[string]*Object)
	for i, p := range gc.action.parameters {
		paramToArg[p.name] = gc.args[i]
	}

	for i := range gc.precondition {
		gc.precondition[i].insertArgs(paramToArg, objects)
	}

	for i := range gc.effect {
		gc.effect[i].insertArgs(paramToArg, objects)
	}
	// TODO: important testing (this was buggy)

	//	gc.precondition.Map(func(l *Literal) {
	//		l.insertArgs(paramToArg, objects)
	//	})
	//	gc.effect.Map(func(l *Literal) { l.insertArgs(paramToArg, objects) })

	return gc
}

// Getter
func (a Action) Name() string             { return a.name }
func (a Action) Precondition() FolFormula { return a.precondition }
func (a Action) Effect() FolFormula       { return a.effect }
func (a Action) Cost() int                { return a.cost }

// Returns the actions name and its parameters as string
func (a Action) NameIdentifier() string {
	paramNames := make([]string, len(a.parameters))
	for i, param := range a.parameters {
		paramNames[i] = param.Name()
	}
	return a.name + "-" + strings.Join(paramNames, "-")
}

// Creates a new action where all private predicates are removed, and where the
// agentId is appended to the actions name.
func (a Action) PublicProjection(agentId int, privates PredicateList,
	objects ObjectList, tm TypeMap) Action {

	precondition := copyFolFormula(a.precondition)
	effect := copyFolFormula(a.effect)

	// remove literals with private predicates
	for i := range privates {
		precondition = removeLiteral(precondition, Literal{false, &privates[i], nil})
		effect = removeLiteral(effect, Literal{false, &privates[i], nil})
	}
	// remove suberfluous parameters
	parameters := make([]Variable, 0, len(a.parameters))
	for _, p := range a.parameters {
		if containsVariable(precondition, p) ||
			containsVariable(effect, p) {
			if len(objects.ofTypeAndSubtypes(tm[p.kind])) > 0 { // TODO: fix inefficient impl.
				// there is at least one object in os that fits the parameters type
				parameters = append(parameters, p)
			} else {
				// remove all literals that contain occurences of p
				for l, ok := nextLiteralWithV(precondition, p); ok; {
					precondition = removeLiteral(precondition, l)
					l, ok = nextLiteralWithV(precondition, p)
				}
				// TODO ^ unit tests for above case
			}
		}
	}

	name := a.name + "@" + strconv.Itoa(agentId)
	return Action{name, parameters, precondition, effect, a.cost}
}

// Creates a PDDL Action from a TokenList. Returns the action and the rest of
// the token list.
func ParseAction(tl TokenList) (Action, TokenList) {
	tokens, rest := extractParanthesis(tl)
	if tokens[1] != ":action" {
		log.Fatalf("error while parsing action: %v", tokens[:2])
	}
	action := Action{}
	action.name = tokens[2]
	tokens = tokens[3:] // skip "(" and ":action" and <action-name>

	for len(tokens) > 0 {
		switch tokens[0] {
		case ":parameters":
			var params TokenList
			params, tokens = extractParanthesis(tokens[1:])
			action.parameters = extractParameters(params)
		case ":precondition":
			var fol TokenList
			fol, tokens = extractParanthesis(tokens[1:])
			action.precondition = createFolFormula(fol)
		case ":effect":
			var fol TokenList
			fol, tokens = extractParanthesis(tokens[1:])
			action.effect = createFolFormula(fol)
		default:
			tokens = tokens[1:]
		}
	}

	return action, rest
}

// Returns a "deep-enough" copy of an Action (needed for grounding)
func (a Action) copy() Action {
	params := []Variable{}
	for _, v := range a.parameters {
		params = append(params, v)
	}
	cond := copyFolFormula(a.precondition)
	eff := copyFolFormula(a.effect)
	return Action{a.name, params, cond, eff, a.cost}
}

// Returns a string representation of a pddl action
func (a Action) String() string {
	params := []string{}
	for _, v := range a.parameters {
		params = append(params, v.name)
	}
	return fmt.Sprintf("A{%s %v\npre:%s\neff:%s}", a.name, params, a.precondition, a.effect)
}

// Returns a PDDL representation of an action
func (a Action) Pddl() string {
	var params string
	if len(a.parameters) > 0 {
		params = ":parameters ("
		for _, v := range a.parameters {
			params += v.Pddl() + " "
		}
		params = params[:len(params)-1] + ")"
	} else {
		params = ":parameters ()"
	}

	precond := ":precondition " + a.precondition.Pddl()
	effect := ":effect " + a.effect.Pddl()

	return fmt.Sprintf("(:action %s %s %s %s)", a.name, params, precond, effect)
}

func DummyAction(name string) *Action { return &Action{name: name} }

type GroundAction struct {
	action       *Action
	args         []*Object
	precondition LiteralList
	effect       LiteralList
}

func (ga *GroundAction) Cost() int                 { return ga.action.cost }
func (ga *GroundAction) Action() *Action           { return ga.action }
func (ga *GroundAction) Args() []*Object           { return ga.args }
func (ga *GroundAction) Precondition() LiteralList { return ga.precondition }
func (ga *GroundAction) Effect() LiteralList       { return ga.effect }

func (ga *GroundAction) NameIdentifier() string {
	s := ga.action.name
	for i, arg := range ga.args {
		if arg == nil {
			s += "-" + ga.action.parameters[i].name
		} else {
			s += "-" + arg.name
		}
	}
	return s
}

func (ga *GroundAction) String() string {
	return ga.NameIdentifier()
}

func (ga *GroundAction) Pddl() string {
	// TODO
	return ""
}

// List of PDDL Actions
type ActionList []Action

// ActionList implements sort.Interface
func (l ActionList) Len() int           { return len(l) }
func (l ActionList) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }
func (l ActionList) Less(i, j int) bool { return l[i].NameIdentifier() < l[j].NameIdentifier() }

func (l ActionList) String() string {
	names := make([]string, len(l))
	for i, a := range l {
		names[i] = a.NameIdentifier()
	}
	return strings.Join(names, "\n")
}
