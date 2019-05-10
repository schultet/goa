package state

type (
	Hash32 uint32
)

const (
	offset32 = 2166136261
	prime32  = 16777619
)

// NewHash32 returns a new hash32 object implementing 32-bit FNV-1 hash.
func NewHash32() *Hash32 {
	var s Hash32 = offset32
	return &s
}

// Hash returns a uint32 hash value for an int slice
func (s *Hash32) Hash(data []int) uint32 {
	hash := Hash32(offset32)
	for _, c := range data {
		hash *= prime32
		hash ^= Hash32(c)
	}
	return uint32(hash)
}

// Hash returns a uint32 hash value for an int slice, ignoring the value at
// index i
func (s *Hash32) HashIgnore(data []int, i int) uint32 {
	hash := Hash32(offset32)
	for j, c := range data {
		if j == i {
			continue
		}
		hash *= prime32
		hash ^= Hash32(c)
	}
	return uint32(hash)
}

// Hash returns a uint32 hash value for an int32 slice
func (s *Hash32) Hash32(data []int32) uint32 {
	hash := Hash32(offset32)
	for _, c := range data {
		hash *= prime32
		hash ^= Hash32(c)
	}
	return uint32(hash)
}
