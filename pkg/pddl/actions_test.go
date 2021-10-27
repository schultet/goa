package pddl

import (
	"fmt"
	"strings"
	"testing"

	"github.com/schultet/goa/pkg/util/slice"
)

func TestGroundCopy(t *testing.T) {
	type testcase struct {
		action string
		args   []int // object position in ObjectList
		expPre string
		expEff string
	}
	tests := []testcase{
		{`( :action A :parameters ( ?a - aircraft ?p - person ?c - city ) 
		    :precondition ( and ( in ?p ?a ) ( at ?a ?c ) ) 
			:effect ( and ( at ?p ?c ) ( not ( in ?p ?a ) ) ) )`,
			[]int{0, 2, 4},
			"(and (in P1 A1) (at A1 C1))",
			"(and (at P1 C1) (not (in P1 A1)))",
		},
		{`( :action B :parameters ( ?a - aircraft ?p - person ?c - city ) 
		    :precondition ( and ( in ?p Foo ) ( at ?a ?c ) ) 
			:effect ( and ( at ?p Bar ) ( not ( in ?p ?a ) ) ) )`,
			[]int{1, 3, 5},
			"(and (in P2 Foo) (at A2 C2))",
			"(and (at P2 Bar) (not (in P2 A2)))",
		},
		{`( :action C :parameters ( ?a - aircraft ?p - person ?c - city ) 
		    :precondition ( and ( in Bar Foo ) ( at Bar Bar ) ) 
			:effect ( and ( at Foo Bar ) ( not ( in Foo Foo ) ) ) )`,
			[]int{0, 2, 4},
			"(and (in Bar Foo) (at Bar Bar))",
			"(and (at Foo Bar) (not (in Foo Foo)))",
		},
		{`( :action D :parameters ( ) 
		    :precondition ( and ( in Bar Foo ) ( at Bar Bar ) ) 
			:effect ( and ( at Foo Bar ) ( not ( in Foo Foo ) ) ) )`,
			[]int{},
			"(and (in Bar Foo) (at Bar Bar))",
			"(and (at Foo Bar) (not (in Foo Foo)))",
		},
	}
	objects := ObjectList{
		{"A1", "aircraft", false},
		{"A2", "aircraft", false},
		{"P1", "person", false},
		{"P2", "person", false},
		{"C1", "city", false},
		{"C2", "city", false},
		{"C3", "city", false},
		{"Foo", "object", false},
		{"Bar", "object", false},
	}
	for _, tc := range tests {
		a, _ := ParseAction(Tokenize(tc.action))
		args := make([]*Object, len(tc.args))
		for i, arg := range tc.args {
			args[i] = &objects[arg]
		}
		gc := a.GroundCopy(args, objects)

		for i := range gc.args {
			if gc.args[i] != args[i] {
				t.Errorf("GroundAction arguments do not match!\nexp:%s\nwas:%s\n",
					args, gc.args)
			}
		}

		expPre := createFolFormula(Tokenize(tc.expPre)).Conjuncts()
		expPrePddl := make([]string, len(expPre))
		resPrePddl := make([]string, len(expPre))
		for i, literal := range expPre {
			expPrePddl[i] = literal.NameIdentifier()
			resPrePddl[i] = gc.precondition[i].NameIdentifier()
		}
		if !slice.ContainSameStrings(expPrePddl, resPrePddl) {
			t.Errorf("Wrong grounded Action Preconditions!\nexp:\nwas:\n",
				expPrePddl, resPrePddl)
		}

		expEff := createFolFormula(Tokenize(tc.expEff)).Conjuncts()
		expEffPddl := make([]string, len(expEff))
		resEffPddl := make([]string, len(expEff))
		for i, literal := range expEff {
			expEffPddl[i] = literal.NameIdentifier()
			resEffPddl[i] = gc.effect[i].NameIdentifier()
		}
		if !slice.ContainSameStrings(expEffPddl, resEffPddl) {
			t.Errorf("Wrong grounded Action Preconditions!\nexp:\nwas:\n",
				expEffPddl, resEffPddl)
		}
	}
}

