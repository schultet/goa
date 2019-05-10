package pddl

import (
	"log"
)

// Grounds all actions of all agents and all predicates
func Ground(task *Task) (map[int][]*GroundAction, LiteralList, LiteralList) {
	tm := task.Types()
	pubAtoms, intAtoms := GroundPredicates(tm, task.objects, task.predicates)

	groundActions := make(map[int][]*GroundAction)

	for agentID, actions := range task.actions {
		if agentID == task.id {
			groundActions[agentID] = GroundActions(actions, task.objects, tm)
		} else {
			groundActions[agentID] = GroundActions(actions, task.PublicObjects(), tm)
		}
	}
	return groundActions, pubAtoms, intAtoms
}

// TODO: testing
// Grounds all actions
func GroundActions(actions []Action, objects ObjectList, tm TypeMap) []*GroundAction {
	result := make([]*GroundAction, 0)
	for i := range actions {
		result = append(result, groundAction(&actions[i], objects, tm)...)
	}
	return result
}

// TODO: testing
// Grounds a single action, returns an array of grounded actions, where each
// precondition and effect formula are atomic formula (i.e. grounded)
func groundAction(a *Action, objects ObjectList, tm TypeMap) []*GroundAction {
	variables := a.parameters
	objectField := make([]ObjectList, len(variables))
	objectCount := make([]int, len(variables))
	n := 1
	for i, v := range variables {
		objectField[i] = objects.ofTypeAndSubtypes(tm[v.kind])
		l := len(objectField[i])
		if l == 0 {
			log.Printf("No object to ground action!\n"+
				"Action:%s:%s\nObjects:%s\nField:%s\n",
				a.name, a.parameters, objects, objectField)
			return []*GroundAction{}
			// TODO: test for this case
		}
		n *= l
		objectCount[i] = l - 1
	}

	result := make([]*GroundAction, 0, n)
	c := newCounter(objectCount)

	for i := 0; i < n; i++ {
		args := make([]*Object, len(c.counter))
		for j, cnt := range c.counter {
			args[j] = &objectField[j][cnt]
		}
		result = append(result, a.GroundCopy(args, objects))
		c.increment()
	}

	return result
}

// Grounds all predicates
func GroundPredicates(tm TypeMap, objects ObjectList,
	predicates PredicateList) (LiteralList, LiteralList) {

	publicAtoms, privateAtoms := make(LiteralList, 0), make(LiteralList, 0)

	for i := range predicates {
		pub, priv := groundPredicate(&predicates[i], objects, tm)
		publicAtoms = append(publicAtoms, pub...)
		privateAtoms = append(privateAtoms, priv...)
	}

	return publicAtoms, privateAtoms
}

// Grounds a single predicate, returns an array of Literals, where each literal
// represents a grounded instance of the predicate
func groundPredicate(p *Predicate, objects ObjectList, tm TypeMap) (
	pub LiteralList, priv LiteralList) {

	objectField := make([]ObjectList, len(p.parameters))
	objectCount := make([]int, len(p.parameters))
	n := 1
	for i, v := range p.parameters {
		objectField[i] = objects.ofTypeAndSubtypes(tm[v.kind])
		l := len(objectField[i])
		if l == 0 {
			log.Printf("there should always be an object to ground predicate")
			return LiteralList{}, LiteralList{}
		}
		n *= l
		objectCount[i] = l - 1
	}

	c := newCounter(objectCount)
	for i := 0; i < n; i++ {
		l := &Literal{
			negated:   false,
			predicate: p,
			args:      make([]*Object, len(p.parameters)),
		}
		isPrivate := false
		for j, count := range c.counter {
			if objectField[j][count].IsPrivate() {
				isPrivate = true
			}
			l.args[j] = &objectField[j][count]
		}
		if isPrivate || p.IsPrivate() {
			priv = append(priv, l)
		} else {
			pub = append(pub, l)
		}
		c.increment()
	}

	return
}

// counter is used to enumerate all combinations of a vector of int values
type counter struct {
	counter []int
	limits  []int
}

// newCounter creates a new counter
func newCounter(limits []int) counter {
	return counter{
		counter: make([]int, len(limits)),
		limits:  limits,
	}
}

// increment increments the counter by one
func (c *counter) increment() []int {
	i := len(c.counter) - 1
	for i >= 0 {
		if c.counter[i] < c.limits[i] {
			c.counter[i] += 1
			return c.counter
		} else {
			c.counter[i] = 0
			i -= 1
		}
	}
	return c.counter
}
