package pddl

import "fmt"

type Task struct {
	name, domain string
	types        TypeMap
	id           int

	objects    ObjectList
	predicates PredicateList
	actions    map[int][]Action

	goal                   LiteralList
	publicInitialLiterals  LiteralList
	privateInitialLiterals LiteralList
}

func NewTask(d *Domain, p *Problem, id int) *Task {
	t := &Task{
		name:   p.name,
		domain: d.name,
		types:  d.types,
		id:     id,
		objects: append(append(append(d.constants, p.objects...),
			d.privateConstants...), p.privateObjects...),
		predicates: append(d.predicates, d.privatePredicates...),
		actions:    make(map[int][]Action),
		goal:       p.Goal.Conjuncts(),
		publicInitialLiterals:  p.initialPredicates,
		privateInitialLiterals: p.privateInitialPredicates,
	}
	t.publicInitialLiterals.Map(func(l *Literal) { l.fixArgs(t.objects) })
	t.privateInitialLiterals.Map(func(l *Literal) { l.fixArgs(t.objects) })
	t.goal.Map(func(l *Literal) { l.fixArgs(t.objects) })
	t.actions[id] = d.actions
	return t
}

func (t *Task) PublicObjects() ObjectList        { return t.objects.Publics() }
func (t *Task) Objects() ObjectList              { return t.objects }
func (t *Task) PrivateObjects() ObjectList       { return t.objects.Privates() }
func (t *Task) PublicPredicates() PredicateList  { return t.predicates.Publics() }
func (t *Task) PrivatePredicates() PredicateList { return t.predicates.Privates() }
func (t *Task) Goal() LiteralList                { return t.goal }
func (t *Task) Types() TypeMap                   { return t.types }
func (t *Task) Actions() map[int][]Action        { return t.actions }

func (t *Task) PublicInitialLiterals() LiteralList {
	return t.publicInitialLiterals
}

func (t *Task) PrivateInitialLiterals() LiteralList {
	return t.privateInitialLiterals
}

// PublicActions computes the set of public action projection for a given domain
func (t *Task) PublicActions() []Action {
	result := make([]Action, len(t.actions[t.id]))
	for i, action := range t.actions[t.id] {
		result[i] = action.PublicProjection(t.id, t.predicates.Privates(), t.objects, t.types)
	}
	return result
}

func (t *Task) AddAction(id int, a Action) {
	if _, ok := t.actions[id]; ok {
		t.actions[id] = append(t.actions[id], a)
	} else {
		t.actions[id] = []Action{a}
	}
}

func (t *Task) AddActions(id int, as []Action) {
	if _, ok := t.actions[id]; ok {
		t.actions[id] = append(t.actions[id], as...)
	} else {
		t.actions[id] = make([]Action, len(as))
		for i, a := range as {
			t.actions[id][i] = a
		}
	}
}

func (t *Task) AddObject(o Object) {
	if !t.objects.contains(o.name) {
		t.objects = append(t.objects, o)
	}
}

func (t *Task) AddPredicate(p Predicate) {
	if _, err := t.predicates.Search(p.name); err != nil {
		t.predicates = append(t.predicates, p)
	}
}

func (t *Task) Dump() {
	fmt.Printf("task (domain): %s (%s)\n", t.name, t.domain)
	fmt.Printf("actions:\n%s\n", ActionList(t.actions[t.id]))
	// TODO: dump everything (verbose option)
}
