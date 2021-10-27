package state

import (
	"testing"
)

//func TestTokenRegistry(t *testing.T) {
//	reg := NewTokenRegistry()
//	tokens := []TokenArray{
//		{1, 2, 3, 4},
//		{35, 77, 100, 33},
//		{1, 2, 3, 4},
//		{11, 12, 13, 14},
//		{11, 12, 13, 14},
//		{1, 2, 3, 4},
//	}
//	ids := []TokenID{
//		0,
//		1,
//		0,
//		2,
//		2,
//		0,
//	}
//
//	for i := range tokens {
//		tid := reg.Register(tokens[i])
//		if tid != ids[i] {
//			t.Fatalf("tokenID not as excpected\nexp:%v\nwas:%v\n", ids[i], tid)
//		}
//		if !reg.Lookup(tid).Equals(tokens[i]) {
//			t.Fatalf("tokenID not as excpected\nexp:%v\nwas:%v\n", ids[i], tid)
//		}
//	}
//}

func TestAdd(t *testing.T) {
	type testcase struct {
		tokens TokenArray
		token  int32
		pos    int
		exp    TokenArray
	}
	tests := []testcase{
		{TokenArray{1, 2, 3}, 35, 0, TokenArray{35, 1, 2, 3}},
		{TokenArray{1, 2, 3}, 35, 1, TokenArray{1, 35, 2, 3}},
		{TokenArray{1, 2, 3}, 35, 2, TokenArray{1, 2, 35, 3}},
		{TokenArray{1, 2, 3}, 35, 3, TokenArray{1, 2, 3, 35}},
		{TokenArray{}, 35, 0, TokenArray{35}},
	}

	for _, tc := range tests {
		res := tc.tokens.Add(tc.token, tc.pos)
		if !res.Equals(tc.exp) || res[tc.pos] != tc.token {
			t.Errorf("ERR: exp:%v was:%v\n", tc.exp, res)
		}
	}
}

func TestExcept(t *testing.T) {
	type testcase struct {
		tokens TokenArray
		pos    int
		exp    TokenArray
	}
	tests := []testcase{
		{TokenArray{1, 2, 3}, 0, TokenArray{2, 3}},
		{TokenArray{1, 2, 3}, 1, TokenArray{1, 3}},
		{TokenArray{1, 2, 3}, 2, TokenArray{1, 2}},
		{TokenArray{1}, 0, TokenArray{}},
	}

	for _, tc := range tests {
		res := tc.tokens.Except(tc.pos)
		if !res.Equals(tc.exp) {
			t.Errorf("ERR: exp:%v was:%v\n", tc.exp, res)
		}
	}
}

func TestTokenArrayJoin(t *testing.T) {
	a := NewTokenArray([]int32{1, 2, 3})
	b := NewTokenArray([]int32{4, 5, 6})
	want := NewTokenArray([]int32{1, 2, 3, 4, 5, 6})
	joint := a.Join(b)
	if !joint.Equals(want) {
		t.Errorf("Error joining TokenArrays!\nexp:%v\nwas:%v\n", want, joint)
	}
}

func TestTokenArraySplit(t *testing.T) {
	a := NewTokenArray([]int32{1, 2, 3, 4, 5, 6})
	want1 := NewTokenArray([]int32{1, 2, 3})
	want2 := NewTokenArray([]int32{4, 5, 6})
	res1, res2 := a.Split()
	if !res1.Equals(want1) || !res2.Equals(want2) {
		t.Errorf("Error joining TokenArrays!\nexp:%v\nwas:%v\n", want1, res1)
	}
}

// TODO: profile hash-functions!!
