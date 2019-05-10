package pddl

import (
	"log"
	"strings"
)

type folError struct {
	s string
}

func (e folError) Error() string {
	return e.s
}

func (e folError) String() string {
	return e.s
}

type FolFormula interface {
	String() string
	Pddl() string
	Conjuncts() LiteralList
}

// returns whether a given string is a variable (starts with a ?)
func isVar(n string) bool {
	return strings.HasPrefix(n, "?")
}

// Creates a first order logic formula from a TokenList tl.
func createFolFormula(tl TokenList) FolFormula {
	if tl[0] != "(" {
		log.Fatal(folError{"error parsing fol formula"})
	}
	if tl[1] == "and" {
		c := Conjunction{}
		tl, rest := extractParanthesis(tl[2:])
		for {
			c.conjuncts = append(c.conjuncts, createFolFormula(tl))
			if len(rest) == 0 || rest[0] == ")" {
				break
			}
			tl, rest = extractParanthesis(rest)
		}
		return c
	} else {
		tl, _ := extractParanthesis(tl)
		return createLiteral(tl)
	}
}

// Creates a literal from a list of tokens.
func createLiteral(tl TokenList) Literal {
	literal := Literal{negated: false}
	if tl[0] != "(" {
		log.Fatal(folError{"error parsing literal, must start with '('" +
			" but started with: " + tl[0]})
	}
	i := 1
	if tl[i] == "not" {
		literal.negated = true
		i++
	}
	if len(tl) <= 2 {
		literal.predicate = &Predicate{}
		literal.args = make([]*Object, 0)
	} else {
		literal.predicate = NewPredicate(tl[i:])
		literal.args = make([]*Object, len(literal.predicate.parameters))
	}
	return literal
}

// Constructs a copy of a fol formula
func copyFolFormula(folFormula FolFormula) FolFormula {
	return createFolFormula(Tokenize(folFormula.Pddl()))
}

// Creates a new FolFormula with a specified literal removed.
func removeLiteral(f FolFormula, literal Literal) FolFormula {
	folFormula := copyFolFormula(f)
	switch formula := folFormula.(type) {
	case Literal:
		if formula.Name() == literal.Name() {
			return Literal{false, &Predicate{}, []*Object{}}
		}
	case Conjunction:
		for i, c := range formula.conjuncts {
			switch conjunct := c.(type) {
			default:
				log.Fatalf("Error while removing literal from folformula!\n"+
					"Conjunct Type not recognized: %T", conjunct)
			case Conjunction:
				// recursive call
				formula.conjuncts[i] = removeLiteral(conjunct, literal)
				folFormula = formula
			case Literal:
				// remove literal from conjunct list
				if conjunct.Name() == literal.Name() {
					if i < len(formula.conjuncts)-1 {
						formula.conjuncts = append(formula.conjuncts[:i],
							formula.conjuncts[i+1:]...)
						i -= 1
					} else {
						formula.conjuncts = formula.conjuncts[:i]
					}

					if len(formula.conjuncts) < 2 {
						// its not a conjunction anymore
						return removeLiteral(formula.conjuncts[0], literal)
					} else {
						return removeLiteral(formula, literal)
					}
				}
			}
		}
	}
	return folFormula
}

// Substitutes object o for variable v in first order formula f. Returns a new
// FolFormula where all occurences of the variable are replaced by the object.
func folSubstitute(f FolFormula, v Variable, o Object) FolFormula {
	return createFolFormula(Tokenize(
		strings.Replace(f.Pddl(), v.name, o.name, -1)))
}

// TODO: this should be implemented more efficient
func multiFolSubstitute(f FolFormula, vs []Variable, os []Object) FolFormula {
	fPddl := f.Pddl()
	for i, v := range vs {
		fPddl = strings.Replace(fPddl, v.name, os[i].name, -1)
	}
	return createFolFormula(Tokenize(fPddl))
}

// TODO: testing
// returns true if formula f contains variable v, otherwise false
func containsVariable(f FolFormula, v Variable) bool {
	folFormula := copyFolFormula(f) // TODO: copy needed?
	switch formula := folFormula.(type) {
	case Literal:
		if formula.predicate.containsTerm(v) {
			return true
		}
	case Conjunction:
		for _, c := range formula.conjuncts {
			if containsVariable(c, v) {
				return true
			}
		}
	}
	return false
}

// TODO: testing, more descriptive name
// Returns first literal in f containing variable v
func nextLiteralWithV(f FolFormula, v Variable) (Literal, bool) {
	folFormula := copyFolFormula(f)
	switch formula := folFormula.(type) {
	case Literal:
		if formula.predicate.containsTerm(v) {
			return formula, true
		}
	case Conjunction:
		for _, c := range formula.conjuncts {
			if l, succ := nextLiteralWithV(c, v); succ {
				return l, true
			}
		}
	}
	return Literal{}, false
}

