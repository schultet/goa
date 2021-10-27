package pddl

import (
	"fmt"
	"strings"
	"testing"
)

func init() {
	fmt.Println("testing fol formula")
}

func TestFolSubstitute(t *testing.T) {
	type testcase struct {
		f   string
		v   string
		o   string
		exp string
	}

	tests := []testcase{
		{"(not (on-table ?a ?b))", "?b", "Table1",
			"(not (on-table ?a Table1))",
		},
		{"(and (on ?x ?y) (at ?x ?z))", "?x", "Cat",
			"(and (on Cat ?y) (at Cat ?z))",
		},
		{"(and (on ?x ?y) (at ?x ?z))", "?a", "Cat",
			"(and (on ?x ?y) (at ?x ?z))",
		},
		{"(and (this ?x ?x ?x) (not (at ?x ?x)))", "?x", "35",
			"(and (this 35 35 35) (not (at 35 35)))",
		},
	}

	for _, tc := range tests {
		inF := createFolFormula(Tokenize(tc.f))
		outF := folSubstitute(inF, Variable{tc.v, ""}, *NewTypedObject(tc.o, ""))
		if outF.Pddl() != tc.exp {
			t.Errorf("ALARM:\nexp:%s\nwas:%s\n", tc.exp, outF.Pddl())
		}

	}
}

func TestMultiFolSubstitute(t *testing.T) {
	type testcase struct {
		f   string
		v   []string
		o   []string
		exp string
	}

	tests := []testcase{
		{"(not (on-table ?a ?b))",
			[]string{"?b", "?a"},
			[]string{"Table1", "Block4"},
			"(not (on-table Block4 Table1))",
		},
		{"(and (on ?x ?y) (at ?x ?z))",
			[]string{"?x"},
			[]string{"Cat"},
			"(and (on Cat ?y) (at Cat ?z))",
		},
		{"(and (on ?x ?y) (at ?x ?z))",
			[]string{"?a", "?z", "?y"},
			[]string{"Cat", "Dog", "Giraffe"},
			"(and (on ?x Giraffe) (at ?x Dog))",
		},
		{"(and (this ?x ?x ?x) (not (at ?x ?x)))",
			[]string{"?x", "?x"},
			[]string{"35", "70"},
			"(and (this 35 35 35) (not (at 35 35)))",
		},
	}

	for _, tc := range tests {
		inF := createFolFormula(Tokenize(tc.f))
		vars, objs := make([]Variable, len(tc.v)), make([]Object, len(tc.o))
		for i, v := range tc.v {
			vars[i] = Variable{v, ""}
			objs[i] = *NewTypedObject(tc.o[i], "")
		}
		outF := multiFolSubstitute(inF, vars, objs)
		if outF.Pddl() != tc.exp {
			t.Errorf("ALARM:\nexp:%s\nwas:%s\n", tc.exp, outF.Pddl())
		}

	}
}

// _____________________________________________________________________________
func TestContainsVariable(t *testing.T) {
	type testCase struct {
		f   string // formula
		v   string // variable
		out bool   // expected output
	}

	tests := []testCase{
		{"(and (a ?x ?y ?z) (b ?a ?b? ?c))", "?x", true},
		{"(and (a ?x ?y ?z) (b ?a ?b? ?c))", "?u", false},
		{"(and (not (ontable ?x)) (not (clear ?x)) (not (handempty ?a)) (holding ?a ?x))", "?x", true},
		{"(and (not (ontable ?x)) (not (clear ?x)) (not (handempty ?y)) (holding ?x ?a))", "?a", true},
		{"(and (not (ontable ?x)) (not (clear ?x)) (not (handempty ?a)) (holding ?a ?x))", "?z", false},
	}

	for _, tc := range tests {
		f := createFolFormula(Tokenize(tc.f))
		v := Variable{tc.v, "kind"}
		if containsVariable(f, v) != tc.out {
			t.Errorf("test Variable in fol formula yields wrong result for:\nf: %s\nv: %s\n", tc.f, tc.v)
		}
	}
}

// _____________________________________________________________________________
func TestIsVar(t *testing.T) {
	vars := []string{"?hal", "?p", "?", "?this-is-a-var", "? a", "?a b"}
	novars := []string{"!no", "notavar", "@nonono", " nope", " no ho "}
	for _, v := range vars {
		if !isVar(v) {
			t.Errorf("Variable detection fault!\nexp: isVar(%s) == true, but was false!", v)
		}
	}
	for _, nv := range novars {
		if isVar(nv) {
			t.Errorf("Variable detection fault!\nexp: isVar(%s) == false, but was true!", nv)
		}
	}
}

