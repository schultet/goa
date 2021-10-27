package pddl

import (
	"fmt"
	"strings"
	"testing"
)

// _____________________________________________________________________________
func init() {
	fmt.Println("testing predicates...")
}

// _____________________________________________________________________________
func TestCopy(t *testing.T) {
	preds := []string{
		"(at ?x hans)",
		"(on t1 ?t2 ?t3)",
		"(below x y z)",
	}

	for _, tc := range preds {
		p := NewPredicate(Tokenize(tc))
		pp := NewPredicate(Tokenize(tc))
		copy1 := p.copy()
		copy2 := p.copy()
		copy3 := p.copy()

		copy1.name = "john"            // change name
		if len(copy2.parameters) > 0 { // change one parameter
			copy2.parameters[0] = Variable{"tj", "hooker"}
		}
		copy3.parameters = []Variable{} // change all parameters

		if p.name != pp.name {
			t.Errorf("Predicate Copy Error!")
		}
		for i, v := range pp.parameters {
			if p.parameters[i] != v {
				t.Errorf("Predicate Copy Error!")
			}
		}
	}
}

// _____________________________________________________________________________
func TestContainsTerm(t *testing.T) {
	type testcase struct {
		p   string // predicate
		t   string // term (variable or typed object)
		res bool   // expected result
	}

	tests := []testcase{
		{"(dc ?l1 - loc ?l2 - loc)", "?l2", true},
		{"(dc ?l1 - loc ?l2 - loc)", "?l3", false},
		{"(dc ?l1 - loc ?l2 - loc)", "?l3", false},
		{"(dc ?a ?b ?c)", "?c", true},
		{"(dc ?a ?b ?c)", "?d", false},
		{"(dc NORBERT ?b BLUEM)", "?b", true},
		{"(dc NORBERT ?b BLUEM)", "NORBERT", true},
		{"(dc NORBERT ?b BLUEM)", "BLUEM", true},
		{"(dc NORBERT ?b BLUEM)", "BLU", false},
		{"(dc ?b)", "", false},
		{"(dc )", "", false},
		{"(dc )", "?a", false},
	}

	for _, tc := range tests {
		p := NewPredicate(Tokenize(tc.p))
		term := Variable{tc.t, "kind"}
		res := tc.res
		if p.containsTerm(term) != res {
			t.Errorf("NONONO")
		}
	}
}

// _____________________________________________________________________________
func TestNewPredicate(t *testing.T) {
	//typetokens := strings.Fields(`( :types location agent - object taxi passenger - agent ) `)
	//types, _ := parseTypes(typetokens)
	// test 1
	tokens := strings.Fields("( directly-connected ?l1 - location ?l2 - location )")
	pred := NewPredicate(tokens)

	if pred.name != "directly-connected" {
		t.Error("predicate name does not match; expected 'directly-connected'"+
			" but was", pred.name)
	} else {
		variableNames := []string{"?l1", "?l2"}
		typeNames := []string{"location", "location"}
		for i, v := range pred.parameters {
			if variableNames[i] != v.name {
				t.Error("variable names do not match", v.name, variableNames[i])
			}
			if typeNames[i] != v.kind {
				t.Error("variable types do not match", v.kind, typeNames[i])
			}
		}
	}

	//tokens = strings.Fields(" ( at ?a - agent ?l - location ) ")

	//tokens = strings.Fields(" ( in ?p - passenger ?t - taxi ) ")

	// TODO: test predicate definition without types &| with mixed type/no-type
}
