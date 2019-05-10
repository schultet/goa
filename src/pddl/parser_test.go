package pddl

import (
	"strings"
	"testing"
)

// TODO: TEST parsing of a complete domain/problem
// important!

func TestParseFunctions(t *testing.T) {

	in := Tokenize(` (:functions (total-cost) - number
		(travel-slow ?f1 - count ?f2 - count) - number
		(travel-fast ?f1 - count ?f2 - count) - number)`)
	res := []string{
		"(total-cost) - number",
		"(travel-slow ?f1 - count ?f2 - count) - number",
		"(travel-fast ?f1 - count ?f2 - count) - number",
	}
	fs := ParseFunctions(in)
	for i, f := range fs {
		if f.String() != res[i] {
			t.Errorf("\nexp:%s\nwas:%s", res[i], f.String())
		}
	}
}

// TODO: testing for a non-empty ObjectList for constants
func TestExtractObjects(t *testing.T) {
	type testCase struct {
		predicates []string
		objects    []string
		result     []string
	}
	tests := []testCase{
		{
			[]string{"(p1 a b c)", "(pred2 d)"},
			[]string{"d", "a", "b"},
			[]string{"a", "b", "c", "d"},
		},
		{
			[]string{"(p1 a b c)", "(pred2 d)"},
			[]string{"d", "a", "c", "b"},
			[]string{"a", "b", "c", "d"},
		},
		{
			[]string{"(p1 a b c)"},
			[]string{},
			[]string{"a", "b", "c"},
		},
		{
			[]string{"(p1 d a b c c)", "(pred2 d d)"},
			[]string{"a", "b"},
			[]string{"a", "b", "c", "d"},
		},
	}

	for _, tc := range tests {
		predicates := make([]*Predicate, 0, len(tc.predicates))
		for _, s := range tc.predicates {
			predicates = append(predicates, NewPredicate(Tokenize(s)))
		}
		objects := ObjectList{}
		for _, o := range tc.objects {
			objects = append(objects, Object{name: o, kind: "object"})
		}
		expectedResult := ObjectList{}
		for _, o := range tc.result {
			expectedResult = append(expectedResult, Object{name: o, kind: "object"})
		}

		extractObjects(predicates, &objects, ObjectList{}, ObjectList{}, ObjectList{})
		if !objects.equals(expectedResult) {
			t.Errorf("exp:%s\nwas:%s", expectedResult, objects)
		}
	}
}

// _____________________________________________________________________________
func TestExtractParanthesis(t *testing.T) {
	type testCase struct {
		in  string
		exp [2]string
	}
	tests := []testCase{
		{"( )", [2]string{"( )", ""}},
		{"( and ( a ) ( and ( b ) ( c ) ) )",
			[2]string{"( and ( a ) ( and ( b ) ( c ) ) )", ""}},
		{"( a ) ( and ( b ) ( c ) ) )",
			[2]string{"( a )", "( and ( b ) ( c ) ) )"}},
	}

	for _, tc := range tests {
		tl := strings.Fields(tc.in)
		result, rest := extractParanthesis(tl)
		if result.String() != tc.exp[0] {
			t.Errorf("wrong result!\nexp:%s\ngot:%s", tc.exp[0], result)
		}
		if rest.String() != tc.exp[1] {
			t.Errorf("wrong result!\nexp:%s\ngot:%s", tc.exp[1], result)
		}
	}
}

// _____________________________________________________________________________
func TestTokenListString(t *testing.T) {
	tl := TokenList{"(", "and", "(", "a", ")", "(", "b", ")", ")"}
	if tl.String() != "( and ( a ) ( b ) )" {
		t.Error("tokenlist was: ", tl)
	}
}

// _____________________________________________________________________________
func TestParsePredicates(t *testing.T) {
	//typetokens := strings.Fields(`( :types location agent - object taxi passenger - agent ) `)
	//types, _ := ParseTypes(typetokens)

	tokens := strings.Fields(` ( :predicates ( directly-connected ?l1 - location ?l2 - location )
				( at ?a - agent ?l - location )
				( in ?p - passenger ?t - taxi )
				( empty ?t - taxi )
				( free ?l - location )

				( :private
				( goal-of ?p - passenger ?l - location )
				) ) `)
	if v := len(tokens); v != 54 {
		t.Error("Expected sth else", v)
	}

	preds, priv_preds, _ := ParsePredicates(tokens[2:])
	if len(preds) != 5 || len(priv_preds) != 1 {
		t.Error("Expected sth else", preds, priv_preds)
	}

	predNames := []string{}
	for _, p := range preds {
		predNames = append(predNames, p.name)
	}
	if !sameSlice(predNames, []string{"directly-connected", "at", "in", "empty", "free"}) {
		t.Error("wrong predicates:", predNames)
	}

	privPredNames := []string{}
	for _, p := range priv_preds {
		privPredNames = append(privPredNames, p.name)
	}
	if !sameSlice(privPredNames, []string{"goal-of"}) {
		t.Error("wrong predicates")
	}
}

// _____________________________________________________________________________
func TestParseConstants(t *testing.T) {
	//typetokens := strings.Fields(`( :types location agent - object taxi passenger - agent ) `)
	//types, _ := ParseTypes(typetokens)

	tokens := strings.Fields(`( :constants t1 t2 t3  - taxi
					p1 p2 p3 p4 p5 p6 p7  - passenger
					some untyped constants
				    ( :private
					 this is sparta - taxi
					 nothing more
				) ) `)
	if v := len(tokens); v != 30 {
		t.Error("Expected sth else")
	}

	constants, priv_constants, _ := ParseConstants(tokens)
	if len(constants) != 13 || len(priv_constants) != 5 {
		t.Error("Expected sth else")
	}
	for _, c := range constants {
		if c.name == "t1" || c.name == "t2" || c.name == "t3" {
			if c.kind == "" {
				t.Error("no type assigned to constant, expected", c.name, "to be of type 'taxi'")
			} else if c.kind != "taxi" {
				t.Error("this is shit")
			}
		}
	}

	tokens = Tokenize(`(:objects
	distributor1 - distributor
	distributor0 - distributor
	crate11 - crate
	crate10 - crate
	crate13 - crate
	crate12 - crate
	depot0 - depot
	crate14 - crate
	depot1 - depot
	pallet5 - pallet
	pallet4 - pallet
	pallet7 - pallet
	pallet6 - pallet
	pallet1 - pallet
	pallet0 - pallet
	pallet3 - pallet
	pallet2 - pallet
	crate9 - crate
	crate8 - crate
	truck1 - truck
	truck0 - truck
	truck3 - truck
	truck2 - truck
	crate5 - crate
	crate4 - crate
	crate7 - crate
	crate6 - crate
	crate1 - crate
	crate0 - crate
	crate3 - crate
	crate2 - crate

	(:private
		hoist5 - hoist
		hoist0 - hoist
	)
)`)
	objects, privObjects, _ := ParseConstants(tokens)
	objectNames := []string{"depot0", "depot1"}
	privObjectNames := []string{"hoist5", "hoist0"}
	for _, name := range objectNames {
		if !objects.contains(name) {
			t.Errorf("Parsing objects error!\n")
		}
	}
	for _, name := range privObjectNames {
		if !privObjects.contains(name) {
			t.Errorf("Parsing private objects error!\n")
		}
	}
}

// _____________________________________________________________________________
func sameSlice(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, va := range a {
		if va != b[i] {
			return false
		}
	}
	return true
}