type Conjunction struct {
	conjuncts []FolFormula
}

// TODO: testing
func (c Conjunction) Conjuncts() LiteralList {
	conjuncts := make(LiteralList, 0)
	for i := range c.conjuncts {
		conjuncts = append(conjuncts, c.conjuncts[i].Conjuncts()...)
	}
	return conjuncts
}

// Returns a string representation of the conjunction
func (c Conjunction) String() string {
	return c.Pddl()
}

// Returns a PDDL representation of the conjunction
func (c Conjunction) Pddl() string {
	s := "(and "
	for _, c := range c.conjuncts {
		s += c.Pddl() + " "
	}
	return s[:len(s)-1] + ")"
}

// Negation is a negated FolFormula
type Negation struct {
	arg FolFormula
}

func (n Negation) Pddl() string           { return n.arg.Pddl() }
func (n Negation) String() string         { return n.Pddl() }
func (n Negation) Conjuncts() LiteralList { return n.arg.Conjuncts() }

type Literal struct {
	negated   bool
	predicate *Predicate
	args      []*Object
}

func (l Literal) Name() string  { return l.predicate.name }
func (l Literal) Negated() bool { return l.negated }

func (l Literal) NameIdentifier() string {
	pnames := make([]string, len(l.args))
	for i, o := range l.args {
		if o == nil {
			pnames[i] = l.predicate.parameters[i].Name()
			//log.Fatalf("LiteralError:\n%s\n%v\n%s\n", l, l.args, *l.predicate)
			continue
		}
		pnames[i] = o.Name()
	}
	return l.Name() + "-" + strings.Join(pnames, "-")
}

// TODO: testing
func (l Literal) Conjuncts() LiteralList {
	lcpy := &Literal{l.negated, l.predicate, make([]*Object, len(l.args))}
	for i := range lcpy.args {
		lcpy.args[i] = l.args[i]
	}
	return LiteralList{lcpy}
}

// Returns a string representation of the literal
func (l Literal) String() string {
	return l.Pddl()
}

// Returns a PDDL representation of the literal
func (l Literal) Pddl() string {
	s := "("
	if l.negated {
		s = "(not ("
	}
	s += l.predicate.name
	for i, o := range l.args {
		if o == nil {
			s += " " + l.predicate.parameters[i].Pddl()
		} else {
			s += " " + o.Pddl()
		}
	}
	if l.negated {
		s += ")"
	}
	return s + ")"
}

// A list of Literals
type LiteralList []*Literal

// LiteralList implements sort.Interface
func (l LiteralList) Len() int           { return len(l) }
func (l LiteralList) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }
func (l LiteralList) Less(i, j int) bool { return l[i].NameIdentifier() < l[j].NameIdentifier() }

// initP initializes a LiteralList with values of a []*Predicate
func (l *LiteralList) initP(ps PredicateList) *LiteralList {
	*l = make(LiteralList, len(ps))
	for i := range ps {
		(*l)[i] = &Literal{
			negated:   false,
			predicate: &ps[i],
			args:      make([]*Object, len(ps[i].parameters)),
		}
	}
	return l
}

// Returns each Literal's NameIdentifier followed by a newline
func (l LiteralList) String() string {
	names := make([]string, len(l))
	for i, e := range l {
		names[i] = e.NameIdentifier()
	}
	return strings.Join(names, "\n")
}

// applies f to all literals in literal list
func (l *LiteralList) Map(f func(*Literal)) {
	for i := range *l {
		f((*l)[i])
	}
}

// TODO: doc
func (l *Literal) fixArgs(objects ObjectList) {
	for j, param := range l.predicate.parameters {
		if !isVar(param.name) {
			o, err := objects.Search(param.name)
			if err == nil {
				l.args[j] = o
			}
		}
	}
}

// insertArgs sets the literals arguments to the objects of pToA or of the
// ObjectList. pToA is a mapping between parameter names (string) and objects.
// The ObjectList has to be provided in case the underlying predicate contains
// parameters that are not in pToA, which is the case when the predicate is
// "partially grounded" (ie. contains one or more object names instead of
// variable names)
func (l *Literal) insertArgs(pToA map[string]*Object, objects ObjectList) {
	for j, p := range l.predicate.parameters {
		if arg, ok := pToA[p.name]; ok {
			l.args[j] = arg
		} else {
			o, err := objects.Search(p.name)
			if err != nil {
				log.Fatalf("Action grounding error!\n")
			}
			l.args[j] = o
		}
	}
}