// TODO: testcase for action with private variables in precondition
func TestPublicProjection(t *testing.T) {
	type testCase struct {
		in        string
		privpreds []string
		expected  string
	}

	tests := []testCase{
		{`( :action debark :parameters ( ?a - aircraft ?p - person ?c - city ) 
		    :precondition ( and ( in ?p ?a ) ( at ?a ?c ) ) 
			:effect ( and ( at ?p ?c ) ( not ( in ?p ?a ) ) ) )`,
			[]string{"(in ?p ?a)"},
			"(:action debark@1 :parameters (?a - aircraft ?p - person ?c - city) " +
				":precondition (at ?a ?c) " +
				":effect (at ?p ?c))"},
		{`( :action X :parameters ( ?a - aircraft ?p - person ?c - city ) 
		    :precondition ( and ( in ?p ?a ) ( at ?a ?c ) ) 
			:effect ( and ( at ?p ?c ) ( not ( in ?p ?a ) ) ) )`,
			[]string{"(at ?x ?y)"},
			"(:action X@1 :parameters (?a - aircraft ?p - person) " +
				":precondition (in ?p ?a) " +
				":effect (not (in ?p ?a)))"},
		{`( :action Y :parameters (?a - A) :precondition (at ?a B) :effect
		    (not (at ?a B)))`, []string{"(at ?x ?y)"},
			`(:action Y@1 :parameters () :precondition () :effect ())`},
	}

	tm, _ := ParseTypes(Tokenize("(:types locatable city flevel - object aircraft person - locatable )"))
	os, pos, _ := ParseConstants(Tokenize(`(:objects
		person2 - person person3 - person person1 - person person6 - person
		person7 - person person4 - person person5 - person city2 - city
		city3 - city person8 - person city1 - city city0 - city city4 - city
		city5 - city fl1 - flevel fl0 - flevel fl3 - flevel fl2 - flevel fl5 - flevel
		fl4 - flevel fl6 - flevel plane2 - aircraft (:private plane1 - aircraft)) `))

	for _, tc := range tests {
		action, _ := ParseAction(Tokenize(tc.in))
		privpreds := []Predicate{}
		for _, p := range tc.privpreds {
			privpreds = append(privpreds, *NewPredicate(Tokenize(p)))
		}
		predicates := PredicateList(privpreds)
		fmt.Println(predicates)
		objects := append(os, pos...)

		publicProjection := action.PublicProjection(1, predicates, objects, tm)
		if publicProjection.Pddl() != tc.expected {
			t.Errorf("wrong public projection\nexp:%s\nwas:%s\n", tc.expected,
				publicProjection.Pddl())
		}
	}
}

// _____________________________________________________________________________
func TestActionPddl(t *testing.T) {
	type testCase struct {
		in       string
		expected string
	}

	tests := []testCase{
		{`( :action debark :parameters ( ?a - aircraft ?p - person ?c - city ) 
		    :precondition ( and ( in ?p ?a ) ( at ?a ?c ) ) 
			:effect ( and ( at ?p ?c ) ( not ( in ?p ?a ) ) ) )
		`, "(:action debark :parameters (?a - aircraft ?p - person ?c - city) " +
			":precondition (and (in ?p ?a) (at ?a ?c)) " +
			":effect (and (at ?p ?c) (not (in ?p ?a))))"},
	}

	for _, tc := range tests {
		tl := strings.Fields(tc.in)
		action, _ := ParseAction(tl)
		if action.Pddl() != tc.expected {
			t.Errorf("wrong pddl representation\nexp:%s\nwas:%s\n", tc.expected, action.Pddl())
		}
	}

}

// _____________________________________________________________________________
func TestParseAction(t *testing.T) {
	type testCase struct {
		in                   string
		expectedName         string
		expectedParameters   string
		expectedPrecondition string
		expectedEffect       string
	}

	tests := []testCase{
		{`( :action enter
		:parameters ( ?t - taxi )
		:precondition ( at p1 ?t ) 
		:effect ( and ( not ( empty ?t ) ) ( at p1 ?t ) ) )`,
			"enter",
			"[?t]",
			"(at p1 ?t)",
			"(and (not (empty ?t)) (at p1 ?t))"},
		{`( :action exit
			:parameters ( ) )`,
			"exit",
			"[]",
			"",
			""},
	}

	for _, tc := range tests {
		tokens := strings.Fields(tc.in)
		a, _ := ParseAction(tokens)
		// check name
		if a.name != tc.expectedName {
			t.Error("Action name does not match")
		}
		// check parameters
		parameters := []string{}
		for _, v := range a.parameters {
			parameters = append(parameters, v.name)
		}
		if fmt.Sprintf("%v", parameters) != tc.expectedParameters {
			t.Errorf("Parameters dont match!\nexp:%s\ngot:%v\n",
				tc.expectedParameters, parameters)
		}
		//check precondition
		preconditionStr := ""
		if a.precondition != nil {
			preconditionStr = a.precondition.String()
		}
		if preconditionStr != tc.expectedPrecondition {
			t.Error("Precondition does not match")
		}
		// check effect
		effectStr := ""
		if a.effect != nil {
			effectStr = a.effect.String()
		}
		if effectStr != tc.expectedEffect {
			t.Errorf("Effect does not match!\nexp:%s\ngot:%s\n",
				tc.expectedEffect, effectStr)
		}
	}
}
