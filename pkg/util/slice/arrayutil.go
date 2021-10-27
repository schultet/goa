package slice

//
// This package provides some useful functions for working with int or string
// slices, such as comparing, copying, appending, etc...
//

// AppendIntArray appends a slice of ints to a slice of ints
func AppendIntArray(a, b *[]int) {
	*a = append(*a, (*b)...)
}

// CopyIntArray returns a copy of an int-slice
func CopyIntArray(a []int) []int {
	b := make([]int, len(a))
	copy(b, a)
	// alternative: b := append([]int{}, a...)
	return b
}

// filtering without allocation
/*
* b := a[:0]
* for _, x := range a {
*	if f(x) {
*		b = append(b, x)
*	}
* }
 */

// ReverseIntArray returns the reversed int-slice
func ReverseIntArray(a []int) {
	for l, r := 0, len(a)-1; l < r; l, r = l+1, r-1 {
		a[l], a[r] = a[r], a[l]
	}
}

// InsertIntArray inserts element e at position i into int-slice a
func InsertIntArray(a *[]int, e, i int) {
	*a = append(*a, 0)
	copy((*a)[i+1:], (*a)[i:])
	(*a)[i] = e
}

// EqualIntArrays returns true iff two int-slices are equal
func EqualIntArrays(a, b []int) bool {
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

// EqualStringArrays returns true iff two string slices are equal
func EqualStringArrays(a, b []string) bool {
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

// ContainSameInts checks whether two int-slices contain the same elements
func ContainSameInts(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	values := make(map[int]int)
	for i, v := range a {
		if _, ok := values[v]; ok {
			values[v]++
		} else {
			values[v] = 1
		}
		if _, ok := values[b[i]]; ok {
			values[b[i]]--
		} else {
			values[b[i]] = -1
		}
	}
	for _, v := range values {
		if v != 0 {
			return false
		}
	}
	return true
}

// ContainSameStrings checks whether two string-slices contain the same elements
func ContainSameStrings(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	values := make(map[string]int)
	for i, v := range a {
		if _, ok := values[v]; ok {
			values[v]++
		} else {
			values[v] = 1
		}
		if _, ok := values[b[i]]; ok {
			values[b[i]]--
		} else {
			values[b[i]] = -1
		}
	}
	for _, v := range values {
		if v != 0 {
			return false
		}
	}
	return true
}

// ContainSameUints checks whether two uint32 slices contain the same values
func ContainSameUints(a, b []uint32) bool {
	if len(a) != len(b) {
		return false
	}

	values := make(map[uint32]int)
	for i, v := range a {
		if _, ok := values[v]; ok {
			values[v]++
		} else {
			values[v] = 1
		}
		if _, ok := values[b[i]]; ok {
			values[b[i]]--
		} else {
			values[b[i]] = -1
		}
	}
	for _, v := range values {
		if v != 0 {
			return false
		}
	}
	return true
}
