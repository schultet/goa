package pddl

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/schultet/goa/pkg/util/slice"
)

type testCase struct {
	GroundedPublicPredicates  int64
	GroundedPrivatePredicates int64
	GroundedOperators         int64
}

func TestGrounding(t *testing.T) {
	testfiles := [][]string{
		[]string{
			"testfiles/bloc-01-exp",
			"testfiles/bloc-01-dom.pddl",
			"testfiles/bloc-01-prob.pddl",
		},
		[]string{
			"testfiles/bloc-02-exp",
			"testfiles/bloc-02-dom.pddl",
			"testfiles/bloc-02-prob.pddl",
		},
		[]string{
			"testfiles/wire-01-exp",
			"testfiles/wire-01-dom.pddl",
			"testfiles/wire-01-prob.pddl",
		},
	}
	for _, test := range testfiles {
		file := test[0]
		dom, prob := Open(test[1], test[2])
		_, _, _ = file, dom, prob
		task := NewTask(dom, prob, 1)
		content, _ := readFile(file)
		var tc testCase
		err := json.Unmarshal(content, &tc)
		if err != nil {
			t.Errorf("%s\n", err)
		}
		_ = task

		pubAtoms, intAtoms := GroundPredicates(task.Types(), task.objects, task.predicates)
		t.Logf("Objects: #pub/#priv: %d/%d\n", len(task.PublicObjects()), len(task.PrivateObjects()))
		t.Logf("Predicates: #pub/#priv: %d/%d\n", len(task.PublicPredicates()), len(task.PrivatePredicates()))
		t.Logf("allpreds: %v", task.predicates)

		if len(pubAtoms) != int(tc.GroundedPublicPredicates) {
			t.Errorf("ALARM\nexp:%d\nwas:%d(%d)\n", tc.GroundedPublicPredicates, len(pubAtoms), len(intAtoms))
		}
		if len(intAtoms) != int(tc.GroundedPrivatePredicates) {
			t.Errorf("ALARM\nexp:%d\nwas:%d(%d)\n", tc.GroundedPublicPredicates, len(pubAtoms), len(intAtoms))
		}
	}
}

var (
	tm      TypeMap
	objects ObjectList
)

func init() {
	// TODO: more tests on type hierarchies -> types_test.go
	t2 := &Type{"b", nil, []*Type{}}
	t1 := &Type{"a", nil, []*Type{t2}}
	t2.parent = t1
	t3 := &Type{"c", nil, []*Type{}}
	t0 := &Type{"object", nil, []*Type{t1, t3}}
	t1.parent = t0
	t3.parent = t0
	t4 := &Type{"d", nil, []*Type{}}
	t4.parent = t0

	tm = TypeMap{
		t0.name: t0,
		t1.name: t1,
		t2.name: t2,
		t3.name: t3,
		t4.name: t4,
	}

	objects = ObjectList{
		*NewTypedObject("a1", "a"),
		*NewTypedObject("a2", "a"),
		*NewTypedObject("b1", "b"),
		*NewTypedObject("b2", "b"),
		*NewTypedObject("b3", "b"),
		*NewTypedObject("object", "object"),
		*NewTypedObject("c1", "c"),
		*NewTypedObject("c2", "c"),
		*NewTypedObject("c3", "c"),
		*NewTypedObject("d1", "d"),
	}

}

// _____________________________________________________________________________
func TestGroundAction(t *testing.T) {
	// TODO: implement this test, it is just printing values for manual
	// verification
	tests := []string{
		`( :action debark :parameters ( ?a - a ?b - b ?c - c ) 
		    :precondition ( and ( in ?b ?a ) ( at ?a ?c ) ) 
			:effect ( and ( at ?b ?c ) ( not ( in ?b ?a ) ) ) )`,
		`( :action drop :parameters ( ?x - d ?y - a )
			:precondition ( and ( at ?x ?y ) ( on ?x ?x ) )
			:effect ( and ( not ( at ?x ?y ) ( not ( on ?x ?x ) ) ) )`,
	}

	action, _ := ParseAction(Tokenize(tests[1]))
	groundedActions := groundAction(&action, objects, tm)
	// TODO: test grounded actions
	for i := range groundedActions {
		fmt.Println("\n", groundedActions[i])
		for _, con := range groundedActions[i].Precondition() {
			fmt.Printf("%v [%v]", con, con.predicate.parameters)
		}
	}

	// test based on taxi/p01
	objectsPddl := `(:objects g1 - location g2 - location c - location 
						  h1 - location h2 - location 
						  t1 t2  - taxi p1 p2  - passenger)`
	typesPddl := "(:types location agent - object taxi passenger - agent ) "
	predicatesPddl := `(:predicates 
		(directly-connected ?l1 - location ?l2 - location) 
		(at ?a - agent ?l - location) (in ?p - passenger ?t - taxi)
		(empty ?t - taxi) (free ?l - location) (:private
		(goal-of ?p - passenger ?l - location)))`
	actionsPddl := []string{
		` (:action enter_p1 :parameters (?t - taxi ?l - location) :precondition (and
		(at p1 ?l) (at ?t ?l) (empty ?t)) :effect (and (not (empty ?t))
		(not (at p1 ?l)) (in p1 ?t))) `,
		` (:action exit_p1 :parameters (?t - taxi ?l - location) :precondition (and
		(in p1 ?t) (at ?t ?l) (goal-of p1 ?l)) :effect (and (not (in p1 ?t)) 
		(empty ?t) (at p1 ?l))) `}

	objects, privateObjects, _ := ParseConstants(Tokenize(objectsPddl))
	types, _ := ParseTypes(Tokenize(typesPddl))
	predicates, privatePredicates, _ := ParsePredicates(Tokenize(predicatesPddl))
	actions := make([]Action, len(actionsPddl))
	for i, ap := range actionsPddl {
		a, _ := ParseAction(Tokenize(ap))
		actions[i] = a
	}
	_, _, _, _, _, _ = objects, privateObjects, types, predicates, privatePredicates, actions
	fmt.Println(objects)
	for i := range actions {
		groundActions := groundAction(&actions[i], objects, types)
		fmt.Printf("%s\n", actions[i].Name())
		for j := range groundActions {
			fmt.Printf("%+v\n", groundActions[j].Precondition())
		}
	}
}

