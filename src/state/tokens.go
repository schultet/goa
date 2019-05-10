package state

// TokenArray holds a token (= unique ID) for each agent (=index)
type TokenArray []int32

// NewTokenArray takes a int32 slice and returns a TokenArray of it
func NewTokenArray(vs []int32) TokenArray {
	return TokenArray(vs)
}

// Ints returns the []int32 representation of a TokenArray
func (a TokenArray) Ints() []int32 {
	res := make([]int32, len(a))
	for i, sid := range a {
		res[i] = int32(sid)
	}
	return res
}

// Equals returns true iff the TokenArrays' underlying int slices are the same
func (a TokenArray) Equals(other TokenArray) bool {
	if len(a) != len(other) {
		return false
	}
	for i, x := range a {
		if x != other[i] {
			return false
		}
	}
	return true
}

// EqualsIgnore returns whether the tokenArrays are equal except for tokenIndex
func (a TokenArray) EqualsIgnore(other TokenArray, tokenIndex int) bool {
	if len(a) != len(other) {
		return false
	}
	for i, x := range a {
		if i != tokenIndex && x != other[i] {
			return false
		}
	}
	return true
}

// Copy returns a copy of a tokenArray
func (a TokenArray) Copy() TokenArray {
	cpy := make(TokenArray, len(a))
	copy(cpy, a)
	return cpy
}

// Except returns a new TokenArray without the token at agentID
func (a TokenArray) Except(agentID int) TokenArray {
	return append(a[:agentID], a[agentID+1:]...)
}

// Join merges two tokenarrays
func (a TokenArray) Join(other TokenArray) TokenArray {
	return append(a, other...)
}

// Split splits a tokenarray in the middle into two tokenarrays
func (a TokenArray) Split() (TokenArray, TokenArray) {
	n := len(a) / 2
	return a[:n], a[n:]
}

// Add adds a token at a specific position in the tokenarray
func (a TokenArray) Add(token int32, pos int) TokenArray {
	res := make(TokenArray, len(a)+1)
	copy(res, a[:pos])
	res[pos] = token
	copy(res[pos+1:], a[pos:])
	return res
}
