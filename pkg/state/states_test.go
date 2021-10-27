package state

import (
	"fmt"
	"testing"

	"github.com/schultet/goa/pkg/util/ints"
)

//func TestReplacePublic(t *testing.T) {
//	type testcase struct {
//		s, pub, res State
//	}
//	tests := []testcase{
//		{State{1, 1, 0, 0, 1}, State{0, 0}, State{0, 0, 0, 0, 1}},
//		{State{0, 0, 0, 0, 0}, State{0, 1}, State{0, 1, 0, 0, 0}},
//		{State{1, 1}, State{0, 0}, State{0, 0}},
//		{State{1}, State{}, State{1}},
//		{State{1, 1}, State{}, State{1, 1}},
//	}
//
//	for _, tc := range tests {
//		res := tc.s.ReplacePublic(tc.pub)
//		if !res.Equals(tc.res) {
//			t.Errorf("exp: %v\nwas: %v\n", tc.res, res)
//		}
//
//	}
//}

//func TestSuccessor(t *testing.T) {
//	type testcase struct {
//		s, res State
//		o      task.Action
//	}
//	tests := []testcase{
//		{
//			State{0, 0, 1, 1},
//			State{1, 1, 1, 1},
//			task.Action{effects: []effect{{0, 1}, {1, 1}}},
//		},
//		{
//			State{0, 0, 1, 1},
//			State{1, 0, 0, 1},
//			task.Action{effects: []effect{{0, 1}, {2, 0}}},
//		},
//		{
//			State{1, 1, 1, 1},
//			State{1, 1, 1, 1},
//			task.Action{effects: []effect{}},
//		},
//		{
//			State{1, 1, 1, 1},
//			State{0, 0, 0, 0},
//			task.Action{effects: []effect{{0, 0}, {1, 0}, {2, 0}, {3, 0}}},
//		},
//	}
//
//	for _, tc := range tests {
//		succ := tc.s.successor(&tc.o)
//		if !succ.Equals(tc.res) {
//			t.Errorf("exp: %s\nwas: %s\n", tc.res, succ)
//		}
//	}
//}
// TODO: move test to location of successor function

//func TestCopy(t *testing.T) {
//	//TODO: implement test
//}
//
//func TestPublic(t *testing.T) {
//	packer := NewPhonyPacker()
//	type testcase struct {
//		s   State
//		n   int
//		pub State
//	}
//	tests := []testcase{
//		{State{0, 1, 35}, 2, State{0, 1}},
//		{State{7}, 1, State{7}},
//		{State{3}, 0, State{}},
//		{State{3, 5}, 2, State{3, 5}},
//	}
//	for _, tc := range tests {
//		reg := NewStateRegistry(len(tc.s), tc.n, packer)
//		if !reg.Public(tc.s).Equals(tc.pub) {
//			t.Errorf("StateRegistry-public() Error:\nexp:%v\nwas:%v",
//				tc.pub, reg.Public(tc.s))
//		}
//	}
//
//}

// TODO: make this a proper unit test (without fmt.Println statements)
func TestStateRegistrySharedState(t *testing.T) {
	s0 := State{1, 2, 3, 4, 5, 6, 34020302349, 19830498}
	s1 := State{0, 1, 0, 1, 1, 0, 2030232131, 0}
	ranges := []int{2, 3, 4, 5, 6, 7, ints.MaxValue}
	r := NewGStateRegistry(4, 2, 1, NewIntPacker(ranges)).(*GstateReg)

	sid0 := r.Register(s0)
	sid1 := r.Register(s1)
	fmt.Printf("SHARED: %v\n", r.SharedState(s0, sid0)) //r.Lookup(sid0)))
	fmt.Printf("GLOBAL: %v\n", r.GlobalState(r.SharedState(s0, sid0)))
	fmt.Printf("SHARED: %v\n", r.SharedState(s1, sid1)) //r.Lookup(sid1)))
	fmt.Printf("GLOBAL: %v\n", r.GlobalState(r.SharedState(s1, sid1)))
}

// TODO: make this a proper unit test (without fmt.Println statements)
func TestStateRegistryRegister(t *testing.T) {
	s0 := State{1, 2, 3, 4, 5, 6, 34020302349, 19830498}
	s1 := State{0, 1, 0, 1, 1, 0, 2030232131, 350}
	ranges := []int{2, 3, 4, 5, 6, 7, ints.MaxValue}

	r := NewGStateRegistry(4, 2, 1, NewIntPacker(ranges))

	sid0 := r.Register(s0)
	sid1 := r.Register(s1)
	fmt.Println("initial state id:", sid0)
	fmt.Println("initial state:", r.Lookup(sid0))
	fmt.Println("pub:", r.Public(s0))
	fmt.Println("pri:", r.Private(s0))
	fmt.Println("tok:", r.Tokens(s0))

	fmt.Println("s1 state id:", sid1)
	fmt.Println("s1 state:", s1)
	fmt.Println("s1 state:", r.Lookup(sid1))
	sha1 := r.SharedState(r.Lookup(sid1), sid1)
	fmt.Println("shared(s1):", sha1)
	fmt.Println("global(s1):", r.GlobalState(sha1))
}

func TestTokenRegistry(t *testing.T) {
	reg := NewGStateRegistry(2, 0, 1, NewPhonyPacker())
	states := []State{
		{1, 2, 3, 0},
		{35, 77, 100, 1},
		{1, 2, 3, 0},
		{11, 12, 13, 2},
		{11, 12, 13, 2},
		{1, 2, 3, 0},
	}

	for i := range states {
		sid := reg.Register(states[i])
		sid2 := reg.Register(states[i])
		if sid != sid2 {
			t.Fatalf("stateIDs not as excpected\nexp:%v\nwas:%v\n", sid, sid2)
		}
		if !reg.Lookup(sid).Equals(states[i]) {
			t.Fatalf("state Lookup failed \nexp:%v\nwas:%v\n", states[i], reg.Lookup(sid))
		}
	}
}
