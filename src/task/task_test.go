package task

import (
	"testing"

	"github.com/schultet/goa/src/util/slice"
)

func varIDs(vs []*Variable) []int {
	res := make([]int, len(vs))
	for i, v := range vs {
		res[i] = v.ID
	}
	return res
}

func TestGetVars(t *testing.T) {
	type testcase struct {
		privateVars []int
		publicVars  []int
	}
	tests := []testcase{
		{[]int{1, 2, 3}, []int{4, 5}},
		{[]int{}, []int{}},
		{[]int{}, []int{4}},
		{[]int{1}, []int{}},
	}
	for _, tc := range tests {
		vars := make([]*Variable, 0)
		for _, v := range tc.privateVars {
			vars = append(vars, &Variable{ID: v, IsPrivate: true})
		}
		for _, v := range tc.publicVars {
			vars = append(vars, &Variable{ID: v, IsPrivate: false})
		}
		task := &Task{Vars: vars}
		prires := task.PrivateVars()
		if !slice.EqualIntArrays(varIDs(prires), tc.privateVars) {
			t.Errorf("ERR: exp: %+v was: %+v\n", tc.privateVars, varIDs(prires))
		}
		pubres := task.PublicVars()
		if !slice.EqualIntArrays(varIDs(pubres), tc.publicVars) {
			t.Errorf("ERR: exp: %+v was: %+v\n", tc.publicVars, varIDs(pubres))
		}
	}
}

func TestIsConsistent(t *testing.T) {
	//type testcase struct {
	//	c   []int
	//	m   map[int]MutexGroup
	//	exp bool
	//}

	//m := make(map[int]MutexGroup, 0)
	//m[0] = MutexGroup{
	//	&VariableValuePair{0, 1},
	//	&VariableValuePair{0, 2},
	//	&VariableValuePair{0, 3},
	//}
	//m[1] = MutexGroup{&VariableValuePair{1, 0}}
	//m[2] = MutexGroup{&VariableValuePair{2, 0}}
	//m[3] = MutexGroup{&VariableValuePair{3, 0}, &VariableValuePair{3, 1}}

	//tests := []testcase{
	//	{[]int{0, 1, 0, 2, 1, 0}, m, false},
	//	{[]int{0, 1, 0, 4}, m, true},
	//}

	//for _, tc := range tests {
	//	c := make([]*VariableValuePair, 0)
	//	for i := 0; i < len(tc.c); i += 2 {
	//		c = append(c, &VariableValuePair{tc.c[i], byte(tc.c[i+1])})
	//	}
	//	res := IsConsistent(c, tc.m)
	//	if res != tc.exp {
	//		t.Fatalf("\nwas %v\n exp %v\n", res, tc.exp)
	//	}
	//}

}
