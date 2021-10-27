package task

import (
	"fmt"
	"sort"
	"testing"

	"github.com/schultet/goa/pkg/util/slice"
)

func createVars(vs []int) []*Variable {
	res := make([]*Variable, len(vs))
	for i, v := range vs {
		res[i] = &Variable{ID: v}
	}
	return res
}

func TestConditionSetProject(t *testing.T) {
	type testcase struct {
		cs   []int
		vars []int
		exp  []int
	}
	tests := []testcase{
		{[]int{0, 0, 1, 1}, []int{1}, []int{1, 1}},
		{[]int{0, 0, 1, 1}, []int{0}, []int{0, 0}},
		{[]int{0, 0, 1, 1}, []int{}, []int{}},
		{[]int{}, []int{1, 1}, []int{}},
		{[]int{0, 0}, []int{1, 1}, []int{}},
		{[]int{0, 0, 1, 1, 2, 2}, []int{0, 2}, []int{0, 0, 2, 2}},
		{[]int{0, 0, 1, 1, 2, 2}, []int{1}, []int{1, 1}},
	}
	for _, tc := range tests {
		cs := CreateConditionSet(tc.cs)
		vars := createVars(tc.vars)
		fmt.Printf("%+v\n", cs.Projection(vars))
	}
}

func TestConditionConsistent(t *testing.T) {
	type testcase struct {
		c1  []int
		exp bool
	}

	tests := []testcase{
		{[]int{0, 1, 0, 1}, false},
		{[]int{0, 1, 0, 2}, false},
		{[]int{0, 0, 1, 1, 2, 2, 0, 1}, false},
		{[]int{}, true},
		{[]int{0, 0}, true},
		{[]int{0, 1, 1, 1}, true},
		{[]int{35, 77, 0, 1, 2, 1}, true},
	}

	for _, tc := range tests {
		c1 := CreateConditionSet(tc.c1)
		res := c1.Consistent()
		if res != tc.exp {
			t.Fatalf("\nexp: %v\nwas: %v\n", tc.exp, res)
		}
	}
}

func TestConditionSetExcept(t *testing.T) {
	type testcase struct {
		c1  []int
		rem []int
		exp []int
	}

	tests := []testcase{
		{[]int{0, 1}, []int{2, 3, 4}, []int{0, 1}},
		{[]int{0, 1, 1, 2}, []int{0, 1}, []int{}},
		{[]int{0, 0, 1, 1, 2, 2, 3, 3}, []int{0, 3}, []int{1, 1, 2, 2}},
		{[]int{0, 1, 1, 1}, []int{}, []int{0, 1, 1, 1}},
		{[]int{}, []int{0, 1, 2, 3}, []int{}},
		{[]int{}, []int{}, []int{}},
	}

	for _, tc := range tests {
		c1 := CreateConditionSet(tc.c1)
		res := c1.Except(tc.rem...)
		exp := CreateConditionSet(tc.exp)
		if !res.Equals(exp) {
			t.Fatalf("\nexp: %v\nwas: %v\n", exp, res)
		}
	}
}

func TestConditionSetSubsetOf(t *testing.T) {
	type testcase struct {
		c1  []int
		c2  []int
		exp bool
	}

	tests := []testcase{
		{[]int{0, 1, 1, 1}, []int{0, 1, 2, 2, 1, 1}, true},
		{[]int{}, []int{35, 77, 0, 1, 2, 1}, true},
		{[]int{0, 1, 1, 1}, []int{0, 1}, false},
		{[]int{0, 1, 1, 1}, []int{0, 1, 2, 1}, false},
		{[]int{0, 1, 1, 1}, []int{0, 1, 2, 1, 3, 3, 1, 1}, true},
		{[]int{35, 77, 0, 1, 2, 1}, []int{}, false},
	}

	for _, tc := range tests {
		c1 := CreateConditionSet(tc.c1)
		c2 := CreateConditionSet(tc.c2)
		res := c1.SubsetOf(c2)
		if res != tc.exp {
			t.Fatalf("\nexp: %v\nwas: %v\n", tc.exp, res)
		}
	}
}

