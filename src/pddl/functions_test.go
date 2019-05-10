package pddl

import "testing"

//(:functions
//	(total-cost) - number
//	(travel-slow ?f1 - count ?f2 - count) - number
//	(travel-fast ?f1 - count ?f2 - count) - number
//)

//	(= (travel-slow n0 n2) 7)
//	(= (travel-slow n0 n3) 8)
//	(= (travel-slow n0 n4) 9)

func TestNumericFluentInit(t *testing.T) {
	in := []string{
		"(= (travel-slow n0 n2) 7)",
		"(= (travel-slow n0 n3) 8)",
		"(= some-name 35)",
	}
	for i, s := range in {
		numFluent := new(NumericFluent).Init(Tokenize(s))
		if numFluent.String() != in[i] {
			t.Errorf("NumericFluent Initialization Error.\n"+
				"exp:%s\nwas:%s\n", in[i], numFluent.String())
		}
	}
}

func TestFunctionInit(t *testing.T) {
	in := []string{
		"(total-cost) - number",
		"(travel-slow ?f1 - count ?f2 - count) - number",
		"(travel-fast ?f1 - count ?f2 - count) - number",
	}
	for i, s := range in {
		f := new(Function).Init(Tokenize(s))
		if f.String() != in[i] {
			t.Errorf("Function Initialization Error.\n"+
				"exp:%s\nwas:%s\n", in[i], f.String())
		}
	}
}

func TestFunctionString(t *testing.T) {
	fs := []Function{
		{"travel-slow",
			[]Object{{"?f1", "count", false}, {"?f2", "count", false}},
			"number",
		},
		{"total-cost",
			[]Object{},
			"number",
		},
	}
	res := []string{
		"(travel-slow ?f1 - count ?f2 - count) - number",
		"(total-cost) - number",
	}

	for i, f := range fs {
		if f.String() != res[i] {
			t.Errorf("Wrong String representation of pddl.Function!\n"+
				"exp:%s\nwas:%s\n", res[i], f.String())
		}
	}
}
