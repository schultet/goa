package pddl

import "errors"

type Variable struct {
	name string
	kind string
}

func (v Variable) Name() string { return v.name }

// Returns a PDDL-string representation of a Variable
func (v Variable) Pddl() string {
	if v.kind == "" {
		return v.name
	}
	return v.name + " - " + v.kind
}

// Returns a string representation of a Variable
func (t Variable) String() string {
	typeStr := "untyped"
	if t.kind != "" {
		typeStr = t.kind
	}
	return "Variable{" + t.name + ", " + typeStr + "}"
}

type Predicate struct {
	name       string
	parameters []Variable
	private    bool
}

func (p Predicate) Name() string    { return p.name }
func (p Predicate) IsPrivate() bool { return p.private }

// Returns a copy of a given predicate
func (p *Predicate) copy() *Predicate {
	params := make([]Variable, len(p.parameters))
	for i := range params {
		params[i] = Variable{p.parameters[i].name, p.parameters[i].kind}
	}
	return &Predicate{p.name, params, p.private}
}

// Creates a new Predicate from a token list
func NewPredicate(tl TokenList) *Predicate {
	i := 0
	for ; tl[i] == "("; i++ { // skip opening braces
	}
	return &Predicate{
		name:       tl[i],
		parameters: extractParameters(tl[i+1:]),
	}
	//pred.parameters = extractParameters(tl[i+1:])
}

// Returns a string representation of a Predicate
func (t Predicate) String() string {
	return t.Pddl()
}

// Returns a PDDL-string representation of a Predicate
func (p Predicate) Pddl() string {
	s := "(" + p.name
	for _, v := range p.parameters {
		s += " " + v.Pddl()
	}
	return s + ")"
}

// returns true when a predicate contains variable v
func (p Predicate) containsTerm(v Variable) bool {
	for _, variable := range p.parameters {
		if variable.name == v.name {
			return true
		}
	}
	return false
}

// returns the list of all variables in the predicates parameter list
func (p Predicate) variables() []Variable {
	res := make([]Variable, 0, len(p.parameters))
	for _, v := range p.parameters {
		if isVar(v.name) {
			res = append(res, v)
		}
	}
	return res
}

// PredicateList implements sort.Interface
type PredicateList []Predicate

// methods to implement sort.Interface
func (pl PredicateList) Len() int           { return len(pl) }
func (pl PredicateList) Swap(i, j int)      { pl[i], pl[j] = pl[j], pl[i] }
func (pl PredicateList) Less(i, j int) bool { return pl[i].Name() < pl[j].Name() }

// Insert predicate p at position i into the predicatelist
func (pl *PredicateList) Insert(p Predicate, i int) {
	*pl = append(*pl, p)
	copy((*pl)[i+1:], (*pl)[i:])
	(*pl)[i] = p
}

// filter returns a list of predicates that match the filter criterion f
func (pl PredicateList) filter(f func(Predicate) bool) PredicateList {
	matches := make(PredicateList, 0)
	for i := range pl {
		if f(pl[i]) {
			matches = append(matches, pl[i])
		}
	}
	return matches
}

// Returns all private Predicates
func (pl PredicateList) Privates() PredicateList {
	return pl.filter(Predicate.IsPrivate)
}

// Returns all public Predicates
func (pl PredicateList) Publics() PredicateList {
	return pl.filter(func(p Predicate) bool { return !p.IsPrivate() })
}

// find searches through the predicate list until a predicate matches the
// pattern or the end of the list is reached.
func (pl PredicateList) find(f func(Predicate) bool) (*Predicate, error) {
	for i := range pl {
		if f(pl[i]) {
			return &pl[i], nil
		}
	}
	return nil, errors.New("Predicate not found!")
}

// Search searches a predicate list for a predicate that matches a given name
func (pl PredicateList) Search(name string) (*Predicate, error) {
	return pl.find(func(p Predicate) bool {
		if p.name == name {
			return true
		}
		return false
	})
}
