package state

import (
	"bytes"
	"fmt"
)

const (
	capacity uint32 = 80000
)

// State is a variable assignment, the index corresponds to the
// variable and the value represents its assigned value
type State []int

// ShareState is like State, but can be transmitted to other agents b/c it does
// not contain variables that are labelled private
type SharedState []int

// StateID is a unique id that corresponds to a position in the statePool
type StateID int32

// Copy creates a state copy
func (s State) Copy() State {
	cpy := make(State, len(s))
	copy(cpy, s)
	return cpy
}

func (s State) CopyExcept(i int) State {
	d := make(State, len(s)-1)
	copy(d, s[:i])
	copy(d[i:], s[i+1:])
	return d
}

// Equals returns true iff two states represent the same variable assignment
func (s State) Equals(other State) bool {
	if len(s) != len(other) {
		return false
	}
	for i := range s {
		if s[i] != other[i] {
			return false
		}
	}
	return true
}

// Bytes returns state s as byte array
func (s State) Bytes() []byte {
	var buf bytes.Buffer
	for _, v := range s {
		fmt.Fprintln(&buf, v)
	}
	return buf.Bytes()
}

// JSON returns a JSON representation of state s
func (s State) JSON() string {
	json := "["
	for v, val := range s {
		json += fmt.Sprintf("{\"%d\": %d},", v, val)
	}
	return json[:len(json)-1] + "]"
}

// TODO: improve documentation
// StateRegistry contains and maintains all generated states. Each state is
// associated with a unique stateID.
type StateRegistry interface {
	// Lookup returns a registered state for a given stateID
	Lookup(StateID) State
	// Register registers a state to the registry and returns a unique ID
	Register(State) StateID

	// Public returns the public part of a state
	Public(State) State
	// Private returns the private part of a state
	Private(State) State
	// Tokens returns the tokens associated with the state
	Tokens(State) State

	// TODO: doc
	SharedState(State, StateID) SharedState
	// TODO: doc
	GlobalState(SharedState) State
	// TODO: doc
	Token(SharedState) StateID
}

type GstateReg struct {
	packer        Packer
	hasher        *Hash32
	statePool     []PackedState
	privateOffset int
	tokenOffset   int
	tokenIndex    int
	agentID       int
	stateRegister map[uint32][]StateID // hash(s) -> sid (= index of s in pool)
}

// NewStateRegistry creates a new GstateReg object that
// implements the GstateReg interface. States are compressed with a
// StatePacker and stored in an int array.
func NewGStateRegistry(pubN, priN, agentID int, packer Packer) StateRegistry {
	return &GstateReg{
		packer:        packer,
		hasher:        NewHash32(),
		statePool:     make([]PackedState, 0, capacity),
		privateOffset: pubN,
		tokenOffset:   pubN + priN,
		tokenIndex:    pubN + priN + agentID,
		agentID:       agentID,
		stateRegister: make(map[uint32][]StateID),
	}
}

// Lookup returns the state registered with StateID sid
func (r *GstateReg) Lookup(sid StateID) State {
	s := append(r.packer.Unpack(r.statePool[int(sid)]), 0)
	copy(s[r.tokenIndex+1:], s[r.tokenIndex:])
	s[r.tokenIndex] = int(sid)
	return s
}

func (r *GstateReg) Token(s SharedState) StateID {
	return StateID(s[r.privateOffset+r.agentID])
}

// Register registers a state to the stateReg and returning a unique StateID
func (r *GstateReg) Register(s State) StateID {
	sid := StateID(len(r.statePool))
	h := r.hasher.HashIgnore(s, r.tokenIndex)
	ps := r.packer.Pack(s.CopyExcept(r.tokenIndex))

	ids, ok := r.stateRegister[h]
	if ok {
		for _, sidx := range ids {
			if r.statePool[sidx].Equals(ps) {
				return sidx
			}
		}
	}
	r.stateRegister[h] = append(ids, sid)
	r.statePool = append(r.statePool, ps)
	return sid
}

// Public returns the public part of the search state
func (r *GstateReg) Public(s State) State {
	return s[:r.privateOffset]
}

func (r *GstateReg) Private(s State) State {
	return s[r.privateOffset:r.tokenOffset]
}

func (r *GstateReg) Tokens(s State) State {
	return s[r.tokenOffset:]
}

// SharedState returns the public part of the state together with the tokens
// i.e. removes the private part from `s`
func (r *GstateReg) SharedState(s State, sid StateID) SharedState {
	snew := make(SharedState, r.privateOffset+len(s[r.tokenOffset:]))
	copy(snew, s[:r.privateOffset])
	copy(snew[r.privateOffset:], s[r.tokenOffset:])
	snew[r.privateOffset+r.agentID] = int(sid)
	return snew
}

// GlobalState returns the global state, given a shared state. It looks up the
// agents token and adds the corresponding private part to the shared state
func (r *GstateReg) GlobalState(s SharedState) State {
	news := r.Lookup(StateID(s[r.privateOffset+r.agentID])).Copy()
	copy(news, s[:r.privateOffset])
	copy(news[r.tokenOffset:], s[r.privateOffset:])
	return news
}

// GetInfo returns a string representing size and 'depth' of the stateRegister
func (r *GstateReg) GetInfo() string {
	l, n := 0, 0
	for _, ls := range r.stateRegister {
		l += len(ls)
		n++
	}
	return fmt.Sprintf("N=%d, len/N=%d\n", n, l/n)
}

//// LocalState consists of the public and private parts
//func (r *gstateReg) LocalState(s State) State {
//}

// StateSplitter splits a state into its private or public part
type StateSplitter struct {
	publicVars, privateVars int
}

// separate returns a copy of s reduced to the range i (incl) to j (excl)
func (sp *StateSplitter) separate(s State, i, j int) State {
	slice := make(State, j-i)
	copy(slice, s[i:j])
	return slice
}

// Public returns a copy of the public variables of state s
func (sp *StateSplitter) Public(s State) State {
	return sp.separate(s, 0, sp.publicVars)
}

// Private returns a copy of the private variables of state s
func (sp *StateSplitter) Private(s State) State {
	return sp.separate(s, sp.publicVars, sp.publicVars+sp.privateVars)
}