// _____________________________________________________________________________
func TestRemoveLiteral(t *testing.T) {
	type testCase struct {
		in  []string
		out string
	}

	tests := []testCase{
		{[]string{"( and ( a ?x ?y ) ( b ?x ?y ) )", "( a ?x ?y )"},
			"(b ?x ?y)"},
		{[]string{"( and ( a ?x ?y ) ( b ?x ?y ) ( c ?x ?y ) )", "( b ?x ?y )"},
			"(and (a ?x ?y) (c ?x ?y))"},
		{[]string{"( and ( a ?x ?y ) ( and ( b ?x ?y ) ( d ?x ?y ) ) ( b ?x ?y ) )", "( b ?x ?y )"},
			"(and (a ?x ?y) (d ?x ?y))"},
		{[]string{"( a ?x ?y )", "( b ?x ?y )"}, "(a ?x ?y)"},
		{[]string{"( a ?x ?y )", "( a ?x ?y )"}, "()"},
		{[]string{"( and ( b ) ( and ( b ) ( and ( a ?x ?y ) ( b ) ) ) )", "( b )"},
			"(a ?x ?y)"},
		{[]string{"( and ( b ) ( c ) ( and ( b ) ( and ( a ?x ?y ) ( b ) ) ) )", "( b )"},
			"(and (c) (a ?x ?y))"},
		{[]string{"( and ( a ) ( a ) ( and ( a ) ( and ( a ?x ?y ) ( a ) ) ) )", "( a )"},
			"()"},
	}

	for _, tc := range tests {
		folTokens := strings.Fields(tc.in[0])
		literal := createLiteral(strings.Fields(tc.in[1]))
		folFormula := createFolFormula(folTokens)
		newFormula := removeLiteral(folFormula, literal)
		if newFormula.Pddl() != tc.out {
			t.Errorf("malformed formula!\nexp:%s\nwas:%s\n", tc.out, newFormula.Pddl())
		}
		origFormula := createFolFormula(strings.Fields(tc.in[0]))
		if folFormula.Pddl() != origFormula.Pddl() {
			t.Errorf("the original formula was modified!\nexp:%s\nwas:%s\n", tc.in[0], folFormula.Pddl())
		}
	}
}

// _____________________________________________________________________________
func TestCreateFolFormula(t *testing.T) {
	type testCase struct {
		in       string
		expected string
	}

	tests := []testCase{
		{" ( at x y ) ", "(at x y)"},
		{" ( and ( at x y ) ( and ( at y z ) ( not ( in a b ) ) ) ( below y ) ) ",
			"(and (at x y) (and (at y z) (not (in a b))) (below y))"},
		{"( and ( a ?x ?y ) ( and ( b ?x ?y ) ( d ?x ?y ) ) ( b ?x ?y ) )",
			"(and (a ?x ?y) (and (b ?x ?y) (d ?x ?y)) (b ?x ?y))"},
	}

	for _, tc := range tests {
		tokens := strings.Fields(tc.in)
		_, rest := extractParanthesis(tokens)
		if len(rest) != 0 {
			t.Error("wrong rest")
		}
		got := createFolFormula(tokens)
		if got.String() != tc.expected {
			t.Errorf("\nexp:%s\ngot:%s", tc.expected, got)
		}
	}
}

// _____________________________________________________________________________
func TestLiteralPddl(t *testing.T) {
	type testCase struct {
		in       string
		expected string
	}

	tests := []testCase{
		{"( at ?l p1 )", "(at ?l p1)"},
		{"( not ( in ?p ?c ) )", "(not (in ?p ?c))"},
	}

	for _, tc := range tests {
		tl := strings.Fields(tc.in)
		literal := createLiteral(tl)
		if literal.Pddl() != tc.expected {
			t.Errorf("wrong pddl representation\nexp:%s\nwas:%s\n", tc.expected, literal.Pddl())
		}
	}
}

// _____________________________________________________________________________
func TestConjunctionPddl(t *testing.T) {
	type testCase struct {
		in       string
		expected string
	}

	tests := []testCase{
		{"( and ( in ?l p1 ) ( on x y ) )", "(and (in ?l p1) (on x y))"},
		{"( and ( not ( in ?p ?c ) ) ( and ( at x y ) ( below y z ) ) )",
			"(and (not (in ?p ?c)) (and (at x y) (below y z)))"},
	}

	for _, tc := range tests {
		tl := strings.Fields(tc.in)
		conjunction := createFolFormula(tl)
		if conjunction.Pddl() != tc.expected {
			t.Errorf("wrong pddl representation\nexp:%s\nwas:%s\n", tc.expected, conjunction.Pddl())
		}
	}
}