// _____________________________________________________________________________
func TestGroundPredicate(t *testing.T) {
	predicates := []string{
		"(below ?x - a ?y - c)",
		"(at ?so - object ?what - d)",
	}
	expectedResult := [][]string{
		[]string{
			"(below a1 - a c1 - c)",
			"(below a1 - a c2 - c)",
			"(below a1 - a c3 - c)",
			"(below a2 - a c1 - c)",
			"(below a2 - a c2 - c)",
			"(below a2 - a c3 - c)",
			"(below b1 - b c1 - c)",
			"(below b1 - b c2 - c)",
			"(below b1 - b c3 - c)",
			"(below b2 - b c1 - c)",
			"(below b2 - b c2 - c)",
			"(below b2 - b c3 - c)",
			"(below b3 - b c1 - c)",
			"(below b3 - b c2 - c)",
			"(below b3 - b c3 - c)",
		},
		[]string{
			"(at object - object d1 - d)",
			"(at a1 - a d1 - d)",
			"(at a2 - a d1 - d)",
			"(at c1 - c d1 - d)",
			"(at c2 - c d1 - d)",
			"(at c3 - c d1 - d)",
			"(at b1 - b d1 - d)",
			"(at b2 - b d1 - d)",
			"(at b3 - b d1 - d)",
		},
	}

	ps := make(PredicateList, len(predicates))
	for i, p := range predicates {
		predicate := NewPredicate(Tokenize(p))
		ps[i] = *predicate
	}
	for i, predicate := range ps.Publics() {

		grounded, _ := groundPredicate(&predicate, objects, tm)
		resultsPddl := make([]string, len(grounded))
		for j := range resultsPddl {
			resultsPddl[j] = grounded[j].Pddl()
		}
		if !slice.ContainSameStrings(resultsPddl, expectedResult[i]) {
			t.Errorf("Error: Wrong grounding of predicate\nExp:%s\nWas:%s\n",
				expectedResult[i], resultsPddl)
		}
	}
}

// _____________________________________________________________________________
func TestCounter(t *testing.T) {
	c := newCounter([]int{2, 3, 1})
	expected := [][]int{
		[]int{0, 0, 0},
		[]int{0, 0, 1},
		[]int{0, 1, 0},
		[]int{0, 1, 1},
		[]int{0, 2, 0},
		[]int{0, 2, 1},
		[]int{0, 3, 0},
		[]int{0, 3, 1},
		[]int{1, 0, 0},
		[]int{1, 0, 1},
		[]int{1, 1, 0},
		[]int{1, 1, 1},
		[]int{1, 2, 0},
		[]int{1, 2, 1},
		[]int{1, 3, 0},
		[]int{1, 3, 1},
		[]int{2, 0, 0},
		[]int{2, 0, 1},
		[]int{2, 1, 0},
		[]int{2, 1, 1},
		[]int{2, 2, 0},
		[]int{2, 2, 1},
		[]int{2, 3, 0},
		[]int{2, 3, 1},
	}
	for i, _ := range make([]int, 24) {
		if !slice.ContainSameInts(c.counter, expected[i]) {
			t.Errorf("Counter Error at %v!\nexp:%v\nwas:%v\n", i, expected[i], c.counter)
		}
		c.increment()
	}

	c = newCounter([]int{1, 1})
	expected = [][]int{
		[]int{0, 0},
		[]int{0, 1},
		[]int{1, 0},
		[]int{1, 1},
	}
	for i, _ := range make([]int, 4) {
		if !slice.ContainSameInts(c.counter, expected[i]) {
			t.Errorf("Counter Error at %v!\nexp:%v\nwas:%v\n", i, expected[i], c.counter)
		}
		c.increment()
	}
}