func TestConditionSetIntersection(t *testing.T) {
	type testcase struct {
		c1  []int
		c2  []int
		exp []int
	}

	tests := []testcase{
		{
			[]int{0, 1, 3, 3, 4, 4, 1, 1},
			[]int{0, 1, 2, 2, 1, 1},
			[]int{0, 1, 1, 1},
		},
		{[]int{0, 1, 1, 1}, []int{1, 1}, []int{1, 1}},
		{[]int{0, 1, 1, 1}, []int{2, 1, 0, 1}, []int{0, 1}},
		{[]int{}, []int{35, 77, 0, 1, 2, 1}, []int{}},
		{[]int{35, 77, 0, 1, 2, 1}, []int{}, []int{}},
		{[]int{}, []int{}, []int{}},
	}

	for _, tc := range tests {
		c1 := CreateConditionSet(tc.c1)
		c2 := CreateConditionSet(tc.c2)
		res := c1.Intersection(c2).Ints()
		if !slice.EqualIntArrays(res, tc.exp) {
			t.Fatalf("\nexp: %+v\nwas: %+v\n", tc.exp, res)
		}
	}
}

func TestSortConditionSet(t *testing.T) {
	type testcase struct {
		inp []int
		exp []int
	}
	tests := []testcase{
		{[]int{4, 1, 2, 0, 3, 7}, []int{2, 0, 3, 7, 4, 1}},
		{[]int{4, 1, 2, 0}, []int{2, 0, 4, 1}},
		{[]int{1, 0}, []int{1, 0}},
		{[]int{}, []int{}},
	}
	for _, tc := range tests {
		res := CreateConditionSet(tc.inp)
		sort.Sort(res)
		if !slice.EqualIntArrays(res.Ints(), tc.exp) {
			t.Fatalf("ERR: exp: %v was: %v\n", tc.exp, res)
		}
	}
}

func TestActionProjection(t *testing.T) {
	type testcase struct {
		pre    []int
		eff    []int
		vars   []int
		expPre []int
		expEff []int
	}
	tests := []testcase{
		{
			[]int{0, 0, 1, 1, 2, 2, 3, 3},
			[]int{4, 4, 0, 0, 1, 1},
			[]int{0, 4},
			[]int{0, 0},
			[]int{4, 4, 0, 0},
		},
		{
			[]int{},
			[]int{},
			[]int{0, 1, 2},
			[]int{},
			[]int{},
		},
		{
			[]int{0, 0},
			[]int{1, 1},
			[]int{0, 1, 2},
			[]int{0, 0},
			[]int{1, 1},
		},
		{
			[]int{0, 0},
			[]int{1, 1},
			[]int{},
			[]int{},
			[]int{},
		},
	}
	for _, tc := range tests {
		o := &Action{
			Preconditions: createPreconditions(tc.pre),
			Effects:       CreateEffects(tc.eff),
		}
		vars := createVars(tc.vars)
		projection := o.Projection(vars)
		resEff := EffectsToIntSlice(projection.Effects)
		resPre := ConditionSetToIntSlice(projection.Preconditions)

		if !slice.EqualIntArrays(resEff, tc.expEff) {
			t.Fatalf("ERR: exp:%v\nwas:%v\n", tc.expEff, resEff)
		}
		if !slice.EqualIntArrays(resPre, tc.expPre) {
			t.Fatalf("ERR: exp:%v\nwas:%v\n", tc.expPre, resPre)
		}
	}
}

func createPreconditions(a []int) []Condition {
	res := make([]Condition, 0, len(a)/2+1)
	for i := 0; i < len(a); i += 2 {
		res = append(res, Condition{a[i], a[i+1]})
	}
	return res
}

//func preconditionsToIntSlice(cs []Condition) []int {
//	res := make([]int, 0, len(cs)*2)
//	for _, c := range cs {
//		res = append(res, c.Variable, int(c.Value))
//	}
//	return res
//}
